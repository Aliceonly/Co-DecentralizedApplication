package Han

import (
	contract "dapp/Connect"
	"math/big"
	"strconv"
	"time"

	// "github.com/ethereum/go-ethereum/common"
	mysql "dapp/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
)
type user struct {
	time  *big.Int
	sigh string
}
func Creat(c *gin.Context) {
	mm := make([]user, 0)
	add := c.PostForm("account")
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	head, _ := contract.GetBlockNumber()
	_, adress := contract.Getaccout()
	taskname := c.PostForm("taskname")
	tasktime := c.PostForm("tasktime")
	taskmoney := c.PostForm("taskmoney")
	// taskmoney2:=taskmoney+"000000000000000000"
	n := new(big.Int)
	n, ok := n.SetString(taskmoney, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	taskplace2 := c.PostForm("taskplace3") //工作类型
	taskplace1 := c.PostForm("taskplace1") //地区
	taskcontent := c.PostForm("taskcontent")
	fmt.Println(taskname, tasktime, taskmoney, taskplace2, taskcontent, taskplace1)
	a:=contract.CreatNewEvent(ins, Txopts, taskname+taskcontent+taskplace1, taskplace2, tasktime, n)
	fmt.Println(a.Hash())
	time.Sleep(time.Duration(5)*time.Second)
	for{
		status:=contract.QueryStatus(a.Hash())
		fmt.Println(status)
	   if status == 1{
			break
		}
	}
	times := contract.Querytime(ins, adress, head)
	fmt.Println("时间戳----->", times)
	sigh,hash := Getsigh(times.String(), taskname)
	task_money, _ := strconv.Atoi(taskmoney)
	err := mysql.Insert(taskname, add, taskplace2, task_money, times, "availablev", tasktime, hash)
	if err != nil {
		panic(err)
	}
	// fmt.Println(contract.GetTasklist(ins, times, adress, head))
	user.time=times

	mm = append(mm,usr.times)
	mm =append(mm, )
	tohtml(c,user{time: times,sigh: sigh})
}

func Getsigh(times string,taskname string) (string,string){
	ins := contract.Getsmartcontract()
	head, _ := contract.GetBlockNumber()
	pr, adress := contract.Getaccout()
	hash := contract.Gettaskhash(ins, adress, head, taskname, times)
	fmt.Println("hash====>", hash)
	sigh := contract.GetthistaskSign(pr, hash)
	cc:=hexutil.Encode(sigh)
	aa:=hexutil.Encode(hash[:])
	return cc,aa
}
