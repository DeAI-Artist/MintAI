package synchronization

import (
	"context"
	"encoding/hex"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/wsrpc"

	"github.com/DeAI-Artist/MintAI/core/internal/testutils"
	"github.com/DeAI-Artist/MintAI/core/logger"
	"github.com/DeAI-Artist/MintAI/core/services/synchronization/telem"
)

func TestUniClient(t *testing.T) {
	t.Skip("Incomplete", "https://smartcontract-it.atlassian.net/browse/BCF-2729")
	privKey, err := hex.DecodeString("TODO")
	require.NoError(t, err)
	pubKey, err := hex.DecodeString("TODO")
	require.NoError(t, err)
	t.Log(len(privKey), len(pubKey))
	lggr := logger.TestLogger(t)
	c, err := wsrpc.DialUniWithContext(testutils.Context(t),
		lggr,
		"TODO",
		privKey,
		pubKey)
	require.NoError(t, err)
	t.Log(c)
	client := telem.NewTelemClient(c)
	ctx, cancel := context.WithTimeout(testutils.Context(t), 500*time.Millisecond)
	resp, err := client.Telem(ctx, &telem.TelemRequest{
		Telemetry: []byte(`hello world`),
		Address:   "myaddress",
	})
	cancel()
	t.Log(resp, err)
	require.NoError(t, c.Close())
}
