package txmgr_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	txmgrcommon "github.com/DeAI-Artist/MintAI/common/txmgr"
	"github.com/DeAI-Artist/MintAI/core/chains/evm/txmgr/mocks"
	"github.com/DeAI-Artist/MintAI/core/internal/testutils"
	"github.com/DeAI-Artist/MintAI/core/internal/testutils/configtest"
)

func Test_SendEveryStrategy(t *testing.T) {
	t.Parallel()

	s := txmgrcommon.SendEveryStrategy{}

	assert.Equal(t, uuid.NullUUID{}, s.Subject())

	ids, err := s.PruneQueue(testutils.Context(t), nil)
	assert.NoError(t, err)
	assert.Len(t, ids, 0)
}

func Test_DropOldestStrategy_Subject(t *testing.T) {
	t.Parallel()
	cfg := configtest.NewGeneralConfig(t, nil)

	subject := uuid.New()
	s := txmgrcommon.NewDropOldestStrategy(subject, 1, cfg.Database().DefaultQueryTimeout())

	assert.True(t, s.Subject().Valid)
	assert.Equal(t, subject, s.Subject().UUID)
}

func Test_DropOldestStrategy_PruneQueue(t *testing.T) {
	t.Parallel()
	cfg := configtest.NewGeneralConfig(t, nil)
	subject := uuid.New()
	queueSize := uint32(2)
	queryTimeout := cfg.Database().DefaultQueryTimeout()
	mockTxStore := mocks.NewEvmTxStore(t)

	t.Run("calls PrineUnstartedTxQueue for the given subject and queueSize, ignoring fromAddress", func(t *testing.T) {
		strategy1 := txmgrcommon.NewDropOldestStrategy(subject, queueSize, queryTimeout)
		mockTxStore.On("PruneUnstartedTxQueue", mock.Anything, queueSize-1, subject, mock.Anything, mock.Anything).Once().Return([]int64{1, 2}, nil)
		ids, err := strategy1.PruneQueue(testutils.Context(t), mockTxStore)
		require.NoError(t, err)
		assert.Equal(t, []int64{1, 2}, ids)
	})
}
