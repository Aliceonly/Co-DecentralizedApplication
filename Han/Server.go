package Han

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sign_in_Handler(c *gin.Context) {
	// c.HTML(200, "sign-in.html", nil) //登录
	c.HTML(200, "sign-in.html", nil)
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
	c.HTML(200, "CWork_order.html", nil) //校园兼职工作页面
}

func detail_shared_Handler(c *gin.Context) {
	c.HTML(200, "Shared_order.html", nil) //共享服务工作页面
}

func detail_qukuilySend_Handler(c *gin.Context) {
	c.HTML(200, "Campus_order.html", nil) //快送页面
}

func redact_candidate_Handler(c *gin.Context) {
	c.HTML(200, "redact-canditate.html", nil) //编辑资料
}
func detail_Handler(c *gin.Context) {
	c.HTML(200, "job-details.html", nil) //编辑资料
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
func delete_succ(c *gin.Context) {
	c.HTML(200, "del_succ.html", nil) //成功发布
}

func create_succ(c *gin.Context) {
	c.HTML(200, "create_succ.html", nil) //成功发布
}

func user_Handler(c *gin.Context) {
	c.HTML(200, "account.html", nil) //成功发布
}

func resume_Handler(c *gin.Context) {
	c.HTML(200, "resume.html", nil) //成功发布
}

func change_user_Handler(c *gin.Context) {
	c.HTML(200, "change_user.html", nil) //成功发布
}

func Self_Order_Handler(c *gin.Context) {
	c.HTML(200, "Self_Order.html", nil) //成功发布
}

func self_order_accept_Handler(c *gin.Context) {
	c.HTML(200, "self_order_accept.html", nil) //成功发布
}

func Release_order_Handler(c *gin.Context) {
	c.HTML(200, "release_order.html", nil)
}

func Read_More_Handler(c *gin.Context) {
	c.HTML(200, "job_more.html", nil) //编辑资料
}

func Unaccept_order_show_handler(c *gin.Context) {
	c.HTML(200, "Unaccept_order_show.html", nil) //编辑资料
}

func Campus_order_Read_more_handler(c *gin.Context) {
	c.HTML(200, "Campus_order_Read_more.html", nil) //编辑资料
}

func CWork_order_Read_more_handler(c *gin.Context) {
	c.HTML(200, "CWork_order_Read_more.html", nil) //编辑资料
}

func Shared_order_Read_more_handler(c *gin.Context) {
	c.HTML(200, "Shared_order_Read_more.html", nil) //编辑资料
}

func Self_info_handler(c *gin.Context) {
	c.HTML(200, "self_info.html", nil) //个人信息
}

func self_new_order_handler(c *gin.Context) {
	c.HTML(200, "self_new_order.html", nil) //个人订单
}

func setting_handler(c *gin.Context) {
	c.HTML(200, "setting.html", nil) //设置
}

func self_accept_order_handler(c *gin.Context) {
	c.HTML(200, "self_accept_order.html", nil) //设置
}

func self_resume_order_handler(c *gin.Context) {
	c.HTML(200, "self_resume_order.html", nil) //设置
}

func quertselfaccept_handler(c *gin.Context) {
	c.HTML(200, "quertselfaccept.html", nil) //查看个人接收订单
}

func collect_order_handler(c *gin.Context) {
	c.HTML(200, "collect_order.html", nil) //查看个人收藏订单
}

func Collect_order_more_Handler(c *gin.Context) {
	c.HTML(200, "collect_order_more.html", nil) //查看个人收藏订单更多情况
}

func GetBalance_show_Handler(c *gin.Context) {
	c.HTML(200, "GetBalance.html", nil) //查看余额
}

func result_hadAcceptOrder_Handler(c *gin.Context) {
	c.HTML(200, "HadAcceptOrder.html", nil) //查看个人已接单订单
}

func Apply_Handler(c *gin.Context) {
	c.HTML(200, "Apply.html", nil) //查看个人已接单订单
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
	r.GET("/detail_order", detail_order_Handler)
	r.GET("/accept_order", accept_order_Handler)
	// r.GET("/all_order", all_order_Handler)
	r.GET("/under_order", under_order_Handler)
	r.GET("/complete_order", complete_order_Handler)
	r.GET("/undone_order", undone_order_Handler)
	r.GET("/publish_order", publish_orser_Handler)
	r.GET("/account", user_Handler)
	r.GET("/resume", resume_Handler)
	r.GET("/change_user", change_user_Handler)
	r.GET("/self_order", Self_Order_Handler)
	r.GET("/self_order_accept", self_order_accept_Handler)           //个人订单
	r.GET("/release_order", Release_order_Handler)                   //查看个人发布订单
	r.GET("/Read_More_detail", Read_More_Handler)                    //读取更多信息
	r.GET("/Unaccept_order_show", Unaccept_order_show_handler)       //展示个人未被接收的订单
	r.GET("/Campus_order_Read_more", Campus_order_Read_more_handler) //校园跑腿订单查看更多详情
	r.GET("/CWork_order_Read_more", CWork_order_Read_more_handler)   //校园兼职订单查看更多详情
	r.GET("/Shared_order_Read_more", Shared_order_Read_more_handler) //共享订单查看更多详情
	r.GET("/Self_info", Self_info_handler)                           //个人信息
	r.GET("/self_new_order", self_new_order_handler)                 //个人订单
	r.GET("/setting", setting_handler)                               //设置
	r.GET("/self_accept_order", self_accept_order_handler)           //个人接收订单
	r.GET("/self_resume_order", self_resume_order_handler)           //个人发布订单
	r.GET("/queryselfaccept_order", quertselfaccept_handler)         //查看个人接收订单
	r.GET("/collect_order", collect_order_handler)                   //查看个人收藏订单
	r.GET("/collect_order_more", Collect_order_more_Handler)         //读取更多信息
	r.GET("/GetBalance_show", GetBalance_show_Handler)               //展示个人余额
	r.GET("/result_hadAcceptOrder", result_hadAcceptOrder_Handler)   //展示个人已接单订单结果
	r.GET("/Apply", Apply_Handler)                                   //跳转申请测试代币
	r.GET("/Frequency", Frequency)                                   //申请次数
	r.GET("/Apply_5", Apply)                                         //申请代币
	r.GET("/succ_login", Succ_Login)                                 //成功登录
	r.GET("/all_order", Self)
	r.POST("/QueryByTime", QueryByTime)
	r.GET("/delete_succ", delete_succ)
	r.GET("/create_succ", create_succ)
	r.GET("/exit_account", Exit_account)  //退出账户
	r.GET("/GetBlance", GetBalanceOfUser) //退出账户

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
		dapp.POST("/deleteData", DeleteData) //删除任务接口
		dapp.POST("/creatUser", CreateUser)
		dapp.POST("/login", Login)
		dapp.POST("/user_info", User_info)
		dapp.POST("/change_user_info", Change_user_info)
		dapp.POST("/Self_Order_show", Self_Order_show)
		dapp.POST("/Read_more", Read_more)
		dapp.POST("/CancelUser", CancelUser)                   //注销用户
		dapp.POST("/Unaccept_Order_show", Unaccept_Order_show) //展示未接单的订单
		dapp.POST("/Campus_order_detail", Campus_order)        //展示校园兼职详情
		dapp.POST("/CWork_order_detail", CWork_order)          //展示校园兼职详情
		dapp.POST("/Shared_order_detail", Shared_order)        //展示校园兼职详情
		dapp.POST("/Campus_order_Read_more", Campus_order_Read_more)
		dapp.POST("/CWork_order_Read_more", CWork_order_Read_more)
		dapp.POST("/Shared_order_Read_more", Shared_order_Read_more)
		dapp.POST("/collectOrder", CollectOrder)
		dapp.POST("/queryCollectOrder", QueryCollectOrder)
		dapp.POST("/Cancle_CollectOrder", CancleCollectOrder)
		dapp.POST("/Query_dim_Order", Query_Dim_order)
		dapp.POST("/Query_hadAcceptOrder", Query_hadAcceptOrder)
		dapp.POST("/ClaimTrust", ClaimTrust)
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
