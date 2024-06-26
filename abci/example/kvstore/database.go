package kvstore

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	db "github.com/tendermint/tm-db"
	"sort"
)

const (
	Stale uint8 = iota
	Ready
	Busy
)

// Additional constants for specific application states
const (
	Registered uint8 = 0 // Indicates an initial state or condition
	Processing uint8 = 1 // Indicates a state where processing is underway
)

type ClientInfo struct {
	Name  string
	Power uint64
}

// BuildKey generates a database key for a given Ethereum address.
func BuildKeyForClientRegistration(ethereumAddress string) []byte {
	return []byte(fmt.Sprintf("clientRegistration_%s", ethereumAddress))
}

// StoreClientInfo stores ClientInfo in the database under the key derived from the Ethereum address.
func StoreClientInfo(db db.DB, ethereumAddress string, info ClientInfo) error {
	dataBytes, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return db.Set(BuildKeyForClientRegistration(ethereumAddress), dataBytes)
}

// GetClientInfo retrieves ClientInfo from the database using the Ethereum address.
func GetClientInfo(db db.DB, ethereumAddress string) (ClientInfo, error) {
	dataBytes, err := db.Get(BuildKeyForClientRegistration(ethereumAddress))
	if err != nil {
		return ClientInfo{}, err
	}
	var info ClientInfo
	err = json.Unmarshal(dataBytes, &info)
	return info, err
}

type MinerInfo struct {
	Name          string   // The name of the miner
	Power         uint64   // The computational power of the miner, possibly in hashes per second
	ServiceTypes  []uint64 // An array of service type identifiers that the miner provides
	IP            string   // The IP address of the miner for network connections
	InitialStatus uint8
}

// BuildKeyForMinerRegistration generates a database key for a given Ethereum address.
func BuildKeyForMinerRegistration(ethereumAddress string) []byte {
	return []byte(fmt.Sprintf("minerRegistration_%s", ethereumAddress))
}

// StoreMinerInfo stores MinerInfo in the database under the key derived from the Ethereum address.
func StoreMinerInfo(db db.DB, ethereumAddress string, info MinerInfo) error {
	dataBytes, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return db.Set(BuildKeyForMinerRegistration(ethereumAddress), dataBytes)
}

// GetMinerInfo retrieves MinerInfo from the database using the Ethereum address.
func GetMinerInfo(db db.DB, ethereumAddress string) (MinerInfo, error) {
	dataBytes, err := db.Get(BuildKeyForMinerRegistration(ethereumAddress))
	if err != nil {
		return MinerInfo{}, err
	}
	var info MinerInfo
	err = json.Unmarshal(dataBytes, &info)
	return info, err
}

// BuildServiceTypeKey generates a database key for service type mappings.
func BuildServiceTypeKey(serviceType uint64) []byte {
	return []byte(fmt.Sprintf("serviceType_%d", serviceType))
}

// GetMinersForServiceType retrieves the list of miners for a given service type from the database.
func GetMinersForServiceType(db db.DB, serviceType uint64) ([]string, error) {
	key := BuildServiceTypeKey(serviceType)
	data, err := db.Get(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return []string{}, nil // No miners registered yet for this service type.
	}

	var miners []string
	err = json.Unmarshal(data, &miners)
	if err != nil {
		return nil, err
	}
	return miners, nil
}

// StoreMinersForServiceType stores the list of miners for a given service type in the database.
func StoreMinersForServiceType(db db.DB, serviceType uint64, miners []string) error {
	key := BuildServiceTypeKey(serviceType)
	data, err := json.Marshal(miners)
	if err != nil {
		return err
	}
	return db.Set(key, data)
}

// RegisterMiner registers a new miner and updates the service type mappings in the database.
func RegisterMiner(db db.DB, miner MinerInfo, minerAddress string) error {
	for _, serviceType := range miner.ServiceTypes {
		miners, err := GetMinersForServiceType(db, serviceType)
		if err != nil {
			return fmt.Errorf("failed to get miners for service type %d: %v", serviceType, err)
		}

		// Check if the miner's address is already listed under this service type.
		found := false
		for _, addr := range miners {
			if addr == minerAddress {
				found = true
				break
			}
		}

		// If not found, add the miner's address to the list and store it back.
		if !found {
			miners = append(miners, minerAddress)
			err = StoreMinersForServiceType(db, serviceType, miners)
			if err != nil {
				return fmt.Errorf("failed to store miners for service type %d: %v", serviceType, err)
			}
		}
	}

	return nil
}

const allMinersKey = "allMinerStatus"

type MinerStatuses map[string]uint8 // Map from address to status

