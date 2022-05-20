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
	c.HTML(http.StatusOK, "se.html", gin.H{
		"data": taskdata,
	})

}