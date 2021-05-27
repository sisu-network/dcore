package main

import (
	"github.com/sisu-network/dcore/chain"
	"github.com/sisu-network/dcore/eth"
)

func main() {
	chain.NewETHChain(nil, nil, nil, eth.Settings{}, false, nil)
}
