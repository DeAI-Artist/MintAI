package services_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/DeAI-Artist/MintAI/core/logger"
	"github.com/DeAI-Artist/MintAI/core/services"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

func TestNewInBackupHealthReport(t *testing.T) {
	lggr, observed := logger.TestLoggerObserved(t, zapcore.InfoLevel)
	ibhr := services.NewInBackupHealthReport(1234, lggr)

	ibhr.Start()
	require.Eventually(t, func() bool { return observed.Len() >= 1 }, time.Second*5, time.Millisecond*100)
	require.Equal(t, "Starting InBackupHealthReport", observed.TakeAll()[0].Message)

	req, err := http.NewRequestWithContext(tests.Context(t), "GET", "http://localhost:1234/health", nil)
	require.NoError(t, err)
	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, res.StatusCode)

	ibhr.Stop()
	require.Eventually(t, func() bool { return observed.Len() >= 1 }, time.Second*5, time.Millisecond*100)
	require.Equal(t, "InBackupHealthReport shutdown complete", observed.TakeAll()[0].Message)
}