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
	fmt.Println("新创建的用户为------>",account)
	fmt.Println("--------------------create_user_data------------------", sid, tele, pd, account)
	
	//数据库操作
	sid_1, _ := strconv.Atoi(sid) //转类型
	mysql.CreateUser(sid_1, tele, pd, account)
	tohtml(c, account)

}
