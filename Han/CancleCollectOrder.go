package Han

import (
	// "github.com/ethereum/go-ethereum/common"
	mysql "dapp/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CancleCollectOrder(c *gin.Context) {
	Timestamp := c.PostForm("timestamp")
	fmt.Println(Timestamp)
	mysql.Cancle_CollectOrder(Timestamp)
	tohtml(c, 1)

}
