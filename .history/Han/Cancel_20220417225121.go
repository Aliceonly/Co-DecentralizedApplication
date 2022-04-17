package Han

import (
	contract "dapp/Connect"

	"math/big"
	// "github.com/ethereum/go-ethereum/common"
	"fmt"

	"github.com/gin-gonic/gin"
)
func Cancel(c *gin.Context){
	ins :=contract.Getsmartcontract()
	Txopts :=contract.GetTxopts()
	times:=c.PostForm("timestap")
	n := new(big.Int)
    n, ok := n.SetString(times, 10)
    if !ok {
        fmt.Println("SetString: error")
        return
    }
	userlist:=contract.Getuser(ins,n,adress,head)
	fmt.Print(userlist)
	tohtml(c,userlist)
}