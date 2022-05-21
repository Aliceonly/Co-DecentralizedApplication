package Han

import (
	contract "dapp/Connect"
	"github.com/gin-gonic/gin"
	"math/big"
	"fmt"
)

func ClaimTrust(c *gin.Context){
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	times := c.PostForm("timestap")
	sign:=c.PostForm("Sign")
	var data []byte = []byte(sign)
	var data1 [3]byte = []byte(sign)
	taskname := c.PostForm("taskname")
	phash:=c.PostForm("phash")
	n := new(big.Int)
	n, ok := n.SetString(times, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	contract.ClaimTrust(ins,Txopts,n,data,phash,taskname)
  
}