package Han
import (
	contract"dapp/Connect"
	"github.com/gin-gonic/gin"
	"github.com/ethereum/go-ethereum/common"
)
func Creat(c *gin.Context){
	add :=common.HexToAddress(c.PostForm("account"))
	ins :=contract.Getsmartcontract()
	Txopts :=contract.GetTxopts()
 data:{taskname:taskname,tasktime:tasktime+tasktime2,account:account,taskmoney:taskmoney,taskplace1:taskplace1+taskplace2,taskplace3:taskplace3,taskcontent:taskcontent},
	Taskname=c.PostForm("taskname")
	Tasktime=c.PostForm("tasktime")
	Taskmoney=c.PostForm("")
	contract.CreatNewEvent(ins,Txopts)
}