package Han

import (
	contract "dapp/Connect"
	"fmt"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"

	"github.com/gin-gonic/gin"
)

func Cancel(c *gin.Context) {
<<<<<<< HEAD
=======
	// add := common.HexToAddress(c.PostForm("account"))
>>>>>>> b33f0d7e88460ce7bdff55c412e8f250b3ccd9f5
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
