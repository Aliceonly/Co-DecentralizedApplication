package Connect

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	contract "dapp/Smartgo"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	// "unsafe"

	// "strconv"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/ptypes/timestamp"

	// "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	// "log"
	"os"

	"github.com/lithammer/fuzzysearch/fuzzy"
	// "fmt"
)

var (
	//本地geth地址
	adress = "http://localhost:8545"
	//本地账户地址
	privatekeyfile = ""
	//本地账户密码
	password = ""
	//合约地址
	contractadress = "0x1c154b7A518876DF365b3d209650d2A2fBfFeA5B"
	//读取用户keystore文件地址
	relativePath = "D:\\y\\geth\\node1\\nodedata\\keystore"
	//本地链chainID交易:修改为本地的chainID
	chainID = 10001
)

var client *ethclient.Client
var rDel *rpc.Client

//连接geth
func init() {
	rpcDel, err := rpc.Dial(adress)
	if err != nil {
		fmt.Println("连接geth====>", err)
		panic(err)
	} else {
		fmt.Println("geth连接成功*****============================***********")
	}
	rDel = rpcDel
	client = ethclient.NewClient(rpcDel)
	//fmt.Println(client)
}

/*
实例化合约
*/
func Getsmartcontract() *contract.TaskDeployerContract {
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
	pwd := password
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
func GetBlockNumber() (*types.Header, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	//fmt.Println(header)
	return header, err
}

//获取区块的详细信息
func Getblockmessage() (*big.Int,uint64) {
	blockNumber := big.NewInt(int64(chainID))
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	blockNow:=block.Number()
	timestamp:=block.Time()
	fmt.Println(block.Number().Uint64()) 
	fmt.Println(block.Time())
	// fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	// fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc38504266149
	// fmt.Println(len(block.Transactions()))   // 144个交易记录
	return blockNow,timestamp
}

// func getBlockTransactionCountByNumber(client *rpc.Client,j string)() {
// 	err:=client.Call(&BlockTransactionCountByNumber,"eth_getBlockTransactionCountByNumber",j)
// 	if err!=nil{
// 		fmt.Println("错误:",err)
// 	}
// 	heights, _ := strconv.ParseUint(j[2:],16,32)
// 	BlockTransactionCountByNumbers, _ := strconv.ParseUint(BlockTransactionCountByNumber[2:],16,32)
// 	fmt.Println("当前块:",heights,"当前块交易数量:",BlockTransactionCountByNumbers)
// 	GetBlockByNumber(client,j)
// }

//设置TransactOpts
func setopts(privateKey *ecdsa.PrivateKey, address common.Address) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainID)))
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
func GetTxopts() *bind.TransactOpts {
	privateKey, publicKey := Getaccout()
	opts := setopts(privateKey, publicKey)
	return opts
}

