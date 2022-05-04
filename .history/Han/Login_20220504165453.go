package Han

import (
	mysql "dapp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	"strconv"
)

func Login(c *gin.Context) {
	Sid := c.PostForm("Sid")
	Password := c.PostForm("Password")
	a, err := strconv.Atoi(Sid) //转类型
	if err != nil {
		fmt.Println("出错了", err)
	}
	fmt.Println("-------------------------data-------------------", a, Password)
	result := mysql.Login(a)
	fmt.Println("密码是=============>>>>>>>>>>", result.Passwd)
	// user_pd := result.Passwd
	// if Password == user_pd {
	// fmt.Println("登陆成功")
	tohtml(c,1)
	// } else {
	// fmt.Print("登陆失败，密码错误")
	// }
	// c.HTML(http.StatusOK, "result.html", gin.H{
	// "Passwd":   result.Passwd,
	// })
}
