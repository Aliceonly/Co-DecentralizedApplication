package Han

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sign_in_Handler(c *gin.Context) {
	// c.HTML(200, "sign-in.html", nil) //登录
	c.HTML(200, "detail-order.html", nil)
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

func redact_candidate_Handler(c *gin.Context) {
	c.HTML(200, "redact-canditate.html", nil) //编辑资料
}
func detail_Handler(c *gin.Context) {
	c.HTML(200, "detail-order.html", nil) //编辑资料
}
func detail_order_Handler(c *gin.Context) {
	c.HTML(200, "detail-order.html", nil) //订单详情
}
func accept_order_Handler(c *gin.Context) {
	c.HTML(200, "accept-order.html", nil) //接单详情
}
func all_order_Handler(c *gin.Context) {
	c.HTML(200, ".html", nil) //个人订单
}
func under_order_Handler(c *gin.Context) {
	c.HTML(200, "under-order.html", nil) //未完成订单
}
func complete_order_Handler(c *gin.Context) {
	c.HTML(200, "complete-order.html", nil) //已完成
}
func undone_order_Handler(c *gin.Context) {
	c.HTML(200, "undone-order.html", nil) //已失效
}
func publish_orser_Handler(c *gin.Context) {
	c.HTML(200, "publish-order.html", nil) //成功发布
}

func Start() error {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置模板
	r.LoadHTMLGlob("html/*")
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	r.StaticFS("/static", http.Dir("./static"))
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数

	r.GET("/", Showdata)
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
	r.GET("/Detail", detail_Handler)
	r.POST("/QueryByTime", QueryByTime)
	r.GET("/detail_order", detail_order_Handler)
	r.GET("/accept_order", accept_order_Handler)
	// r.GET("/all_order", all_order_Handler)
	r.GET("/under_order", under_order_Handler)
	r.GET("/complete_order", complete_order_Handler)
	r.GET("/undone_order", undone_order_Handler)
	r.GET("/publish_order", publish_orser_Handler)
	r.GET("/all_order", Self)

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务

	//设置路由组
	dapp := r.Group("dapp")
	{
		dapp.POST("/creatTask", Creat)            //创建任务接口
		dapp.POST("/querytasklist", Query)        //调用合约直接查任务mapping方法
		dapp.POST("/queryuselist", Queryuser)     //调用合约直接查用户mapping方法
		dapp.POST("/canceltask", Cancel)          //发布者取消任务接口
		dapp.POST("/queryselftask", Selftask)     //发布者查询个人发布的任务的数据
		dapp.POST("/confirmtask", Confirmtask)    //接受任务接口
		dapp.POST("/quertselfaccept", Selfaccept) //发布者查询个人接收的任务的数据
		dapp.POST("/detailData", DetailData)
		dapp.POST("/delete", DeleteData)
	}

	err := r.Run()
	return err
}

func tohtml(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

func returnError(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"code":    1,
		"message": msg,
	})
}
