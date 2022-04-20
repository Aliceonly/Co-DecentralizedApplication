package Han

import (
	// contract "dapp/Connect"

	// "math/big"
	// "github.com/ethereum/go-ethereum/common"
	"fmt"
    mysql "dapp/utils"
	"github.com/gin-gonic/gin"
)
func Index_Handler(c *gin.Context) {
	// c.HTML(200, "index.html", nil) //首页
	taskdata := mysql.Showdata()
	fmt.Println("get____taskdata=======>>>>>>>>>>>>>>", taskdata)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"data": taskdata,
	}) // -------------------------第三步---------------------

}