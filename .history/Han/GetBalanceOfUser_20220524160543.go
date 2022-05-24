package Han

func GetBalanceOfUser(c *gin.Context){
	ins := c.Getsmartcontract()
	head, _ := c.GetBlockNumber()
	_, adress := c.Getaccout()
	fmt.Println(c.GetuserBanlance(ins,adress,head))
}