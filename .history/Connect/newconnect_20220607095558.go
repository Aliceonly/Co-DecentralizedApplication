package Connect

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
    "time"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type accountBody struct {
	File    string
	Passwd  string
	Address string
}

func newAcc(file string) *accountBody {
	u := accountBody{File: file}
	u.Passwd = "111111"
	return &u
}

type clientManage struct {
	rpcConn *rpc.Client
	ethConn *ethclient.Client
}

func newClient(rawurl string) *clientManage {
	rpcDial, err := rpc.Dial(rawurl)
	if err != nil {
		panic(err)
	}
	client := ethclient.NewClient(rpcDial)

	return &clientManage{
		rpcConn: rpcDial,
		ethConn: client,
	}
}

func A( adress string) {
	userA := newAcc("D:\\y\\geth\\node1\\nodedata\\keystore\\UTC--2021-09-12T17-06-06.881126000Z--00dc6e8b60fa02a5d83e525bbef3240e8ea54dc5")
	toAddress := adress

	c := newClient("http://127.0.0.1:8545")
	defer c.ethConn.Close()

	privateKey, fromAddress, err := keyStoreToPrivateKey(&userA.File, &userA.Passwd)

	if err != nil {
		log.Fatal(err)
	}

	printItemBalance := func(address string) {
		ethValue, _ := c.getBalance(address)
		fmt.Printf("Address:%s Eth:%s \n", address, ethValue)

	}

	printItemBalance(fromAddress)
	printItemBalance(toAddress)

	weiValue := EthToWei(1.5)

	//本人私钥签名，广播给节点
	txSend, _ := c.transferEth(privateKey, fromAddress, toAddress, weiValue)
	fmt.Printf("txSend: %s \n", txSend)
	
	//等待挖矿确认
	time.Sleep(time.Duration(5) * time.Second)

	printItemBalance(fromAddress)
	printItemBalance(toAddress)

}

// transferEth ETH转账
func (ec *clientManage) transferEth(pk, fromAddress, toAddress string, weiValue *big.Int) (string, error) {
	//加载私钥
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal("privateKey ", err)

	}

	//派生公钥用于签名
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicfromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//获得帐户的随机数
	nonce, err := ec.ethConn.PendingNonceAt(context.Background(), publicfromAddress)
	if err != nil {
		log.Fatal("nonce ", err)
	}

	gasLimit := uint64(21000) // in units

	//获取市场波动的gas价格
	gasPrice, err := ec.ethConn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("gasPrice ", err)
	}

	//发送给谁
	tAddress := common.HexToAddress(toAddress)

	//生成未签名的事务
	tx := types.NewTransaction(nonce, tAddress, weiValue, gasLimit, gasPrice, nil)

	//获取chainID
	// chainID, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	log.Fatal("chainID ", err)
	// }
	//私链需要指定创世区块的ID
	chainID := big.NewInt(10001)

	//EIP155签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("check your chainID!!! ", err)
	}

	//广播到整个网络
	if err := ec.ethConn.SendTransaction(context.Background(), signedTx); err != nil {
		log.Fatal("SendTransaction ", err)
	}

	return signedTx.Hash().Hex(), err
}

// getBalance 根据address 从状态树 获取ETH余额
func (ec *clientManage) getBalance(address string) (string, error) {
	account := common.HexToAddress(address)
	balance, err := ec.ethConn.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "", err
	}
	//单位wei转化eth
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue.String(), err
}

//EthToWei eth单位安全转wei
//https://stackoverrun.com/cn/q/13021596
func EthToWei(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	// Set precision if required.
	// bigval.SetPrec(64)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}

// keyStoreToPrivateKey 根据keystore文件和密码 获取私钥和地址
func keyStoreToPrivateKey(privateKeyFile, password *string) (string, string, error) {
	keyJSON, err := ioutil.ReadFile(*privateKeyFile)
	if err != nil {
		fmt.Println("read keyjson file failed：", err)
	}
	unlockedKey, err := keystore.DecryptKey(keyJSON, *password)
	if err != nil {
		return "", "", err
	}
	privKey := hex.EncodeToString(unlockedKey.PrivateKey.D.Bytes())
	addr := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)

	return privKey, addr.String(), nil

}