package Han

import (
	contract "dapp/Connect"
	"time"

	"github.com/gin-gonic/gin"
)

func ClaimTrust(c *gin.Context){
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	contract.ClaimTrust(ins,Txopts,time)
  
}