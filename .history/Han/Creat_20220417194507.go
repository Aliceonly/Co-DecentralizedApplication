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
	contract.CreatNewEvent(ins,)
}