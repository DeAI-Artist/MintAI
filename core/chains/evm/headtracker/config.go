package headtracker

import (
	"time"

	"github.com/DeAI-Artist/MintAI/core/chains/evm/config"
)

//go:generate mockery --quiet --name Config --output ./mocks/ --case=underscore

// Config represents a subset of options needed by head tracker
type Config interface {
	BlockEmissionIdleWarningThreshold() time.Duration
	FinalityDepth() uint32
}

type HeadTrackerConfig interface {
	config.HeadTracker
}
