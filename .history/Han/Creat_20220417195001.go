package Han

import (
	contract "dapp/Connect"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)
func Creat(c *gin.Context){
	// add :=common.HexToAddress(c.PostForm("account"))
	ins :=contract.Getsmartcontract()
	Txopts :=contract.GetTxopts()
	Taskname=c.PostForm("taskname")
	Tasktime=c.PostForm("tasktime")
	Taskmoney=c.PostForm("taskmoney")
	Taskplace=c.PostForm("taskplace1")
	Taskplace2=c.PostForm("taskplace3")
	Taskcontent=c.postForm("taskcontent")
	fmt.Println(Taskname,Tasktime,Taskmoney,Taskplace,Taskplace2,Taskcontent)
	// contract.CreatNewEvent(ins,Txopts)
}