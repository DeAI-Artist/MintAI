package kvstore

import (
	"encoding/json"
	dbm "github.com/tendermint/tm-db"
	"reflect"
	"strconv"
	"testing"
)

// TestStoreAndGetClientInfo tests the storing and retrieval of ClientInfo in the database.
func TestStoreAndGetClientInfo(t *testing.T) {
	// Initialize the in-memory database
	db := dbm.NewMemDB()

	// Example client information
	clientInfo := ClientInfo{
		Name:  "Alice",
		Power: 10,
	}
	ethereumAddress := "0x123abc"

	// Store client information
	err := StoreClientInfo(db, ethereumAddress, clientInfo)
	if err != nil {
		t.Fatalf("StoreClientInfo failed: %s", err)
	}

	// Retrieve client information
	retrievedInfo, err := GetClientInfo(db, ethereumAddress)
	if err != nil {
		t.Fatalf("GetClientInfo failed: %s", err)
	}

	// Compare the stored and retrieved information
	if retrievedInfo.Name != clientInfo.Name || retrievedInfo.Power != clientInfo.Power {
		t.Errorf("Retrieved info does not match stored info. Got %+v, want %+v", retrievedInfo, clientInfo)
	}
}

func TestMapInMemDB(t *testing.T) {
	// Initialize an in-memory database
	memDB := dbm.NewMemDB()

	// Original map
	originalMap := map[string]int{"a": 1, "b": 2, "c": 3}

	// Serialize the map to JSON for storage
	jsonData, err := json.Marshal(originalMap)
	if err != nil {
		t.Fatalf("Failed to serialize map: %v", err)
	}

	// Store serialized data
	if err := memDB.Set([]byte("mapKey"), jsonData); err != nil {
		t.Fatalf("Failed to set data in memDB: %v", err)
	}

	// Retrieve and deserialize the map
	retrievedData, err := memDB.Get([]byte("mapKey"))
	if err != nil {
		t.Fatalf("Failed to get data from memDB: %v", err)
	}
	retrievedMap := make(map[string]int)
	if err := json.Unmarshal(retrievedData, &retrievedMap); err != nil {
		t.Fatalf("Failed to deserialize map: %v", err)
	}

	// Compare the original and retrieved maps
	if len(retrievedMap) != len(originalMap) {
		t.Fatalf("Mismatch in map sizes: expected %v, got %v", len(originalMap), len(retrievedMap))
	}
	for key, expectedValue := range originalMap {
		if value, exists := retrievedMap[key]; !exists || value != expectedValue {
			t.Fatalf("Mismatch for key %s: expected %v, got %v", key, expectedValue, value)
		}
	}
}

func TestStoreAndGetMinerInfo(t *testing.T) {
	// Initialize the in-memory database
	db := dbm.NewMemDB()

	// Example miner information
	minerInfo := MinerInfo{
		Name:         "Miner Bob",
		Power:        500,
		ServiceTypes: []uint64{1, 2, 3},
		IP:           "192.168.1.100:8080", // IP address and port combined in one string
	}
	ethereumAddress := "0x456def"

	// Store miner information
	err := StoreMinerInfo(db, ethereumAddress, minerInfo)
	if err != nil {
		t.Fatalf("StoreMinerInfo failed: %s", err)
	}

	// Retrieve miner information
	retrievedInfo, err := GetMinerInfo(db, ethereumAddress)
	if err != nil {
		t.Fatalf("GetMinerInfo failed: %s", err)
	}

	// Compare the stored and retrieved information
	if retrievedInfo.Name != minerInfo.Name || retrievedInfo.Power != minerInfo.Power ||
		len(retrievedInfo.ServiceTypes) != len(minerInfo.ServiceTypes) || retrievedInfo.IP != minerInfo.IP {
		t.Errorf("Retrieved info does not match stored info. Got %+v, want %+v", retrievedInfo, minerInfo)
	}

	// Optionally check if all service types match
	for i, serviceType := range retrievedInfo.ServiceTypes {
		if serviceType != minerInfo.ServiceTypes[i] {
			t.Errorf("Service type mismatch at index %d: got %d, want %d", i, serviceType, minerInfo.ServiceTypes[i])
		}
	}
}

