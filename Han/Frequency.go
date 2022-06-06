package Han

import (
	contract "dapp/Connect"

	"github.com/gin-gonic/gin"
)

func Frequency(c *gin.Context) {
	ins := contract.Getsmartcontract2()
	_, adress := contract.Getaccout()
	head, _ := contract.GetBlockNumber()
	a := contract.Getfrequency(ins, adress, head)
	tohtml(c, a)
}
