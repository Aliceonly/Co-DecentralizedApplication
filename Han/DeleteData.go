package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	"strconv"
)

func DeleteData(c *gin.Context) {
	timestamp := c.PostForm("timestamp")
	a, err := strconv.Atoi(timestamp) //转类型
	if err != nil {
		fmt.Println("出错了", err)
	}
	fmt.Println("-------------------------timestamp-------------------", a)
	mysql.DeletTask(a)
}
