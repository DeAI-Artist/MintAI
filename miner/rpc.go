package miner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RPCResponse defines the expected JSON response structure
type RPCResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		NodeInfo struct {
			ProtocolVersion struct {
				P2P   string `json:"p2p"`
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"protocol_version"`

			ID         string `json:"id"`
			ListenAddr string `json:"listen_addr"`
			Network    string `json:"network"`
			Version    string `json:"version"`
			Channels   string `json:"channels"`
			Moniker    string `json:"moniker"`
			Other      struct {
				TxIndex    string `json:"tx_index"`
				RPCAddress string `json:"rpc_address"`
			} `json:"other"`
		} `json:"node_info"`
		SyncInfo struct {
			LatestBlockHash     string `json:"latest_block_hash"`
			LatestAppHash       string `json:"latest_app_hash"`
			LatestBlockHeight   string `json:"latest_block_height"`
			LatestBlockTime     string `json:"latest_block_time"`
			EarliestBlockHash   string `json:"earliest_block_hash"`
			EarliestAppHash     string `json:"earliest_app_hash"`
			EarliestBlockHeight string `json:"earliest_block_height"`
			EarliestBlockTime   string `json:"earliest_block_time"`
			CatchingUp          bool   `json:"catching_up"`
		} `json:"sync_info"`
		ValidatorInfo struct {
			Address string `json:"address"`
			PubKey  struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"pub_key"`
			VotingPower string `json:"voting_power"`
		} `json:"validator_info"`
	} `json:"result"`
}

// QueryRPCStatus sends a GET request to the RPC endpoint and checks the response
func QueryRPCStatus(rpcEndpoint string) error {
	resp, err := http.Get(fmt.Sprintf("http://%s/status", rpcEndpoint))
	if err != nil {
		return fmt.Errorf("failed to reach the RPC endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("node unreachable, received non-200 status code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	var rpcResponse RPCResponse
	if err := json.Unmarshal(body, &rpcResponse); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	// Assuming 'jsonrpc' field must be "2.0" to consider the response valid
	if rpcResponse.Jsonrpc != "2.0" {
		return errors.New("invalid JSONRPC version returned or node unreachable")
	}

	fmt.Println("Node is reachable and JSON is valid.")
	return nil
}
