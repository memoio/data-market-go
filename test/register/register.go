package main

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	cfdid "did-solidity/go-contracts/controlfiledid"
	"did-solidity/go-contracts/proxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	com "github.com/memoio/contractsv2/common"
	inst "github.com/memoio/contractsv2/go_contracts/instance"
)

var (
	eth   string
	hexSk string
	sks   [5]string
	as    [5]common.Address

	hash  []byte
	signs [5][]byte

	txfee = big.NewInt(1e12)

	bigone       = big.NewInt(1)
	defaultNonce = big.NewInt(0)
	set          = true

	allGas = make([]uint64, 0)

	methodType = "EcdsaSecp256k1VerificationKey2019"

	// user sk and did, as the controller of this mfiledid
	// test chain
	// user_sk = "11f797550cd4d77d08fd160047f9d55c8f468260c87e53a1f74505de4d9454be"
	// did     = "f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f18"
	// dev chain
	user_sk = "9db5e51e62c438bc32e0137bab95d73892d057faeea15d9868eb71c983945a80"
	did     = "f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f20"

	// file did
	fdid        = "bafkreibkkylda7ub52hkpl4ulbysvazjm2mcs2zjbjvfvy4hdaa2qnk4ne"
	encode      = "mid"
	ftype       = uint8(0)        // 0:private; 1:public
	price       = big.NewInt(100) // attomemo
	keywords    = []string{"picture", "animal", "dog", "white", "big"}
	newPrice    = big.NewInt(200)
	newFtype    = uint8(1)
	newKeywords = []string{"picture", "animal", "dog", "white", "big", "smile"}

	instanceAddr common.Address
)

//go run register.go -eth=test -sk=0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4
//go run register.go -eth=dev -sk=0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4

