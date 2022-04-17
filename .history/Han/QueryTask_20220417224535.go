package Han

import (
	contract "dapp/Connect"

	"math/big"
	// "github.com/ethereum/go-ethereum/common"
	"fmt"

	"github.com/gin-gonic/gin"
)
func Query(c *gin.Context){
	ins :=contract.Getsmartcontract()
	head,_:=contract.GetBlockNumber()
	_,adress:=contract.Getaccout()
	times:=big.Int(c.PostForm("timestap")
	n := new(big.Int)
    n, ok := n.SetString(taskmoney, 10)
    if !ok {
        fmt.Println("SetString: error")
        return
    }
	contract.GetTasklist(ins,times,adress,head)
}