// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chain

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sisu-network/dcore/accounts/keystore"
	"github.com/sisu-network/dcore/core"
	"github.com/sisu-network/dcore/core/rawdb"
	"github.com/sisu-network/dcore/core/types"
	"github.com/sisu-network/dcore/eth"
	"github.com/sisu-network/dcore/eth/ethconfig"
	"github.com/sisu-network/dcore/params"
)

var (
	basicTxGasLimit       = 21000
	fundedKey, bob, alice *keystore.Key
	initialBalance        = big.NewInt(100000000000000000)
	chainID               = big.NewInt(1)
	value                 = big.NewInt(1000000000000)
	gasLimit              = 10000000
	gasPrice              = big.NewInt(1000000000)
)

func init() {
	genKey, err := keystore.NewKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	fundedKey = genKey
	genKey, err = keystore.NewKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	bob = genKey
	genKey, err = keystore.NewKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	alice = genKey
}

func NewDefaultChain(t *testing.T) (*ETHChain, chan *types.Block, chan core.NewTxPoolHeadEvent, <-chan core.NewTxsEvent) {
	// configure the chain
	config := ethconfig.Defaults
	chainConfig := &params.ChainConfig{
		ChainID:             chainID,
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        big.NewInt(0),
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
	}

	config.Genesis = &core.Genesis{
		Config:     chainConfig,
		Nonce:      0,
		Number:     0,
		ExtraData:  hexutil.MustDecode("0x00"),
		GasLimit:   100000000,
		Difficulty: big.NewInt(0),
		Alloc:      core.GenesisAlloc{fundedKey.Address: {Balance: initialBalance}},
	}

	chain := NewETHChain(&config, nil, rawdb.NewMemoryDatabase(), eth.DefaultSettings, true, nil)

	if err := chain.Accept(chain.GetGenesisBlock()); err != nil {
		t.Fatal(err)
	}

	newBlockChan := make(chan *types.Block)
	chain.SetOnSealFinish(func(block *types.Block) error {
		if err := chain.SetPreference(block); err != nil {
			t.Fatal(err)
		}
		if err := chain.Accept(block); err != nil {
			t.Fatal(err)
		}
		newBlockChan <- block
		return nil
	})

	newTxPoolHeadChan := make(chan core.NewTxPoolHeadEvent, 1)
	chain.GetTxPool().SubscribeNewHeadEvent(newTxPoolHeadChan)

	txSubmitCh := chain.GetTxSubmitCh()
	return chain, newBlockChan, newTxPoolHeadChan, txSubmitCh
}
