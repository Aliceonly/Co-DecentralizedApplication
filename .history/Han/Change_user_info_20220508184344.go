package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Change_user_info(c *gin.Context) {
	Account := c.PostForm("account")
	Telephone := c.PostForm("telephone")
	Sid := c.PostForm("sid")
	Sname := c.PostForm("Sname")
	Sage := c.PostForm("Sage")
	Major := c.PostForm("Major")
	Grade := c.PostForm("Grade")
	a, err := strconv.Atoi(Sid) //转类型
	fmt.Println("bbbb--->",Account, a, Sname, Sage, Telephone, Major, Grade)
	if err != nil {
		fmt.Println("出错了", err)
	}
	mysql.Update_User(Account, a, Sname, Sage, Telephone, Major, Grade)
	fmt.Println("查询出来的结果是=============>>>>>>>>>>")
	tohtml(c, 1)

}
