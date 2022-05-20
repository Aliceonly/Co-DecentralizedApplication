package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func User_info(c *gin.Context) {
	Account := c.PostForm("Account")
	fmt.Println("-------------------------Account-------------------", Account)
	result := mysql.Query_User(Account)
	fmt.Println("查询出来的结果是=============>>>>>>>>>>", result)
	tohtml(c, result)
}
