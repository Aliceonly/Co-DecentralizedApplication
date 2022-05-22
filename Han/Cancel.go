package Han

import (
	contract "dapp/Connect"
	"fmt"
	// "github.com/ethereum/go-ethereum/common"
	mysql "dapp/utils"
	"github.com/gin-gonic/gin"
	"math/big"
)

func Cancel(c *gin.Context) {
	state := "invalid"
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
	mysql.Update_state(state, times)
}
