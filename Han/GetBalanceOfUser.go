package Han

import (
	contract "dapp/Connect"
	"fmt"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

//获取余额
func GetBalanceOfUser(c *gin.Context) {
	ins := contract.Getsmartcontract()
	head, _ := contract.GetBlockNumber()
	_, adress := contract.Getaccout()
	a := contract.GetuserBanlance(ins, adress, head)
	fmt.Println("userbalance:", a)
	tohtml(c, a)
}
