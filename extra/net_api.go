package extra

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type NetAPI struct {
	NetworkId string
}

// Listening returns an indication if the node is listening for network connections.
func (s *NetAPI) Listening() bool { return true } // always listening

// PeerCount returns the number of connected peers
func (s *NetAPI) PeerCount() hexutil.Uint { return hexutil.Uint(0) } // TODO: report number of connected peers

// Version returns the current ethereum protocol version.
func (s *NetAPI) Version() string { return s.NetworkId }
