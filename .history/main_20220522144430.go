package main

import (
	c "dapp/Connect"

	// "fmt"
	s "dapp/Han"

	"github.com/ethereum/go-ethereum/common"
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
	s.Start()
	a:="0xfb098bd5f3a03750599a69c46d8d1a191f0d4c5c68cf948f9115291ff1299184"
	c.QueryStatus(common.HexToHash(a))
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

