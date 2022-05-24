package Han
import (
	contract "dapp/Connect"
	"fmt"
	// "github.com/ethereum/go-ethereum/common"
	mysql "dapp/utils"
	"github.com/gin-gonic/gin"
	"math/big"
)

func GetBalanceOfUser(c *gin.Context){
	ins := c.Getsmartcontract()
	head, _ := c.GetBlockNumber()
	_, adress := c.Getaccout()
	fmt.Println(c.GetuserBanlance(ins,adress,head))
}