func main() {
	chain := flag.String("eth", "test", "eth api Address;") //dev test or product
	// TODO: read sk from local config
	sk := flag.String("sk", "", "signature for sending transaction")
	//sk1 := flag.String("sk1", "", "signature for on-chain")

	flag.Parse()

	// get instance address and endpoint with chain
	instanceAddr, eth = com.GetInsEndPointByChain(*chain)
	hexSk = *sk
	fmt.Println("instance address: ", instanceAddr)

	// tx sk is a must
	if len(hexSk) == 0 {
		log.Fatal("please input sk")
	}

	fmt.Println()

	// connect endpoint
	client, err := ethclient.DialContext(context.Background(), eth)
	if err != nil {
		log.Fatal(err)
	}

	// get chain id
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// make auth
	txAuth, err := com.MakeAuth(chainId, hexSk)
	if err != nil {
		log.Fatal(err)
	}

	// new instanceIns
	instanceIns, err := inst.NewInstance(instanceAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	// check controlFileDid address
	cfDidAddr, _ := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeFileDidControl)
	fmt.Println("controlFileDid address:", cfDidAddr.Hex())
	cfdidIns, err := cfdid.NewControlFileDid(cfDidAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	_ = cfdidIns
	var n uint64
	// fmt.Println("call cfdid.GetNonce")
	// // get nonce with did
	// n, err := cfdidIns.GetNonce(&bind.CallOpts{}, did)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("nonce of did: ", n)

	// check FileDid address
	fDidAddr, _ := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeFileDid)
	fmt.Println("FileDid address:", fDidAddr.Hex())

	// get proxy address
	proxyAddr, _ := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeDidProxy)
	fmt.Println("proxy address:", proxyAddr.Hex())

	// proxyIns, err := proxy.NewProxy(common.HexToAddress("0x98b10003642E517C8eCc1cee052E31Bb1C4d1d29"), client)
	proxyIns, err := proxy.NewProxy(proxyAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("call proxy.GetFileDidNonce")
	// test call proxy.getNonce
	n, err = proxyIns.GetFileDidNonce(&bind.CallOpts{}, did)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce: ", n)

	// str to ecdsa
	ecdsaSk, err := crypto.HexToECDSA(user_sk)
	if err != nil {
		// a random sk for wrong input
		ecdsaSk, err = crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
	}
	privateKeyBytes := ecdsaSk.D.Bytes() // D 是私钥的 big.Int 值
	fmt.Println("user sk:", hex.EncodeToString(privateKeyBytes))
	fmt.Println("user addr:", crypto.PubkeyToAddress(ecdsaSk.PublicKey))
	fmt.Println("did: ", did)

	//
	fmt.Println("fdid: ", fdid)

	// nonce
	var nonceBuf = make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBuf, 3)

	// sign in controlFileDid
	// bytes memory data = abi.encodePacked(
	// 	"registerMfileDid",
	// 	mfileDid,
	// 	encode,
	// 	ftype,
	// 	controller,
	// 	price,
	// 	nonce[controller]
	// );

	// type
	var typeBuf = make([]byte, 8)
	binary.BigEndian.PutUint64(typeBuf, 0)

	// price
	price := new(big.Int).SetInt64(3)
	// 目标字节数组（32字节）
	priceBuf := make([]byte, 32)
	// 使用 FillBytes 填充（大端序）
	price.FillBytes(priceBuf)
	fmt.Printf("Price Bytes: %x\n", priceBuf) // 输出 32 字节的 16 进制表示
	fmt.Println("buf len: ", len(priceBuf))

	// make msg for sign
	message := string("registerMfileDid") + fdid + encode + string(typeBuf[0]) + did + string(priceBuf) + string(nonceBuf)
	// append keywords into message
	for _, v := range keywords {
		message += v
	}
	fmt.Println("message: ", message)
	fmt.Printf("unprefixed message bytes: %x\n", []byte(message))

	// add ethereum prefix to message
	message = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	fmt.Println("prefixed message: ", message)
	fmt.Printf("prefixed message bytes: %x\n", []byte(message))

	// ethereum hash with message
	hash := crypto.Keccak256([]byte(message))
	fmt.Println("hash:", hex.EncodeToString(hash))

	// sign
	signature, err := crypto.Sign(hash, ecdsaSk)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("signature:", hex.EncodeToString(signature))

	fmt.Println("call proxy.RegisterMfileDid")

	// admin register a file did
	tx, err := proxyIns.RegisterMfileDid(txAuth, fdid, encode, ftype, did, price, keywords, signature)
	if err != nil {
		log.Fatal(err)
	}
	_ = tx

	fmt.Println("get gas used")
	go com.PrintGasUsed(eth, tx.Hash(), "admin register file did gas:", &allGas)
	fmt.Println("check tx")
	fmt.Println("endpoint: ", eth)
	err = com.CheckTx(eth, tx.Hash(), "admin register file did")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("waiting tx")
	time.Sleep(time.Second * 10)

	// get encode,ftype,controller,price,keywords,deactivated,read
	// encode
	fmt.Println("get encode")
	mfileEncode, err := proxyIns.GetEncode(&bind.CallOpts{From: com.AdminAddr}, fdid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file encode:", mfileEncode)
	if mfileEncode != encode {
		log.Fatal("file encode should be", encode, ", but is", mfileEncode)
	}
	// ftype
	mfileType, err := proxyIns.GetFtype(&bind.CallOpts{From: com.AdminAddr}, fdid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file type:", mfileType)
	if mfileType != ftype {
		log.Fatal("file type should be", ftype, ", but is", mfileType)
	}
	// controller
	mfileController, err := proxyIns.GetController(&bind.CallOpts{From: com.AdminAddr}, fdid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file's controller:", mfileController)
	if mfileController != did {
		log.Fatal("file controller should be", did, ", but is", mfileController)
	}
	// price
	mfilePrice, err := proxyIns.GetPrice(&bind.CallOpts{From: com.AdminAddr}, fdid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file read permission price:", mfilePrice)
	if mfilePrice.Cmp(price) != 0 {
		log.Fatal("price should be", price, ", but is ", mfilePrice)
	}
	// keywords
	mfileKeywords, err := proxyIns.GetKeywords(&bind.CallOpts{From: com.AdminAddr}, fdid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mfile keywords:", mfileKeywords)
	if len(mfileKeywords) != len(keywords) {
		log.Fatal("keywords should be", keywords, ", but is ", mfileKeywords)
	}
	// deactivated
	mfileDeactivated, err := proxyIns.Deactivated(&bind.CallOpts{From: com.AdminAddr}, fdid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file did deactivated:", mfileDeactivated)
	if mfileDeactivated {
		log.Fatal("mfile should be activated")
	}
	// read
	mfileRead, err := proxyIns.Read(&bind.CallOpts{From: com.AdminAddr}, fdid, did)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("has read permission:", mfileRead)
	if mfileRead != 0 {
		log.Fatal("should not has read permission")
	}
	fmt.Println()
}

// make hash
func makeHash() []byte {
	// 假设输入数据（替换为实际值）
	ftype := uint8(1) // enum 在 Solidity 中默认是 uint8
	controller := "exampleController"
	price := big.NewInt(1000000000000000000) // 1 ETH in wei (uint256)
	nonce := uint64(0)                       // uint64

	// 1. 将各字段按 abi.encodePacked 规则转换为字节
	var msgBytes []byte

	// 字符串类型：直接拼接字节
	msgBytes = append(msgBytes, []byte(fdid)...)
	msgBytes = append(msgBytes, []byte(encode)...)

	// enum (视为 uint8): 1字节
	msgBytes = append(msgBytes, byte(ftype))

	// 字符串类型：直接拼接字节
	msgBytes = append(msgBytes, []byte(controller)...)

	// uint256: 32字节小端序，高位补零
	priceBytes := make([]byte, 32)
	price.FillBytes(priceBytes) // 大端序转字节
	msgBytes = append(msgBytes, priceBytes...)

	// uint64: 8字节小端序
	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, nonce)
	msgBytes = append(msgBytes, nonceBytes...)

	// message := string("registerMfileDid") + fdid + encode + string(typeBuf) + did + string(bPrice) + string(nonceBuf)
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msgBytes), msgBytes)

	// ethereum hash with message
	hash := crypto.Keccak256([]byte(msg))

	return hash
}
