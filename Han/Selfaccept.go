package Han

import (
	// contract "dapp/Connect"

	// "math/big"
	// "github.com/ethereum/go-ethereum/common"
	"fmt"
    mysql "dapp/utils"
	"github.com/gin-gonic/gin"
)
func Selfaccept(c *gin.Context){
    add :=c.PostForm("account")
	selftask_list:=mysql.Selectaccept(add)
	fmt.Print(selftask_list)
	tohtml(c,selftask_list)
}