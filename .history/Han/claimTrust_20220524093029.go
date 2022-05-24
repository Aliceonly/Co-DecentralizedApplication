package Han

import (
	"crypto/ecdsa"
	contract "dapp/Connect"
	"fmt"
	"hash"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

func ClaimTrust(c *gin.Context){
	
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	times := c.PostForm("timestap")
    hash :=c.PostForm("hash")
	n := new(big.Int)
	n, ok := n.SetString(times, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	va
	contract.ClaimTrust(ins,Txopts,n)
	tohtml(c,"ClaimOK")
  
}

func Va(hash []byte,signature []byte) bool {
	pr, _ := contract.Getaccout()
	publice:=pr.Public()
	publicKeyECDSA, ok := publice.(*ecdsa.PublicKey)
    if !ok {
        panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	a:=contract.Validation(hash,signature,publicKeyBytes)
	return a
}