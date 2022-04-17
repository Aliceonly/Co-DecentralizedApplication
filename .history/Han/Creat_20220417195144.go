package Han

import (
	// contract "dapp/Connect"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)
func Creat(c *gin.Context){
	// add :=common.HexToAddress(c.PostForm("account"))
	// ins :=contract.Getsmartcontract()
	// Txopts :=contract.GetTxopts()
	taskname :=c.PostForm("taskname")
	tasktime :=c.PostForm("tasktime")
	taskmoney:=c.PostForm("taskmoney")
	taskplace:=c.PostForm("taskplace1")
	taskplace2=c.PostForm("taskplace3")
	taskcontent=c.postForm("taskcontent")
	fmt.Println(taskname,Tasktime,Taskmoney,Taskplace,Taskplace2,Taskcontent)
	// contract.CreatNewEvent(ins,Txopts)
}