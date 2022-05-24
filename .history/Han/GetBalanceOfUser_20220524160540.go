package Han

func GetBalanceOfUser(c *Context.context){
	ins := c.Getsmartcontract()
	head, _ := c.GetBlockNumber()
	_, adress := c.Getaccout()
	fmt.Println(c.GetuserBanlance(ins,adress,head))
}