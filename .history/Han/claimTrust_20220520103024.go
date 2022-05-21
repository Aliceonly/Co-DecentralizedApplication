package Han

import (
	"github.com/gin-gonic/gin"
	contract "dapp/Connect"

)

func ClaimTrust(c *gin.Context){
	ins := contract.Getsmartcontract()
	contract.ClaimTrust()
  
}