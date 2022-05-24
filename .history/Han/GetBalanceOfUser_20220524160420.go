func GetBalanceOfUser(c *gin.context){
	ins := c.Getsmartcontract()
	head, _ := c.GetBlockNumber()
	_, adress := c.Getaccout()
	fmt.Println(c.GetuserBanlance(ins,adress,head))
}