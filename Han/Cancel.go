package Han

import (
	contract "dapp/Connect"

	"fmt"
	"math/big"

	"github.com/gin-gonic/gin"
)

func Cancel(c *gin.Context) {
	// add := common.HexToAddress(c.PostForm("account"))
	_, adress := contract.Getaccout()
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	times := c.PostForm("timestap")
	n := new(big.Int)
	n, ok := n.SetString(times, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	contract.CancelEvent(ins, Txopts, n, adress)
	tohtml(c, "canceltrue")
}
