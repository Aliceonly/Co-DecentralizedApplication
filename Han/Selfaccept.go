package Han

import (
	// contract "dapp/Connect"

	// "math/big"
	// "github.com/ethereum/go-ethereum/common"
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Selfaccept(c *gin.Context) {
	add := c.PostForm("account")
	selftask_list := mysql.Query_selfacceptOrder(add)
	fmt.Print(selftask_list)
	tohtml(c, selftask_list)
}