/*
获取合约余额
*/
func GetcontractBanlance(ins *contract.TaskDeployerContract, address common.Address, header *types.Header) *big.Int {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	balance, err := ins.GetBalanceOfContract(&opts)
	if err != nil {
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
	Tasklist, err := ins.Tasklist(&opts, timestap)
	if err != nil {
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
	UserList, err := ins.Userlist(&opts, timestap)
	if err != nil {
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
	amount *big.Int,
) (*big.Int,uint64) {
	ops.Value = amount
	_, err := ins.CreateNewEvent(ops, launchTime, Taskcatagory, Taskname, amount)
	if err != nil {
		fmt.Println("CreatNewEvent error ===>", err)
		panic(err)
	}
	// blocknumber := ops.Nonce
	blocknumber,timestap:=Getblockmessage()
	// fmt.Println(timestap.AccessList())
	// fmt.Println(timestap.Cost())
	// fmt.Println(timestap.Type())
	return blocknumber, timestap
}

/*
获取任务时间戳
*/
func Querytime(ins *contract.TaskDeployerContract,
	address common.Address,
	header *types.Header) *big.Int {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	timestamp, err := ins.Query(&opts)
	if err != nil {
		panic(err)
	}
	return timestamp
}

/*
取消任务函数
*/

func CancelEvent(
	ins *contract.TaskDeployerContract,
	ops *bind.TransactOpts,
	timestamp *big.Int,
	address common.Address) {
	ops.From = address
	_, err := ins.CancelEvent(ops, timestamp)
	if err != nil {
		panic(err)
	}
}

/*
接受任务函数
*/
func Confirmtask(
	ins *contract.TaskDeployerContract,
	timestamp *big.Int,
	ops *bind.TransactOpts,
	address common.Address) {
	ops.From = address
	_, err := ins.Confirmtask(ops, timestamp)
	if err != nil {
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
	taskname string) {
	_, err := ins.ClaimTrust(ops, timestamp, sigs, PHash, taskname)
	if err != nil {
		panic(err)
	}
}

//获取交易的hash值
func Gettaskhash(ins *contract.TaskDeployerContract, address common.Address, header *types.Header, taskname string, timestamp string) [32]byte {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	hash, err := ins.GetTxHash(&opts, taskname, timestamp)
	if err != nil {
		panic(hash)
	}
	return hash
}

//获取发布任务的用户对当前任务的确认签名
func GetthistaskSign(PrivateKey *ecdsa.PrivateKey, hash [32]byte) []byte {
	signature, err := ecdsa.SignASN1(rand.Reader, PrivateKey, hash[:])
	if err != nil{
		panic(err)	
	}
	// fmt.Printf("signature: %x\n", signature)
	fmt.Println(signature)


	// signature, err := crypto.Sign(hash[:], PrivateKey)

	return signature
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
	password string) bool {
	_, err := ins.Createuser(opts, name, phonenumber, studentid, password)
	if err != nil {
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
	_, err := ins.Changepassword(opts, password, studentid)
	if err != nil {
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
	_, err := ins.Changename(opts, name, studentid, password)
	if err != nil {
		panic(err)
	}
	return true
}

var newAccount string
var accounts []string

func CreatnewActogeth(pd string) string {
	// fmt.Print("why----->", rDel)
	err := rDel.Call(&newAccount, "personal_newAccount", pd)
	if err != nil {
		panic(err)
	}
	rDel.Call(&accounts, "personal_listAccounts")
	// if err!=nil {
	// 	panic(err)
	// }
	fmt.Print(accounts)
	return newAccount
}

//针对不同用户登入获取不同用户的信息来对交易签名
func Changeuser(ad string, pw string) {
	// relativePath := "d://Test_Block_chain//data//keystore"
	var FileInfo []os.FileInfo
	var err error

	if FileInfo, err = ioutil.ReadDir(relativePath); err != nil {
		fmt.Println("读取 keystore 文件夹出错")
		return
	}
	a := make([]string, 0)
	for _, fileInfo := range FileInfo {
		a = append(a, fileInfo.Name())
		// fmt.Println(fileInfo.Name())
	}
	// ac:="5c595872e02b0613658036bdf5daa6d9f42954be"
	fmt.Print("keystore", a)
	matches2 := fuzzy.Find(ad[2:], a)
	fmt.Println("当前登入的用户为", ad)
	// print(relativePath+"//"+matches1[0])
	fmt.Println(".....", matches2)
	privatekeyfile = relativePath + "\\" + matches2[0]
	password = pw
	fmt.Println("privatekeyfile and pw:", privatekeyfile, pw)
	// matches2 = nil
}

//用户退出状态 文件处于空状态
func Userexit() {
	privatekeyfile = ""
	password = ""
}
func Get() (string, string) {
	return password, privatekeyfile
}

//注销用户
func Cancellation(ad string) string {
	var FileInfo []os.FileInfo
	var err error
	relativePath := "D://Test_Block_chain//data//keystore"

	if FileInfo, err = ioutil.ReadDir(relativePath); err != nil {
		fmt.Println("读取 keystore 文件夹出错")
		return err.Error()
	}
	a := make([]string, 0)
	for _, fileInfo := range FileInfo {
		a = append(a, fileInfo.Name())
	}
	matches2 := fuzzy.Find(ad[2:], a)
	adfile := relativePath + "//" + matches2[0]
	err2 := os.Remove(adfile)
	if err2 != nil {
		panic(err2)
	} else {
		fmt.Println("用户注销完毕")
		matches2 = nil
		result := "成功注销"
		return result
	}

}
