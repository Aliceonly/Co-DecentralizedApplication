package Han

import (
	contract "dapp/Connect"
	// "github.com/ethereum/go-ethereum/common"

	"fmt"

	"github.com/gin-gonic/gin"
)

func Succ_Login(c *gin.Context) {
	password, privatekeyfile := contract.Get()
	fmt.Print(password, privatekeyfile)
	tohtml(c, privatekeyfile)
}
