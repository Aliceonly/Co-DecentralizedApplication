package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index_Handler(c *gin.Context) {
	c.HTML(200, "index.html", nil) //首页
}

func sign_in_Handler(c *gin.Context) {
	c.HTML(200, "sign-in.html", nil) //登录
}

func sign_up_Handler(c *gin.Context) {
	c.HTML(200, "sign-up.html", nil) //注册
}

func post_job_Handler(c *gin.Context) {
	c.HTML(200, "post-job.html", nil) //发送订单
}

func find_job_Handler(c *gin.Context) {
	c.HTML(200, "find-job.html", nil) //接收订单
}

func email_sign_in_Handler(c *gin.Context) {
	c.HTML(200, "email_sign_in.html", nil) //邮件验证登录
}

func personal_detail_Handler(c *gin.Context) {
	c.HTML(200, "candidate-details.html", nil) //个人信息
}

func detail_campus_Handler(c *gin.Context) {
	c.HTML(200, "details-of-CampusJob.html", nil) //校园兼职工作页面
}

func detail_shared_Handler(c *gin.Context) {
	c.HTML(200, "details-of-SharedService.html", nil) //共享服务工作页面
}

func detail_qukuilySend_Handler(c *gin.Context) {
	c.HTML(200, "details-of-QukuilySend.html", nil) //快送页面
}

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

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
