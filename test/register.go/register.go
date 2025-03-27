package main

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/big"

	"did-solidity/go-contracts/proxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	com "github.com/memoio/contractsv2/common"
	inst "github.com/memoio/contractsv2/go_contracts/instance"
)

var (
	eth      string
	hexSk    string
	sks      [5]string
	as       [5]common.Address
	endPoint string

	hash  []byte
	signs [5][]byte

	txfee = big.NewInt(1e12)

	bigone       = big.NewInt(1)
	defaultNonce = big.NewInt(0)
	set          = true

	allGas = make([]uint64, 0)

	methodType = "EcdsaSecp256k1VerificationKey2019"

	fdid        = "bafkreibkkylda7ub52hkpl4ulbysvazjm2mcs2zjbjvfvy4hdaa2qnk4na"
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

func main() {
	chain := flag.String("eth", "test", "eth api Address;") //dev test or product
	// TODO: read sk from local config
	sk := flag.String("sk", "", "signature for sending transaction")
	//sk1 := flag.String("sk1", "", "signature for on-chain")

	flag.Parse()

	// get instance address and endpoint with chain
	instanceAddr, eth = com.GetInsEndPointByChain(*chain)
	hexSk = *sk

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
	// get proxy address
	proxyAddr, _ := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeDidProxy)
	fmt.Println("proxy address:", proxyAddr.Hex())

	// proxyIns, err := proxy.NewProxy(common.HexToAddress("0x98b10003642E517C8eCc1cee052E31Bb1C4d1d29"), client)
	proxyIns, err := proxy.NewProxy(proxyAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	// sk and did, as the controller of this mfiledid
	user_sk := "11f797550cd4d77d08fd160047f9d55c8f468260c87e53a1f74505de4d9454be"
	fmt.Println("sk: ", user_sk)
	did := "f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f18"
	fmt.Println("did: ", did)

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

	//
	fmt.Println("fdid: ", fdid)

	// nonce
	var nonceBuf = make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBuf, 0)

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

	price := new(big.Int).SetInt64(123)

	// make msg for sign
	message := string("registerMfileDid") + fdid + encode + string(typeBuf) + did + string(price.Bytes()) + string(nonceBuf)

	// ethereum hash with message
	hash := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
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
	go com.PrintGasUsed(endPoint, tx.Hash(), "admin register file did gas:", &allGas)
	err = com.CheckTx(endPoint, tx.Hash(), "admin register file did")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	// get encode,ftype,controller,price,keywords,deactivated,read
	// encode
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
