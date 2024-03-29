package telemetry

import (
	ocrtypes "github.com/smartcontractkit/libocr/commontypes"

	"github.com/DeAI-Artist/MintAI/core/services/synchronization"
)

type MonitoringEndpointGenerator interface {
	GenMonitoringEndpoint(network string, chainID string, contractID string, telemType synchronization.TelemetryType) ocrtypes.MonitoringEndpoint
}
