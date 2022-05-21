package Han

import (
	contract "dapp/Connect"
	

	"github.com/gin-gonic/gin"
)

func ClaimTrust(c *gin.Context){
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	n := new(big.Int)
	n, ok := n.SetString(times, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	contract.ClaimTrust(ins,Txopts,)
  
}