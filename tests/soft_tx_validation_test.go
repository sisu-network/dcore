package tests

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mighty-chain/coreth/core/types"
)

func TestSoftValidation(t *testing.T) {
	chain, _, _, _ := NewDefaultChain(t)

	// Mark the genesis block as accepted and start the chain
	chain.Start()
	defer chain.Stop()

	contract := "6080604052348015600f57600080fd5b50602a60008190555060b9806100266000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80631003e2d214602d575b600080fd5b605660048036036020811015604157600080fd5b8101908080359060200190929190505050606c565b6040518082815260200191505060405180910390f35b60008160005401600081905550600054905091905056fea264697066735822122066dad7255aac3ea41858c2a0fe986696876ac85b2bb4e929d2062504c244054964736f6c63430007060033"
	code := common.Hex2Bytes(contract)

	nonce := uint64(0)
	tx := types.NewContractCreation(nonce, big.NewInt(0), uint64(gasLimit), gasPrice, code)
	err := chain.SoftValidateTx(tx, false)
	if err == nil {
		t.Fatal("Soft validation should fail because the transaction is unsigned.")
	}

	// Sign the transaction
	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), fundedKey.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	err = chain.SoftValidateTx(tx, false)
	if err != nil {
		t.Fatal("Soft validation fails")
	}
}
