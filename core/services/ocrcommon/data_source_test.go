package ocrcommon_test

import (
	"math/big"
	"testing"

	promtestutil "github.com/prometheus/client_golang/prometheus/testutil"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/DeAI-Artist/MintAI/core/internal/testutils"
	"github.com/DeAI-Artist/MintAI/core/logger"
	"github.com/DeAI-Artist/MintAI/core/services/job"
	"github.com/DeAI-Artist/MintAI/core/services/ocrcommon"
	"github.com/DeAI-Artist/MintAI/core/services/pipeline"
	pipelinemocks "github.com/DeAI-Artist/MintAI/core/services/pipeline/mocks"
)

var (
	mockValue          = "100000000"
	jsonParseTaskValue = "1234"
)

func Test_InMemoryDataSource(t *testing.T) {
	runner := pipelinemocks.NewRunner(t)
	runner.On("ExecuteRun", mock.Anything, mock.AnythingOfType("pipeline.Spec"), mock.Anything, mock.Anything).
		Return(&pipeline.Run{}, pipeline.TaskRunResults{
			{
				Result: pipeline.Result{
					Value: mockValue,
					Error: nil,
				},
				Task: &pipeline.HTTPTask{},
			},
		}, nil)

	ds := ocrcommon.NewInMemoryDataSource(runner, job.Job{}, pipeline.Spec{}, logger.TestLogger(t))
	val, err := ds.Observe(testutils.Context(t), types.ReportTimestamp{})
	require.NoError(t, err)
	assert.Equal(t, mockValue, val.String()) // returns expected value after pipeline run
}

func Test_InMemoryDataSourceWithProm(t *testing.T) {
	runner := pipelinemocks.NewRunner(t)

	jsonParseTask := pipeline.JSONParseTask{
		BaseTask: pipeline.BaseTask{},
	}
	bridgeTask := pipeline.BridgeTask{
		BaseTask: pipeline.BaseTask{},
	}

	bridgeTask.BaseTask = pipeline.NewBaseTask(1, "ds1", []pipeline.TaskDependency{{
		PropagateResult: true,
		InputTask:       nil,
	}}, []pipeline.Task{&jsonParseTask}, 1)

	jsonParseTask.BaseTask = pipeline.NewBaseTask(2, "ds1_parse", []pipeline.TaskDependency{{
		PropagateResult: false,
		InputTask:       &bridgeTask,
	}}, []pipeline.Task{}, 2)

	runner.On("ExecuteRun", mock.Anything, mock.AnythingOfType("pipeline.Spec"), mock.Anything, mock.Anything).
		Return(&pipeline.Run{}, pipeline.TaskRunResults([]pipeline.TaskRunResult{
			{
				Task:   &bridgeTask,
				Result: pipeline.Result{},
			},
			{
				Result: pipeline.Result{Value: jsonParseTaskValue},
				Task:   &jsonParseTask,
			},
		}), nil)

	ds := ocrcommon.NewInMemoryDataSource(
		runner,
		job.Job{
			Type: "offchainreporting",
		},
		pipeline.Spec{},
		logger.TestLogger(t),
	)
	val, err := ds.Observe(testutils.Context(t), types.ReportTimestamp{})
	require.NoError(t, err)

	assert.Equal(t, jsonParseTaskValue, val.String()) // returns expected value after pipeline run
	assert.Equal(t, cast.ToFloat64(jsonParseTaskValue), promtestutil.ToFloat64(ocrcommon.PromOcrMedianValues))
	assert.Equal(t, cast.ToFloat64(jsonParseTaskValue), promtestutil.ToFloat64(ocrcommon.PromBridgeJsonParseValues))

}

type mockSaver struct {
	r *pipeline.Run
}

func (ms *mockSaver) Save(r *pipeline.Run) {
	ms.r = r
}

func Test_NewDataSourceV2(t *testing.T) {
	runner := pipelinemocks.NewRunner(t)
	ms := &mockSaver{}
	runner.On("ExecuteRun", mock.Anything, mock.AnythingOfType("pipeline.Spec"), mock.Anything, mock.Anything).
		Return(&pipeline.Run{}, pipeline.TaskRunResults{
			{
				Result: pipeline.Result{
					Value: mockValue,
					Error: nil,
				},
				Task: &pipeline.HTTPTask{},
			},
		}, nil)

	ds := ocrcommon.NewDataSourceV2(runner, job.Job{}, pipeline.Spec{}, logger.TestLogger(t), ms, nil)
	val, err := ds.Observe(testutils.Context(t), types.ReportTimestamp{})
	require.NoError(t, err)
	assert.Equal(t, mockValue, val.String()) // returns expected value after pipeline run
	assert.Equal(t, &pipeline.Run{}, ms.r)   // expected data properly passed to channel
}

func Test_NewDataSourceV1(t *testing.T) {
	runner := pipelinemocks.NewRunner(t)
	ms := &mockSaver{}
	runner.On("ExecuteRun", mock.Anything, mock.AnythingOfType("pipeline.Spec"), mock.Anything, mock.Anything).
		Return(&pipeline.Run{}, pipeline.TaskRunResults{
			{
				Result: pipeline.Result{
					Value: mockValue,
					Error: nil,
				},
				Task: &pipeline.HTTPTask{},
			},
		}, nil)

	ds := ocrcommon.NewDataSourceV1(runner, job.Job{}, pipeline.Spec{}, logger.TestLogger(t), ms, nil)
	val, err := ds.Observe(testutils.Context(t), ocrtypes.ReportTimestamp{})
	require.NoError(t, err)
	assert.Equal(t, mockValue, new(big.Int).Set(val).String()) // returns expected value after pipeline run
	assert.Equal(t, &pipeline.Run{}, ms.r)                     // expected data properly passed to channel
}
