package main

import (
	c "dapp/Connect"
	"fmt"
	"math/big"

	// "fmt"

	// s "dapp/Han"
	// "github.com/ethereum/go-ethereum/common"
	// "dapp/controller"
	// "fmt"
	// "math/big"
	// "net/http"
	// "github.com/go-sql-driver/mysql"
	// mysql "dapp/utils"
	// "io/ioutil"
	// "os"
	// "github.com/lithammer/fuzzysearch/fuzzy"
)

func main() {
	// s.Start()
	ins := c.Getsmartcontract()
	head, _ := c.GetBlockNumber()
	pr, adress := c.Getaccout()
	taskname:="买饮料校园区域主要建筑"
	times:="1653212002"
	// Txopts := c.GetTxopts()
	hash := c.Gettaskhash(ins, adress, head, taskname, times)
	sigh := c.GetthistaskSign(pr, hash)
	// a:=big.NewInt(1653212002)
	// c.ClaimTrust(ins,Txopts,a,sigh,hash,taskname)
	fmt.Println(c.Validation(hash[:],sigh,pr.Public()))
	// a:="0xc69c51337ca843b19648226e3626d0851c8b9b394201f25c61b3985a13d39811"
	// c:=c.QueryStatus(common.HexToHash(a))
	// fmt.Println(c)
// c.GetBlockNumber()
// c.Getblockmessage()
	// c.Getblockmessage()
	// a:="0x5e9ed07526dd85415af6a2f367a66efb02637ba0"
	// fmt.Println(a[2:])
	// c.Changeuser("aebdbdb1455a7d56d56fe2cf967fbe7118845e66","abc")
	// a,b:=c.Get()
	// print(a,b)
	// a:=c.CreatnewActogeth("123")
	// print("newaccount",a)
	// mysql.DeletTask(1650442602)
	// big:=big.NewInt(2)
	// err:=mysql.Insert("aa","0xxx","a",2,big,"availablev","200202")
	// if err!=nil {
	// 	panic(err)
	// }
	// fmt.Println(mysql.Select("0xxx"))
	// fmt.Print(mysql.Selectaccept("0xxx000"))
}

