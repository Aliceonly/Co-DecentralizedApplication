package Han

import (
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
	if Sid == "" {
		tohtml(c, 3)
	}
	if Password == "" && Sid == "" {
		tohtml(c, 4)
	}
	if Password != "" && Sid != "" {
		// a, err := strconv.Atoi(Sid) //转类型
		// if err != nil {
		// 	fmt.Println("出错了", err)
		// }
		fmt.Println("-------------------------data-------------------", a, Password)
		result := mysql.Login(Sid)
		fmt.Println("密码是=============>>>>>>>>>>", result)
		user_pd := result
		if Password == user_pd && user_pd != "" {
			fmt.Println("登陆成功")
			tohtml(c, 1)
		} else {
			fmt.Print("登陆失败，密码错误")
			tohtml(c, 0)
		}
	}

}
