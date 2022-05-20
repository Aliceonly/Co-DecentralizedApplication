package main

import (
	// c"dapp/Connect"
	"fmt"
	// s "dapp/Han"
	// "dapp/controller"
	// "fmt"
	// "math/big"
	// "net/http"
	// "github.com/go-sql-driver/mysql"
	// mysql "dapp/utils"
	"io/ioutil"
	"os"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

func main() {
	// s.Start()
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
	test2()

}
func test2(){
	var FileInfo []os.FileInfo
	var err error
	relativePath := "D:\\y\\geth\\node1\\nodedata\\keystore"
	
	if FileInfo,err = ioutil.ReadDir( relativePath ); err != nil{
		fmt.Println("读取 img 文件夹出错")
		return
	}
  var a:=make([]string,0)
	for _,fileInfo := range FileInfo {
		a = append(a,fileInfo.Name())
		// fmt.Println(fileInfo.Name())
	}
	// var b = []string{"UTC--2021-09-12T17-06-06.881126000Z--00dc6e8b60fa02a5d83e525bbef3240e8ea54dc5", "UTC--2022-05-02T08-26-11.435276500Z--5c595872e02b0613658036bdf5daa6d9f42954be"}
	ac:="5c595872e02b0613658036bdf5daa6d9f42954be"
	matches1 := fuzzy.Find(ac, b)
	print(matches1[0])

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