// SaveMinerStatuses helper to store the entire map in the database.
func SaveMinerStatuses(db db.DB, statuses MinerStatuses) error {
	data, err := json.Marshal(statuses)
	if err != nil {
		return fmt.Errorf("error marshaling miner statuses: %v", err)
	}
	return db.Set([]byte(allMinersKey), data)
}

// LoadMinerStatuses helper to retrieve the entire map from the database.
func LoadMinerStatuses(db db.DB) (MinerStatuses, error) {
	data, err := db.Get([]byte(allMinersKey))
	if err != nil {
		return nil, fmt.Errorf("error retrieving miner statuses: %v", err)
	}
	if data == nil {
		return make(MinerStatuses), nil // Return an empty map if no data found
	}

	var statuses MinerStatuses
	err = json.Unmarshal(data, &statuses)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling miner statuses: %v", err)
	}
	return statuses, nil
}

// AddOrUpdateMinerStatus adds a new miner or updates an existing miner's status.
func AddOrUpdateMinerStatus(db db.DB, address string, status uint8) error {
	statuses, err := LoadMinerStatuses(db)
	if err != nil {
		return err
	}
	statuses[address] = status // Add new or update existing
	return SaveMinerStatuses(db, statuses)
}

// RemoveMinerStatus removes a miner's status from the map.
func RemoveMinerStatus(db db.DB, address string) error {
	statuses, err := LoadMinerStatuses(db)
	if err != nil {
		return err
	}
	delete(statuses, address) // Remove the miner from the map
	return SaveMinerStatuses(db, statuses)
}

// GetMinerStatus queries a single miner's status.
func GetMinerStatus(db db.DB, address string) (uint8, error) {
	statuses, err := LoadMinerStatuses(db)
	if err != nil {
		return 0, err
	}
	status, found := statuses[address]
	if !found {
		return 0, fmt.Errorf("miner status not found for address: %s", address)
	}
	return status, nil
}

// RemoveMinerFromServiceTypeMapping removes a miner's address from the service type mapping in the database.
func RemoveMinerFromServiceTypeMapping(db db.DB, serviceType uint64, minerAddr string) error {
	miners, err := GetMinersForServiceType(db, serviceType)
	if err != nil {
		return fmt.Errorf("failed to get miners for service type %d: %v", serviceType, err)
	}

	// Find and remove the miner's address
	for i, addr := range miners {
		if addr == minerAddr {
			miners = append(miners[:i], miners[i+1:]...)
			break
		}
	}

	// Store the updated miner list for the service type
	if err := StoreMinersForServiceType(db, serviceType, miners); err != nil {
		return fmt.Errorf("failed to store miners for service type %d: %v", serviceType, err)
	}

	return nil
}

// AddMinerToServiceTypeMapping adds a miner's address to the service type mapping in the database.
func AddMinerToServiceTypeMapping(db db.DB, serviceType uint64, minerAddr string) error {
	miners, err := GetMinersForServiceType(db, serviceType)
	if err != nil {
		return fmt.Errorf("failed to get miners for service type %d: %v", serviceType, err)
	}

	// Check if the miner's address is already in the list
	for _, addr := range miners {
		if addr == minerAddr {
			return nil // Already exists, no need to add
		}
	}

	// Add the miner's address to the list
	miners = append(miners, minerAddr)

	// Store the updated miner list for the service type
	if err := StoreMinersForServiceType(db, serviceType, miners); err != nil {
		return fmt.Errorf("failed to store miners for service type %d: %v", serviceType, err)
	}

	return nil
}

// BuildKeyForMinerRating generates a database key for a given miner's address.
func BuildKeyForMinerRating(minerAddress string) []byte {
	return []byte(fmt.Sprintf("minerRating_%s", minerAddress))
}

// StoreClientRating stores the map of client ratings in the database under the key derived from the miner's address.
func StoreClientRating(db db.DB, minerAddress string, ratings map[string]uint8) error {
	dataBytes, err := json.Marshal(ratings)
	if err != nil {
		return err
	}
	return db.Set(BuildKeyForMinerRating(minerAddress), dataBytes)
}

