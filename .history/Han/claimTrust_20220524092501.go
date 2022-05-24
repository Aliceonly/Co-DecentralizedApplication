package Han

import (
	contract "dapp/Connect"
	"github.com/gin-gonic/gin"
	"math/big"
	"fmt"
)

func ClaimTrust(c *gin.Context){
	contract.Validation()
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	times := c.PostForm("timestap")
	// taskname := c.PostForm("taskname")
	// // phash:=c.PostForm("phash")
	_, adress := contract.Getaccout()
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