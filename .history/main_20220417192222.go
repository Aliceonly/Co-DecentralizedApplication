package main

import (

	c"dapp/Connect"
	"fmt"
)


func main() {

	ins:=c.Getsmartcontract()
	head,_:=c.GetBlockNumber()
	_,adress:=c.Getaccout()
	fmt.Println("contractBalance",c.GetcontractBanlance(ins,adress,head))
	r.Run()
}
