package extra

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// Web3API offers helper API methods
type Web3API struct{}

// ClientVersion returns the version of the vm running
func (s *Web3API) ClientVersion() string { return "0.1.0" }

// Sha3 returns the bytes returned by hashing [input] with Keccak256
func (s *Web3API) Sha3(input hexutil.Bytes) hexutil.Bytes { return ethcrypto.Keccak256(input) }
