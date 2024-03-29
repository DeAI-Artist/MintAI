package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	evmclient "github.com/DeAI-Artist/MintAI/core/chains/evm/client"
	"github.com/DeAI-Artist/MintAI/core/internal/testutils"
)

func Test_NodeWrapError(t *testing.T) {
	t.Parallel()

	t.Run("handles nil errors", func(t *testing.T) {
		err := evmclient.Wrap(nil, "foo")
		assert.NoError(t, err)
	})

	t.Run("adds extra info to context deadline exceeded errors", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(testutils.Context(t), 0)
		defer cancel()

		err := ctx.Err()

		err = evmclient.Wrap(err, "foo")

		assert.EqualError(t, err, "foo call failed: remote eth node timed out: context deadline exceeded")
	})
}
