package txs

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

// SelectPseudorandomMiner selects a miner pseudorandomly based on the block height, app hash, and service ID.
func SelectPseudorandomMiner(miners []string, blockHeight int64, appHash []byte, serviceID string) string {
	// Combine block height, app hash, and service ID into a single string
	combinedInput := fmt.Sprintf("%d%s%s", blockHeight, appHash, serviceID)

	// Generate a SHA-256 hash of the combined input
	hash := sha256.New()
	hash.Write([]byte(combinedInput))
	hashBytes := hash.Sum(nil)

	// Convert the hash to a big integer
	hashInt := new(big.Int)
	hashInt.SetBytes(hashBytes)

	// Get the index within the range of the list of miners
	index := new(big.Int).Mod(hashInt, big.NewInt(int64(len(miners)))).Int64()

	// Select the miner at the generated index
	selectedMiner := miners[index]

	return selectedMiner
}

/*
func main() {
	// Example usage
	miners := []string{"miner1", "miner2", "miner3", "miner4"}
	blockHeight := 123456
	appHash := "exampleAppHash"
	nonce := 7890

	selectedMiner := selectPseudorandomMiner(miners, blockHeight, appHash, nonce)
	fmt.Println("Selected Miner:", selectedMiner)
}
*/
