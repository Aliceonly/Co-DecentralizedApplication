package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	"strconv"
)

func Read_more(c *gin.Context) {
	timestamp := c.PostForm("timestamp")
	int, err := strconv.Atoi(timestamp) //转类型
	if err != nil {
		fmt.Println("出错了", err)
	}
	fmt.Println("-------------------------timestamp-------------------", int)
	result := mysql.Query_bytime(int)
	fmt.Println("查询出来的结果是=============>>>>>>>>>>", result)
	tohtml(c, result)
}
