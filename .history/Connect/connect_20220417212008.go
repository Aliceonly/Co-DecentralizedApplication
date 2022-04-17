package Connect

import (
	"context"
	"crypto/ecdsa"
	contract "dapp/Smartgo"
	"fmt"
	"io/ioutil"
	"math/big"
     "fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var(
	//本地geth地址
	adress ="http://localhost:8545"
	//本地账户地址
	privatekeyfile="D://y//geth//node1//nodedata//keystore//UTC--2021-09-12T17-06-06.881126000Z--00dc6e8b60fa02a5d83e525bbef3240e8ea54dc5"
	//本地账户密码
	password="1111"
	//合约地址
	contractadress="0x2623F6cf09E3e6b6D43605ade1A30fDFA24E30f4"
)

var client *ethclient.Client
//连接geth
func init() {
	rpcDel, err := rpc.Dial(adress)
	if err != nil {
		panic(err)
	}
	client = ethclient.NewClient(rpcDel)
	//fmt.Println(client)
}
/*
实例化合约
 */
func Getsmartcontract() *contract.TaskDeployerContract{
	ins, err := contract.NewTaskDeployerContract(common.HexToAddress(contractadress), client)
	if err != nil {
		panic(err)
	}
	return ins
}

func Getaccout() (*ecdsa.PrivateKey, common.Address) {
	 file := privatekeyfile
	account, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	pwd :=password
	key, err := keystore.DecryptKey(account, pwd)
	if err != nil {
		panic(err)
	}
	//fmt.Println(key.PrivateKey, key.Address)
	return key.PrivateKey, key.Address
}
//获取gasprice
func GetgasPrice() (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return big.NewInt(0), err
	} else {
		return gasPrice, nil
	}

}
//获取nonce
func Getnonce(address common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, err
	} else {
		return nonce, nil
	}
}
//获取区块数
func GetBlockNumber() (*types.Header,error) {
	header,err :=client.HeaderByNumber(context.Background(),nil)
	if err!=nil {
		panic(err)
	}
	//fmt.Println(header)
	return header, err
}
//设置TransactOpts
func setopts(privateKey *ecdsa.PrivateKey, address common.Address) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(10001))
	if err != nil {
		panic(err)
	}
	nonce, err := Getnonce(address)
	if err != nil {
		panic(err)
	}
	gasPrice, err := GetgasPrice()
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth
}
func GetTxopts() *bind.TransactOpts{
	privateKey, publicKey := Getaccout()
	opts := setopts(privateKey, publicKey)
	return opts
}
/*
获取合约余额
 */
func GetcontractBanlance(ins *contract.TaskDeployerContract, address common.Address, header *types.Header) (*big.Int)  {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	balance, err := ins.GetBalanceOfContract(&opts)
	if err!=nil {
		panic(err)
	}
	return balance
}
/*
获取任务列表函数
*/
func GetTasklist(ins *contract.TaskDeployerContract, timestap *big.Int, address common.Address, header *types.Header) struct {
	Taskname    string
	Sponsor     common.Address
	Beneficiary common.Address
	Category    string
	Amount      *big.Int
	Timestamp   *big.Int
	State       string
	LaunchTime  string
} {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	Tasklist,err:=ins.Tasklist(&opts,timestap)
	if err!=nil {
		panic(err)
	}
	return Tasklist
}
/*
获取用户信息列表
*/
func Getuser(ins *contract.TaskDeployerContract, timestap *big.Int, address common.Address, header *types.Header) struct {
	Name        string
	Phonenumber string
	Studentid   *big.Int
	Password    string
} {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	UserList,err:=ins.Userlist(&opts,timestap)
	if err!=nil {
		panic(err)
	}
	return UserList

}
/*
创建任务函数
*/
func CreatNewEvent(
	ins *contract.TaskDeployerContract,
	ops *bind.TransactOpts,
	Taskname string,
	Taskcatagory string,
	launchTime string,
	amount *big.Int) *types.Transaction{
	timestap,err:=ins.CreateNewEvent(ops,Taskname,Taskcatagory,launchTime,amount)
	if err!=nil {
		panic(err)
	}
	fmt.
	(timestap.AccessList())
	print(timestap.Cost())
	return timestap
}
/*
取消任务函数
*/

func CancelEvent(
	ins *contract.TaskDeployerContract,
	ops *bind.TransactOpts,
	timestamp *big.Int){
	_,err:=ins.CancelEvent(ops,timestamp)
	if err!=nil {
		panic(err)
	}
}
/*
接受任务函数
*/
func Confirmtask(
	ins *contract.TaskDeployerContract,
	timestamp *big.Int,
    ops *bind.TransactOpts){
	_,err:=ins.Confirmtask(ops,timestamp)
	if err!=nil {
		panic(err)
	}
}
/*
发布任务者确认接受任务者完成任务
*/
func ClaimTrust(
	ins *contract.TaskDeployerContract,
	ops *bind.TransactOpts,
	timestamp *big.Int,
	sigs []byte,
	PHash [32]byte,
	taskname string){
	_,err:=ins.ClaimTrust(ops,timestamp,sigs,PHash,taskname)
	if err!=nil {
		panic(err)
	}
}
/*
创建用户信息
*/
func Creatuser(
	ins *contract.TaskDeployerContract,
    opts *bind.TransactOpts,
	name string,
	phonenumber string,
	studentid *big.Int,
	password string) bool{
   _,err:=ins.Createuser(opts,name,phonenumber,studentid,password)
	if err!=nil {
		panic(err)
	}
	return true
}
/*
用户修改信息
*/
   //修改密码
func Userchangepsword(
	ins *contract.TaskDeployerContract,
	opts *bind.TransactOpts,
	password string,
	studentid *big.Int) bool {
	_,err:=ins.Changepassword(opts,password,studentid)
	if err!=nil {
		panic(err)
	}
	return true
}
   //修改名称
func Userchangename(ins *contract.TaskDeployerContract,
	opts *bind.TransactOpts,
	name string,
	studentid *big.Int,
	password string) bool {
	_,err:=ins.Changename(opts,name,studentid,password)
	if err!=nil {
		panic(err)
	}
	return true
}

