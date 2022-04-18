package main

import (

	// c"dapp/Connect"
	// "fmt"
	s"dapp/Han"
	"fmt"
	"dapp/controller"
	"net/http"
	"os/exec"
)


func main() {
    s.Start()
	// ins:=c.Getsmartcontract()
	// head,_:=c.GetBlockNumber()
	// _,adress:=c.Getaccout()
	// fmt.Println("contractBalance",c.GetcontractBanlance(ins,adress,head))

}

func test(){
	//处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//解析html
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	//去首页
	http.HandleFunc("/main", controller.GetHome)
	//邮箱密码登录接口
	http.HandleFunc("/sign_in", controller.ELogin)
	//邮箱验证码登录接口
	http.HandleFunc("/Esign_in", controller.Login)
	//注销接口
	http.HandleFunc("/logout", controller.Logout)
	//注册接口
	http.HandleFunc("/sign_up", controller.Regist)
	//超级管理员注册接口
	http.HandleFunc("/Adregist", controller.RegistAsAdministrator)
	/*
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	http.HandleFunc("/getUserInfo", controller.GetUserInfo)
	http.HandleFunc("/deleteUser", controller.DeleteUser)
	http.HandleFunc("/toUpdateUser", controller.ToUpdateUser)
	http.HandleFunc("/updateOraddUser", controller.UpdateOrAddUser)
	*/



}
