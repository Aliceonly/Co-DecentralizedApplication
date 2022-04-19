package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func QueryByTime(c *gin.Context) {
	timestamp := c.PostForm("timestamp")
	int, err := strconv.Atoi(timestamp) //转类型
	if err != nil {
		fmt.Println("出错了", err)
	}
	fmt.Println("-------------------------timestamp-------------------", int)
	result := mysql.Query_bytime(int)
	fmt.Println("查询出来的结果是=============>>>>>>>>>>", result)
	c.HTML(http.StatusOK, "result.html", gin.H{
		"task":       result,
		"Taskname":   result.Taskname,
		"Timestamp":  result.Timestamp,
		"Amount":     result.Amount,
		"State":      result.State,
		"LaunchTime": result.LaunchTime,
		"Category":   result.Category,
	})
}
