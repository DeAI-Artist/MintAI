package config

import (
	"github.com/DeAI-Artist/MintAI/core/services/keystore/keys/p2pkey"
)

type P2P interface {
	V2() V2
	PeerID() p2pkey.PeerID
	IncomingMessageBufferSize() int
	OutgoingMessageBufferSize() int
	TraceLogging() bool
	Enabled() bool
}
