package Han

import (
	contract "dapp/Connect"
	"math/big"

	// "github.com/ethereum/go-ethereum/common"
	"fmt"
    mysql"dapp/utils"
	"github.com/gin-gonic/gin"
)

func Creat(c *gin.Context) {
	add :=c.PostForm("account")
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
	taskplace2 := c.PostForm("taskplace3")//工作类型
	taskplace1 := c.PostForm("taskplace1")//地区
	taskcontent := c.PostForm("taskcontent")
	fmt.Println(taskname, tasktime, taskmoney, taskplace2, taskcontent,taskplace1)
	contract.CreatNewEvent(ins, Txopts, taskname+taskcontent+taskplace1, taskplace2, tasktime, n)
	times := contract.Querytime(ins, adress, head)
	fmt.Println("时间戳----->", times)
	tohtml(c, times)
	Getbalance()
	mysql.Insert(taskname,add,taskplace2,taskmoney,times,)
	fmt.Println(contract.GetTasklist(ins, times, adress, head))
}

func Getbalance(){
	ins:=contract.Getsmartcontract()
	head,_:=contract.GetBlockNumber()
	_,adress:=contract.Getaccout()
	fmt.Println("contractBalance",contract.GetcontractBanlance(ins,adress,head))
}