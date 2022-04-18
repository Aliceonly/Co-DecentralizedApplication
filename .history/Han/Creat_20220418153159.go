package Han

import (
	contract "dapp/Connect"
	"math/big"

	// "github.com/ethereum/go-ethereum/common"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Creat(c *gin.Context) {
	// add :=common.HexToAddress(c.PostForm("account"))
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	head, _ := contract.GetBlockNumber()
	_, adress := contract.Getaccout()
	taskname := c.PostForm("taskname")
	tasktime := c.PostForm("tasktime")
	taskmoney := c.PostForm("taskmoney")
	n := new(big.Int)
	n, ok := n.SetString(taskmoney, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	// fmt.Println(n)
	// taskplace :=c.PostForm("taskplace1")
	taskplace2 := c.PostForm("taskplace3")
	taskplace1 := c.PostForm("taskplace1")
	taskcontent := c.PostForm("taskcontent")
	fmt.Println(taskname, tasktime, taskmoney, taskplace2, taskcontent)
	contract.CreatNewEvent(ins, Txopts, taskname+taskcontent+taskplace1, taskplace2, tasktime, n)
	times := contract.Querytime(ins, adress, head)
	fmt.Println("时间戳----->", times)
	tohtml(c, times)
	fmt.Println(contract.GetTasklist(ins, times, adress, head))
}

func Getbalance