package Han

import (
	contract "dapp/Connect"

	mysql "dapp/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
)

func Confirmtask(c *gin.Context) {
	state := "已接单"
	account := c.PostForm("account")
	add := common.HexToAddress(account)
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	times := c.PostForm("timestap")
	n := new(big.Int)
	n, ok := n.SetString(times, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	contract.Confirmtask(ins, n, Txopts, add)
	mysql.Update_beneficiary(times, state, account)
	tohtml(c, "confirmtrue")
}
