package Han

import (
	// contract "dapp/Connect"

	// "math/big"
	// "github.com/ethereum/go-ethereum/common"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)
func Selftask(c *gin.Context){

	mysql
	fmt.Print(userlist)
	tohtml(c,userlist)
}