func TestRegisterMiner(t *testing.T) {
	db := dbm.NewMemDB()

	minerInfo := MinerInfo{
		Name:         "Miner Bob",
		Power:        500,
		ServiceTypes: []uint64{1, 2},
		IP:           "192.168.1.100:8080",
	}
	minerAddress := "0x456def"

	// Case 1: No service type key exists yet
	err := RegisterMiner(db, minerInfo, minerAddress)
	if err != nil {
		t.Fatalf("RegisterMiner failed: %s", err)
	}
	for _, serviceType := range minerInfo.ServiceTypes {
		miners, _ := GetMinersForServiceType(db, serviceType)
		if len(miners) != 1 || miners[0] != minerAddress {
			t.Errorf("Expected [%s] for service type %d, got %v", minerAddress, serviceType, miners)
		}
	}

	// Case 2: Service type key exists and already includes the specified miner
	err = RegisterMiner(db, minerInfo, minerAddress)
	if err != nil {
		t.Fatalf("RegisterMiner failed on second call: %s", err)
	}
	for _, serviceType := range minerInfo.ServiceTypes {
		miners, _ := GetMinersForServiceType(db, serviceType)
		if len(miners) != 1 || miners[0] != minerAddress { // No duplicate should be added
			t.Errorf("Unexpected duplicate or change in miners list for service type %d, got %v", serviceType, miners)
		}
	}

	// Case 3: Service type key exists but does not include the specified miner
	newMinerAddress := "0x123abc"
	err = RegisterMiner(db, minerInfo, newMinerAddress)
	if err != nil {
		t.Fatalf("RegisterMiner failed on third call: %s", err)
	}
	for _, serviceType := range minerInfo.ServiceTypes {
		miners, _ := GetMinersForServiceType(db, serviceType)
		if !reflect.DeepEqual(miners, []string{minerAddress, newMinerAddress}) {
			t.Errorf("Expected [%s, %s] for service type %d, got %v", minerAddress, newMinerAddress, serviceType, miners)
		}
	}

	// Additional case: Database operations error handling
	// You could mock the database to throw errors and verify that RegisterMiner handles it properly.
	// This would typically require an interface for the database and a mock implementation.
}

// TestMinerStatusManagement tests adding, updating, querying, and removing miner statuses.
func TestMinerStatusManagement(t *testing.T) {
	db := dbm.NewMemDB()
	address := "0x123"
	initialStatus := Ready
	updatedStatus := Busy

	// Add a new miner
	if err := AddOrUpdateMinerStatus(db, address, initialStatus); err != nil {
		t.Fatalf("Failed to add new miner status: %v", err)
	}

	// Verify initial status
	status, err := GetMinerStatus(db, address)
	if err != nil || status != initialStatus {
		t.Errorf("Expected initial status %d for address %s; got %d", initialStatus, address, status)
	}

	// Update the miner's status
	if err := AddOrUpdateMinerStatus(db, address, updatedStatus); err != nil {
		t.Errorf("Failed to update miner status: %v", err)
	}

	// Verify updated status
	status, err = GetMinerStatus(db, address)
	if err != nil || status != updatedStatus {
		t.Errorf("Expected updated status %d for address %s; got %d", updatedStatus, address, status)
	}

	// Remove the miner's status
	if err := RemoveMinerStatus(db, address); err != nil {
		t.Errorf("Failed to remove miner status: %v", err)
	}

	// Verify removal
	_, err = GetMinerStatus(db, address)
	if err == nil {
		t.Errorf("Expected error for fetching status of removed miner, got none")
	}
}

// BenchmarkStatusOperations benchmarks adding, retrieving, and removing miner statuses.
func BenchmarkStatusOperations(b *testing.B) {
	db := dbm.NewMemDB()
	for n := 1; n <= b.N; n++ {
		address := "miner" + strconv.Itoa(n) // Create unique miner address
		status := uint8(n % 3)               // Cycle through 0, 1, 2 statuses

		// Add or update miner status
		b.Run("AddOrUpdate", func(b *testing.B) {
			if err := AddOrUpdateMinerStatus(db, address, status); err != nil {
				b.Error("Failed to add/update miner status:", err)
			}
		})

		// Get miner status
		b.Run("GetStatus", func(b *testing.B) {
			if _, err := GetMinerStatus(db, address); err != nil {
				b.Error("Failed to get miner status:", err)
			}
		})

		// Remove miner status
		b.Run("RemoveStatus", func(b *testing.B) {
			if err := RemoveMinerStatus(db, address); err != nil {
				b.Error("Failed to remove miner status:", err)
			}
		})
	}
}
