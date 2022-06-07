package Han

import (
	contract "dapp/Connect"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Apply(c *gin.Context) {
	ins := contract.Getsmartcontract2()
	// Txopts := contract.GerTxopts()
	_, adress := contract.Getaccout()
	// a := contract.Dotrasfer(ins, Txopts, adress)
	for {
		status := contract.QueryStatus(a.Hash())
		fmt.Println(status)
		if status == 1 {
			tohtml(c, "ok")
			break
		}
	}

}
