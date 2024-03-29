package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/DeAI-Artist/MintAI/core/chains/evm/utils"
	"github.com/DeAI-Artist/MintAI/core/internal/testutils"
)

// ChainID   *big.Big   `json:"chainID"`
// FromBlock uint64     `json:"fromBlock"`

// // Contract-specific
// EffectiveTransmitterAddress null.String    `json:"effectiveTransmitterAddress"`
// SendingKeys                 pq.StringArray `json:"sendingKeys"`

// // Mercury-specific
// FeedID *common.Hash `json:"feedID"`
func Test_RelayConfig(t *testing.T) {
	cid := testutils.NewRandomEVMChainID()
	fromBlock := uint64(2222)
	feedID := utils.NewHash()
	rawToml := fmt.Sprintf(`
ChainID = "%s"
FromBlock = %d
FeedID = "0x%x"
`, cid, fromBlock, feedID[:])

	var rc RelayConfig
	err := toml.Unmarshal([]byte(rawToml), &rc)
	require.NoError(t, err)

	assert.Equal(t, cid.String(), rc.ChainID.String())
	assert.Equal(t, fromBlock, rc.FromBlock)
	assert.Equal(t, feedID.Hex(), rc.FeedID.Hex())
}
