package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Query_Dim_order(c *gin.Context) {
	name := c.PostForm("name")
	new_name := "%" + name + "%"
	a := mysql.Query_Dim_Order(new_name)
	fmt.Println(new_name)
	fmt.Println(a)
}
