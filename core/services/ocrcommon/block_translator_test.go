package ocrcommon_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DeAI-Artist/MintAI/core/internal/testutils"
	"github.com/DeAI-Artist/MintAI/core/internal/testutils/evmtest"
	v2 "github.com/DeAI-Artist/MintAI/core/internal/testutils/evmtest/v2"
	"github.com/DeAI-Artist/MintAI/core/logger"
	"github.com/DeAI-Artist/MintAI/core/services/ocrcommon"
)

func Test_BlockTranslator(t *testing.T) {
	t.Parallel()

	ethClient := evmtest.NewEthClientMock(t)
	ctx := testutils.Context(t)
	lggr := logger.TestLogger(t)

	t.Run("for L1 chains, returns the block changed argument", func(t *testing.T) {
		bt := ocrcommon.NewBlockTranslator(v2.ChainEthMainnet(t).EVM(), ethClient, lggr)

		from, to := bt.NumberToQueryRange(ctx, 42)

		assert.Equal(t, big.NewInt(42), from)
		assert.Equal(t, big.NewInt(42), to)
	})

	t.Run("for optimism, uses the default translator", func(t *testing.T) {
		bt := ocrcommon.NewBlockTranslator(v2.ChainOptimismMainnet(t).EVM(), ethClient, lggr)
		from, to := bt.NumberToQueryRange(ctx, 42)
		assert.Equal(t, big.NewInt(42), from)
		assert.Equal(t, big.NewInt(42), to)

	})

	t.Run("for arbitrum, returns the ArbitrumBlockTranslator", func(t *testing.T) {
		bt := ocrcommon.NewBlockTranslator(v2.ChainArbitrumMainnet(t).EVM(), ethClient, lggr)
		assert.IsType(t, &ocrcommon.ArbitrumBlockTranslator{}, bt)

		bt = ocrcommon.NewBlockTranslator(v2.ChainArbitrumRinkeby(t).EVM(), ethClient, lggr)
		assert.IsType(t, &ocrcommon.ArbitrumBlockTranslator{}, bt)
	})
}
