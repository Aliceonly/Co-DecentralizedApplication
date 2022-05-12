package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Self_Order_show(c *gin.Context) {
	Account := c.PostForm("Account")
	Self_taskdata := mysql.Self_Order_show(Account)
	fmt.Println("get____taskdata=======>>>>>>>>>>>>>>", Self_taskdata)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"data": Self_taskdata,
	})

}
