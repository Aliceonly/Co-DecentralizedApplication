package main

import (
	"net/http"
	c"dapp/Connect"
	"github.com/gin-gonic/gin"
	"fmt"
)


func main() {

	ins:=c.Getsmartcontract()
	head,_:=c.GetBlockNumber()
	_,adress:=c.Getaccout()
	fmt.Println("contractBalance",c.GetcontractBanlance(ins,adress,head))
	r.Run()
}
