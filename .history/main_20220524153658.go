package main

import (
	c "dapp/Connect"


	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/crypto"
	// // "fmt"
	s "dapp/Han"
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
	// ins := c.Getsmartcontract()

	pr, adress := c.Getaccout()

	// hash:="0xe67604b8d0af2baf6eb092db37142e859e515e5505826c09843093b4a8c25f88"
	// sigh :="0xbde63ad0c40820b2a148c8f24f916795b2ad87ce9a85c95c13218fcf04f2e32e730fab525f0944660d0a5d8325de1f21374da6a301053e7ccbc71d415e47bd3d01"
	// Txopts := c.GetTxopts()
	// s,_:=hexutil.Decode(sigh)
	// h,_:=hexutil.Decode(hash)
	// fmt.Println(s)
	// a:=big.NewInt(1653359260)
	// // c.ClaimTrust(ins,Txopts,a,sigh,hash,taskname)
	// publice:=pr.Public()
	// publicKeyECDSA, ok := publice.(*ecdsa.PublicKey)
    // if !ok {
    //     panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    // }
	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// if c.Validation(h,s,publicKeyBytes) ==true{
	// 	c.ClaimTrust(ins,Txopts,a)
	// 	fmt.Println("ok")
	// }else{
	// 	fmt.Println("err")
	// }
}



