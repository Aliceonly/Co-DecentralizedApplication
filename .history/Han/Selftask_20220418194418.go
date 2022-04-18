package Han

import (
	// contract "dapp/Connect"

	// "math/big"
	// "github.com/ethereum/go-ethereum/common"
	"fmt"
    mysql "dapp/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)
func Selftask(c *gin.Context){

    add :=c.PostForm("account")
	selftask_list:=mysql.Select(add)
	fmt.Print(userlist)
	tohtml(c,userlist)
}