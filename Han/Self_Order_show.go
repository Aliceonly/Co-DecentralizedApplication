package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Self_Order_show(c *gin.Context) {
	Account := c.PostForm("Account")
	Self_taskdata := mysql.Self_Order_show(Account)
	fmt.Println("get____taskdata=======>>>>>>>>>>>>>>", Self_taskdata)
	tohtml(c, Self_taskdata)

}