// GetClientRating retrieves the map of client ratings from the database using the miner's address.
func GetClientRating(db db.DB, minerAddress string) (map[string]uint8, error) {
	dataBytes, err := db.Get(BuildKeyForMinerRating(minerAddress))
	if err != nil {
		return nil, err
	}
	if dataBytes == nil {
		return make(map[string]uint8), nil // Return an empty map if no data found
	}

	var ratings map[string]uint8
	err = json.Unmarshal(dataBytes, &ratings)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

// GenerateHashForMinerInfo creates a unique hash for given miner information which includes
// the miner's address, any relevant metadata, and the block height.
func GenerateHashForServiceInfo(clientAddress string, metadata []byte, blockHeight int64) string {
	hasher := sha256.New()
	hasher.Write([]byte(clientAddress))                  // Include the miner's address
	hasher.Write(metadata)                               // Include the service request metadata
	hasher.Write([]byte(fmt.Sprintf("%d", blockHeight))) // Include the block height

	// Return the resulting hash as a hexadecimal string
	return hex.EncodeToString(hasher.Sum(nil))
}

// JobInfo represents information about a specific job or task associated with a service.
type JobInfo struct {
	ServiceID    string `json:"service_id"`    // The unique identifier for the service
	ClientID     string `json:"client_id"`     // The identifier of the client requesting the service
	ServiceType  uint64 `json:"service_type"`  // The numeric identifier of the type of service
	JobStatus    uint8  `json:"job_status"`    // The status of the job
	TimeoutBlock int64  `json:"timeout_block"` // The block index when the job will expire
}

// BuildKeyForMinerJob generates a database key for a given miner's job.
func BuildKeyForMinerJob(minerID string) []byte {
	return []byte(fmt.Sprintf("minerjobs_%s", minerID))
}

// StoreJobInfo updates or appends a JobInfo in the list stored in the database under the miner's ID.
func StoreJobInfo(db db.DB, minerID string, newJob JobInfo) error {
	key := BuildKeyForMinerJob(minerID)
	jobs, err := GetJobInfos(db, minerID)
	if err != nil {
		return err
	}

	// Check if the job with the same ServiceID already exists and update it
	updated := false
	for i, job := range jobs {
		if job.ServiceID == newJob.ServiceID {
			jobs[i] = newJob
			updated = true
			break
		}
	}

	// If the job is not found, append the new job to the list
	if !updated {
		jobs = append(jobs, newJob)
	}

	dataBytes, err := json.Marshal(jobs)
	if err != nil {
		return err
	}

	return db.Set(key, dataBytes)
}

// GetJobInfos retrieves a list of JobInfo from the database using the miner's ID.
func GetJobInfos(db db.DB, minerID string) ([]JobInfo, error) {
	key := BuildKeyForMinerJob(minerID)
	dataBytes, err := db.Get(key)
	if err != nil {
		return nil, err
	}
	if dataBytes == nil {
		return []JobInfo{}, nil // Return an empty list if no jobs are found
	}
	var jobs []JobInfo
	err = json.Unmarshal(dataBytes, &jobs)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

// GetJobInfoByServiceID retrieves a single JobInfo by ServiceID from a specified miner's list of jobs.
func GetJobInfoByServiceID(db db.DB, minerID string, serviceID string) (JobInfo, error) {
	key := BuildKeyForMinerJob(minerID)
	dataBytes, err := db.Get(key)
	if err != nil {
		return JobInfo{}, fmt.Errorf("failed to retrieve jobs for miner ID '%s': %v", minerID, err)
	}
	if dataBytes == nil {
		return JobInfo{}, fmt.Errorf("no jobs found for miner ID '%s'", minerID)
	}

	var jobs []JobInfo
	err = json.Unmarshal(dataBytes, &jobs)
	if err != nil {
		return JobInfo{}, fmt.Errorf("error unmarshaling jobs data: %v", err)
	}

	for _, job := range jobs {
		if job.ServiceID == serviceID {
			return job, nil
		}
	}

	return JobInfo{}, fmt.Errorf("no job found with ServiceID '%s'", serviceID)
}

// RemoveJobInfo removes a JobInfo from the list stored in the database under the miner's ID based on the ServiceID.
func RemoveJobInfo(db db.DB, minerID string, serviceID string) error {
	key := BuildKeyForMinerJob(minerID)
	jobs, err := GetJobInfos(db, minerID)
	if err != nil {
		return err
	}

	// Filter out the job with the given ServiceID
	updatedJobs := make([]JobInfo, 0)
	found := false
	for _, job := range jobs {
		if job.ServiceID != serviceID {
			updatedJobs = append(updatedJobs, job)
		} else {
			found = true
		}
	}

	// If no job was found with the given ServiceID, return an error
	if !found {
		return fmt.Errorf("no job found with ServiceID '%s'", serviceID)
	}

	// Serialize the updated list of jobs back to JSON
	dataBytes, err := json.Marshal(updatedJobs)
	if err != nil {
		return err
	}

	// Save the updated list back to the database
	return db.Set(key, dataBytes)
}

const allServiceRequestsKey = "allServiceRequests"

type ServiceRequest struct {
	ServiceID string
	MinerID   string
	Height    int64
}

type ServiceRequests []ServiceRequest

func SaveServiceRequests(db db.DB, requests ServiceRequests) error {
	data, err := json.Marshal(requests)
	if err != nil {
		return fmt.Errorf("error marshaling service requests: %v", err)
	}
	return db.Set([]byte(allServiceRequestsKey), data) // Use SetSync for immediate writes
}

func LoadServiceRequests(db db.DB) (ServiceRequests, error) {
	data, err := db.Get([]byte(allServiceRequestsKey))
	if err != nil {
		return nil, fmt.Errorf("error retrieving service requests: %v", err)
	}
	if data == nil {
		return make(ServiceRequests, 0), nil // Return an empty slice if no data found
	}

	var requests ServiceRequests
	err = json.Unmarshal(data, &requests)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling service requests: %v", err)
	}
	return requests, nil
}

func AddServiceRequest(db db.DB, serviceID, minerID string, height int64) error {
	requests, err := LoadServiceRequests(db)
	if err != nil {
		return err
	}
	requests = append(requests, ServiceRequest{ServiceID: serviceID, MinerID: minerID, Height: height})
	return SaveServiceRequests(db, requests)
}

func RemoveServiceRequest(db db.DB, serviceID string) error {
	requests, err := LoadServiceRequests(db)
	if err != nil {
		return err
	}

	// Filter out the service request by ServiceID only
	filteredRequests := make(ServiceRequests, 0)
	for _, req := range requests {
		if req.ServiceID != serviceID {
			filteredRequests = append(filteredRequests, req)
		}
	}

	return SaveServiceRequests(db, filteredRequests)
}

func RetainServiceRequestsAboveHeight(db db.DB, retainHeight int64) error {
	requests, err := LoadServiceRequests(db)
	if err != nil {
		return err
	}

	var lastIndex = -1
	for i, request := range requests {
		if request.Height <= retainHeight {
			lastIndex = i
		} else {
			break
		}
	}
	if lastIndex != -1 {
		requests = requests[lastIndex+1:]
	}
	return SaveServiceRequests(db, requests)
}

// ServiceTypeCount maps a service type to the number of times a miner has served that type.
type ServiceTypeCount map[uint64]int

// MinerServices contains a miner ID and a map of service types with their respective counts.
type MinerServices struct {
	MinerID      string           `json:"miner_id"`
	ServiceTypes ServiceTypeCount `json:"service_types"`
}

// BlockServices holds the block height and a list of miner services provided at that height.
type BlockServices struct {
	BlockHeight   int64           `json:"block_height"`
	MinerServices []MinerServices `json:"miner_services"`
}

// MinerWorkRecords manages a collection of BlockServices, tracking services provided by miners over time.
type MinerWorkRecords []BlockServices

func (mwr *MinerWorkRecords) AddOrUpdateRecord(newBlock BlockServices) error {
	index, found := mwr.FindByHeight(newBlock.BlockHeight)
	if found {
		(*mwr)[index] = newBlock
	} else {
		*mwr = append(*mwr, newBlock)
		sort.Slice(*mwr, func(i, j int) bool { return (*mwr)[i].BlockHeight < (*mwr)[j].BlockHeight })
	}
	return nil
}

func (mwr MinerWorkRecords) FindByHeight(height int64) (int, bool) {
	for i, bs := range mwr {
		if bs.BlockHeight == height {
			return i, true
		}
	}
	return -1, false
}

func (mwr *MinerWorkRecords) UpdateMinerServices(height int64, minerServices MinerServices) error {
	index, found := mwr.FindByHeight(height)
	if !found {
		return mwr.AddOrUpdateRecord(BlockServices{
			BlockHeight:   height,
			MinerServices: []MinerServices{minerServices},
		})
	}

	minerIndex, found := findMiner((*mwr)[index].MinerServices, minerServices.MinerID)
	if found {
		(*mwr)[index].MinerServices[minerIndex] = minerServices
	} else {
		(*mwr)[index].MinerServices = append((*mwr)[index].MinerServices, minerServices)
	}
	return nil
}

func findMiner(miners []MinerServices, minerID string) (int, bool) {
	for i, miner := range miners {
		if miner.MinerID == minerID {
			return i, true
		}
	}
	return -1, false
}

func (mwr *MinerWorkRecords) IncrementServiceType(height int64, minerID string, serviceType uint64) error {
	index, found := mwr.FindByHeight(height)
	if !found {
		newBlock := BlockServices{
			BlockHeight: height,
			MinerServices: []MinerServices{{
				MinerID:      minerID,
				ServiceTypes: ServiceTypeCount{serviceType: 1},
			}},
		}
		return mwr.AddOrUpdateRecord(newBlock)
	}

	minerIndex, found := findMiner((*mwr)[index].MinerServices, minerID)
	if !found {
		(*mwr)[index].MinerServices = append((*mwr)[index].MinerServices, MinerServices{
			MinerID:      minerID,
			ServiceTypes: ServiceTypeCount{serviceType: 1},
		})
		return nil
	}

	(*mwr)[index].MinerServices[minerIndex].ServiceTypes[serviceType]++
	return nil
}
