package Han

import (
	contract "dapp/Connect"
	"fmt"
	"math/big"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

func ClaimTrust(c *gin.Context){
	
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	times := c.PostForm("timestap")
	// taskname := c.PostForm("taskname")
	// // phash:=c.PostForm("phash")
	// head, _ := contract.GetBlockNumber()
	// hash:=contract.Gettaskhash(ins,adress,head,taskname,times)
	// var data1 [32]byte = [32]byte(phash)
	n := new(big.Int)
	n, ok := n.SetString(times, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	
	contract.ClaimTrust(ins,Txopts,n)
	tohtml(c,"ClaimOK")
  
}

func Va()  {
	pr, _ := contract.Getaccout()
	publice:=pr.Public()
	publicKeyECDSA, ok := publice.(*ecdsa.PublicKey)
    if !ok {
        panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	contract.Validation(hash,signa)
}