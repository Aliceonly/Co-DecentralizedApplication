package main

import (
	c"dapp/Connect"
	// "fmt"
	s "dapp/Han"
	// "dapp/controller"
	// "fmt"
	// "math/big"
	// "net/http"
	// "github.com/go-sql-driver/mysql"
	// mysql "dapp/utils"
	// "io/ioutil"
	// "os"
	// "github.com/lithammer/fuzzysearch/fuzzy"
)

func main() {
	s.Start()
	privc.Getaccout()
	// c.Getblockmessage()
	// a:="0x5e9ed07526dd85415af6a2f367a66efb02637ba0"
	// fmt.Println(a[2:])
	// c.Changeuser("aebdbdb1455a7d56d56fe2cf967fbe7118845e66","abc")
	// a,b:=c.Get()
	// print(a,b)
	// a:=c.CreatnewActogeth("123")
	// print("newaccount",a)
	// mysql.DeletTask(1650442602)
	// s.Creat()
	// big:=big.NewInt(2)
	// err:=mysql.Insert("aa","0xxx","a",2,big,"availablev","200202")
	// if err!=nil {
	// 	panic(err)
	// }
	// fmt.Println(mysql.Select("0xxx"))
	// fmt.Print(mysql.Selectaccept("0xxx000"))


}


// func test() {
// 	//处理静态资源
// 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
// 	//解析html
// 	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
// 	//去首页
// 	http.HandleFunc("/main", controller.GetHome)
// 	//邮箱密码登录接口
// 	http.HandleFunc("/sign_in", controller.ELogin)
// 	//邮箱验证码登录接口
// 	http.HandleFunc("/Esign_in", controller.Login)
// 	//注销接口
// 	http.HandleFunc("/logout", controller.Logout)
// 	//注册接口
// 	http.HandleFunc("/sign_up", controller.Regist)
// 	//超级管理员注册接口
// 	http.HandleFunc("/Adregist", controller.RegistAsAdministrator)
// 	/*
// 		http.HandleFunc("/checkUserName", controller.CheckUserName)
// 		http.HandleFunc("/getUserInfo", controller.GetUserInfo)
// 		http.HandleFunc("/deleteUser", controller.DeleteUser)
// 		http.HandleFunc("/toUpdateUser", controller.ToUpdateUser)
// 		http.HandleFunc("/updateOraddUser", controller.UpdateOrAddUser)
// 	*/

// 	http.ListenAndServe(":8080", nil)
// }
