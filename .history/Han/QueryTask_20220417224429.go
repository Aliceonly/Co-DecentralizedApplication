package Han

import (
	contract "dapp/Connect"
	"time"
	// "math/big"
	// "github.com/ethereum/go-ethereum/common"
	"fmt"

	"github.com/gin-gonic/gin"
)
func Query(c *gin.Context){
	ins :=contract.Getsmartcontract()
	Txopts :=contract.GetTxopts()
	head,_:=contract.GetBlockNumber()
	_,adress:=contract.Getaccout()
	times=
	contract.GetTasklist(ins,times,adress,head)
}