package Han

import (
	// "github.com/ethereum/go-ethereum/common"
	mysql "dapp/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CollectOrder(c *gin.Context) {
	Timestamp := c.PostForm("timestamp")
	Account := c.PostForm("account")
	fmt.Println(Timestamp, Account)
	result := mysql.CreateCollectOrder(Account, Timestamp)
	tohtml(c, result)

}
