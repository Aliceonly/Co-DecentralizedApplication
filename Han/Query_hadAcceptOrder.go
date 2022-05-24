package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Query_hadAcceptOrder(c *gin.Context) {
	status := c.PostForm("status")
	account := c.PostForm("account")
	data := mysql.Query_hadAcceptOrder(account, status)
	fmt.Println("get____taskdata=======>>>>>>>>>>>>>>", data)
	tohtml(c, data)

}
