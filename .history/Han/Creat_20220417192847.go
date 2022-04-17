package Han
import (
	contract"dapp/Connect"
	"github.com/gin-gonic/gin"
)
func Creat(c *gin.Context){
	add :=common.HexToAddress(c.PostForm("account"))
	ins :=contract.Getsmartcontract()
	Txopts :=contract.GetTxopts()
}