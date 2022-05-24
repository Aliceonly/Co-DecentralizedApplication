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
	ins := contract.Getsmartcontract()
	head, _ := contract.GetBlockNumber()
	_, adress := contract.Getaccout()
	fmt.Println(usercontract.GetuserBanlance(ins,adress,head))
}