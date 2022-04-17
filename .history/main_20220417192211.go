package main

import (
	"net/http"
	c"dapp/Connect"
	"github.com/gin-gonic/gin"
	"fmt"
)


func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置模板
	r.LoadHTMLGlob("html/*")
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	r.StaticFS("/static", http.Dir("./static"))
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数

	r.GET("/", index_Handler)
	r.GET("/sign_in", sign_in_Handler)
	r.GET("/sign_up", sign_up_Handler)
	r.GET("/post_job", post_job_Handler)
	r.GET("/find_job", find_job_Handler)
	r.GET("/email_sign_in", email_sign_in_Handler)
	r.GET("/personal_detail", personal_detail_Handler)
	r.GET("/detail_campus", detail_campus_Handler)
	r.GET("/detail_qukuilySend", detail_qukuilySend_Handler)
	r.GET("/detail_shared", detail_shared_Handler)
	r.GET("/redact_candidate", redact_candidate_Handler)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	ins:=c.Getsmartcontract()
	head,_:=c.GetBlockNumber()
	_,adress:=c.Getaccout()
	fmt.Println("contractBalance",c.GetcontractBanlance(ins,adress,head))
	r.Run()
}
