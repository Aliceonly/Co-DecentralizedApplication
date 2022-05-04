package Han

import (
	contract "dapp/Connect"
	// "math/big"
	"strconv"
	// "github.com/ethereum/go-ethereum/common"
	mysql "dapp/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	sid := c.PostForm("Sid")
	tele := c.PostForm("Telephone")
	pd := c.PostForm("Password")
	account := contract.CreatnewActogeth(pd)
	fmt.Println("--------------------create_user_data------------------", sid, tele, pd, account)
	//数据库操作
	sid_1, _ := strconv.Atoi(sid) //转类型
	tele_1, _ := strconv.Atoi(tele)
	mysql.CreateUser(sid_1, tele_1, pd, account)
}
