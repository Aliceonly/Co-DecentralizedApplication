package Han

import (
	// "github.com/ethereum/go-ethereum/common"
	mysql "dapp/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func QueryCollectOrder(c *gin.Context) {
	Account := c.PostForm("Account")
	fmt.Println(Account)
	result := mysql.Query_collect_order(Account)
	tohtml(c, result)

}
