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
	pr, _ := c.Getaccout()
	hash:="0x9a15ca3125764c240946b2eb7fcacadb9bd18fb44bb4355ee6362a07de452f47"
	sigh:="0x9a15ca3125764c240946b2eb7fcacadb9bd18fb44bb4355ee6362a07de452f47"
	Txopts := c.GetTxopts()
	hash := c.Gettaskhash(ins, adress, head, taskname, times)
	// sigh := c.GetthistaskSign(pr, []byte(hash))
	a:=big.NewInt(1653359260)
	
	publice:=pr.Public()
	publicKeyECDSA, ok := publice.(*ecdsa.PublicKey)
    if !ok {
        panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	if c.Validation([]byte(hash),[]byte(sigh),publicKeyBytes) ==true{
		c.ClaimTrust(ins,Txopts,a)
		fmt.Println("ok")
	}else{
		fmt.Println("err")
	}
}

