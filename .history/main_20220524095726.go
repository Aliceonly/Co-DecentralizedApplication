package main

import (
	c "dapp/Connect"
	"fmt"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	// "fmt"
	// s "dapp/Han"
	// "github.com/ethereum/go-ethereum/common"
	// "dapp/controller"
	// "fmt"
	// "math/big"
	// "net/http"
	// "github.com/go-sql-driver/mysql"
	// mysql "dapp/utils"
	// "io/ioutil"
	// "os"
	// "github.com/lithammer/fuzzysearch/fuzzy"
)

func main() {
	// s.Start()
	ins := c.Getsmartcontract()
	head, _ := c.GetBlockNumber()
	pr, adress := c.Getaccout()
	taskname:="a校园区域主要建筑"
	times:="1653357157"
	Txopts := c.GetTxopts()
	hash := c.Gettaskhash(ins, adress, head, taskname, times)
	sigh := c.GetthistaskSign(pr, hash)
	// a:=big.NewInt(1653212002)
	// c.ClaimTrust(ins,Txopts,a,sigh,hash,taskname)
	n := new(big.Int)
	n, ok := n.SetString(taskmoney, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	publice:=pr.Public()
	publicKeyECDSA, ok := publice.(*ecdsa.PublicKey)
    if !ok {
        panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	if c.Validation(hash[:],sigh,publicKeyBytes) ==true{
		c.ClaimTrust(ins,Txopts,times)
	}
	fmt.Println()

}

