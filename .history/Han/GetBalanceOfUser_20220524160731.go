package Han
import (
	contract "dapp/Connect"
	"fmt"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GetBalanceOfUser(c *gin.Context){
	ins := contract.Getsmartcontract()
	head, _ := contract.GetBlockNumber()
	_, adress := contract.Getaccout()
	contract.GetuserBanlance(ins,adress,head)
	fmt.Println()
}