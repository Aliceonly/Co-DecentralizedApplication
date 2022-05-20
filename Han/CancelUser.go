package Han

import (
	contract "dapp/Connect"
	mysql "dapp/utils"
	// "github.com/ethereum/go-ethereum/common"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CancelUser(c *gin.Context) {
	add := c.PostForm("account")
	result := contract.Cancellation(add)
	if result == "成功注销" {
		fmt.Println("成功注销账户")
		mysql.CancleUser(add)
		tohtml(c, 1)
	} else {
		tohtml(c, 0)
	}
}
