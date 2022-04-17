package Han

import (
	contract "dapp/Connect"
	"math/big"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"fmt"
)
func Creat(c *gin.Context){
	// add :=common.HexToAddress(c.PostForm("account"))
	ins :=contract.Getsmartcontract()
	Txopts :=contract.GetTxopts()
	taskname :=c.PostForm("taskname")
	tasktime :=c.PostForm("tasktime")
	taskmoney :=c.PostForm("taskmoney")
	n := new(big.Int)
    n, ok := n.SetString(taskmoney, 10)
    if !ok {
        fmt.Println("SetString: error")
        return
    }
	fmt.Println(n)
	taskplace :=c.PostForm("taskplace1")
	taskplace2 :=c.PostForm("taskplace3")
	taskcontent :=c.PostForm("taskcontent")
	fmt.Println(taskname,tasktime,taskmoney,taskplace,taskplace2,taskcontent)
	a:=contract.CreatNewEvent(ins,Txopts,taskname+taskcontent+taskplace2,taskplace2,tasktime,n)
	fmt.Print(&a)
}