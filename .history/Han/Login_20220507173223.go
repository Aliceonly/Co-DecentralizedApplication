package Han

import (
	contract "dapp/Connect"
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	// "strconv"
)

func Login(c *gin.Context) {
	Account := c.PostForm("Account")
	Password := c.PostForm("Password")
	if Password == "" {
		tohtml(c, 2)
	}
	if Account == "" {
		tohtml(c, 3)
	}
	if Password == "" && Account == "" {
		tohtml(c, 4)
	}
	if Password != "" && Account != "" {
		fmt.Println("-------------------------data-------------------", Account, Password)
		result := mysql.Login(Account)
		fmt.Println("密码是=============>>>>>>>>>>", result)
		user_pd := result
		if Password == user_pd && user_pd != "" {
			fmt.Println("登陆成功")
			tohtml(c, 1)
			contract.Changeuser(Account,)
		} else {
			fmt.Print("登陆失败，密码错误")
			tohtml(c, 0)
		}
	}

}
