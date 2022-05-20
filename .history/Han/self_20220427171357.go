package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Self(c *gin.Context) {
	taskdata := mysql.Showdata()
	fmt.Println("get____taskdata=======>>>>>>>>>>>>>>", taskdata)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"data": taskdata,
	})

}

// func all_order_Handler(c *gin.Context) {
// 	c.HTML(200, "self.html", nil) //个人订单
// }
