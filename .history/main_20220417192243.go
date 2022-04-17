package main

import (

	c"dapp/Connect"
	"fmt"
	"dapp/Han"
)


func main() {

	ins:=c.Getsmartcontract()
	head,_:=c.GetBlockNumber()
	_,adress:=c.Getaccout()
	fmt.Println("contractBalance",c.GetcontractBanlance(ins,adress,head))

}
