package Han

import (
	contract "dapp/Connect"
	

	"github.com/gin-gonic/gin"
)

func ClaimTrust(c *gin.Context){
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	contract.ClaimTrust(ins,Txopts,)
  
}