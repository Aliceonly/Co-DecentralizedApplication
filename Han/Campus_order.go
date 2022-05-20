package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Campus_order(c *gin.Context) {
	entry := c.PostForm("entry")
	Self_taskdata := mysql.Campus_order(entry)
	fmt.Println("get____taskdata=======>>>>>>>>>>>>>>", Self_taskdata)
	tohtml(c, Self_taskdata)

}
