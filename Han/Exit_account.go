package Han

import (
	contract "dapp/Connect"

	// "github.com/ethereum/go-ethereum/common"

	"fmt"

	"github.com/gin-gonic/gin"
)

func Exit_account(c *gin.Context) {
	contract.Userexit()
	password, privatekeyfile := contract.Get()
	fmt.Println("成功退出账户", password, privatekeyfile)
	tohtml(c, privatekeyfile)
}
