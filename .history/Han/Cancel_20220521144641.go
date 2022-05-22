package Han

import (
	contract "dapp/Connect"

	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
)

func Cancel(c *gin.Context) {
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
	contract.CancelEvent(ins, Txopts, n, add)
	tohtml(c, "canceltrue")
}
