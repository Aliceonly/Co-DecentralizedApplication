package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Unaccept_Order_show(c *gin.Context) {
	Account := c.PostForm("Account")
	state := "availablev"
	data := mysql.Query_unget_order(Account, state)
	fmt.Println("get____taskdata=======>>>>>>>>>>>>>>", data)
	tohtml(c, data)

}
