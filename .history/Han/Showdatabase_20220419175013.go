package Han

func Index_Handler(c *gin.Context) {
	// c.HTML(200, "index.html", nil) //首页
	taskdata := mysql.Showdata()
	fmt.Println("get____taskdata=======>>>>>>>>>>>>>>", taskdata)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"data": taskdata,
	}) // -------------------------第三步---------------------

}