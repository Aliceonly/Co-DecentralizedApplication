package main

import (
	"crypto/ecdsa"
	c "dapp/Connect"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	// // "fmt"
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
	taskname:="dc校园区域主要建筑"
	times:="1653359260"
	// hash:="0x9a15ca3125764c240946b2eb7fcacadb9bd18fb44bb4355ee6362a07de452f47"
	sigh:="0x9a20a166ff4166647eaceff9a9341127403788f1d1ccd96ce07391d725f179a04fdf1f6a705ae6d0c64de264d0880f39e2252c62ac700394554733a8783aa2360135"
	Txopts := c.GetTxopts()
	hash := c.Gettaskhash(ins, adress, head, taskname, times)
	kk := c.GetthistaskSign(pr, hash)
	a:=big.NewInt(1653359260)
	// c.ClaimTrust(ins,Txopts,a,sigh,hash,taskname)
	
	publice:=pr.Public()
	publicKeyECDSA, ok := publice.(*ecdsa.PublicKey)
    if !ok {
        panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	if c.Validation(hash[:],[]byte(sigh),publicKeyBytes) ==true{
		c.ClaimTrust(ins,Txopts,a)
		fmt.Println("ok")
	}else{
		fmt.Println("err")
	}
}

