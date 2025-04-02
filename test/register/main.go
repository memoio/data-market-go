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

	// sks   [5]string
	// as    [5]common.Address

	// hash  []byte
	// signs [5][]byte

	// txfee = big.NewInt(1e12)

	// bigone       = big.NewInt(1)
	// defaultNonce = big.NewInt(0)
	// set          = true

	allGas = make([]uint64, 0)

	// methodType = "EcdsaSecp256k1VerificationKey2019"

	// user sk and did, as the controller of this mfiledid
	// test chain
	// user_sk = "11f797550cd4d77d08fd160047f9d55c8f468260c87e53a1f74505de4d9454be"
	// did     = "f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f18"
	// dev chain
	controller_sk  = "7ad6e373d75363a20a7851a00aa6204c52e70b26f5499c0ba32a119058d4afdd"
	controller_did = "f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f81"

	// file did
	fdid   = "bafkreih6n5g5w4y6u7uvc4mh7jhjm7gidmkrbbpi7phyiyg54gplvngcpn"
	encode = "mid"
	ftype  = uint8(0) // 0:private; 1:public
	// price       = big.NewInt(100) // attomemo
	keywords = []string{"sport", "basketball"}
	// newPrice    = big.NewInt(200)
	// newFtype    = uint8(1)
	// newKeywords = []string{"picture", "animal", "dog", "white", "big", "smile"}

	instanceAddr common.Address
)

//go run main.go -eth=test -sk=0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4
//go run main.go -eth=dev -sk=0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4

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
	n, err = proxyIns.GetFileDidNonce(&bind.CallOpts{}, controller_did)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce: ", n)

	// str to ecdsa
	ecdsaSk, err := crypto.HexToECDSA(controller_sk)
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
	fmt.Println("did: ", controller_did)

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

	// price
	price := new(big.Int).SetInt64(3)
	// 目标字节数组（32字节）
	priceBuf := make([]byte, 32)
	// 使用 FillBytes 填充（大端序）
	price.FillBytes(priceBuf)
	fmt.Printf("Price Bytes: %x\n", priceBuf) // 输出 32 字节的 16 进制表示
	fmt.Println("buf len: ", len(priceBuf))

	// make msg for sign
	message := string("registerMfileDid") + fdid + encode + string(typeBuf[0]) + controller_did + string(priceBuf) + string(nonceBuf)
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
	tx, err := proxyIns.RegisterMfileDid(txAuth, fdid, encode, ftype, controller_did, price, keywords, signature)
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
	if mfileController != controller_did {
		log.Fatal("file controller should be", controller_did, ", but is", mfileController)
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
	mfileRead, err := proxyIns.Read(&bind.CallOpts{From: com.AdminAddr}, fdid, controller_did)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("has read permission:", mfileRead)
	if mfileRead != 0 {
		log.Fatal("should not has read permission")
	}

	fmt.Println("register filedid ok")
}

// go run register.go -eth=dev -sk=0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4
// instance address:  0x10790c26fB42AaDB87c10b91a25090AF1Ff5352E

// controlFileDid address: 0x9dCf172c6179044be7dd5F7967231697dA131476
// FileDid address: 0x7272d844d5d78E2F8Fb0780582739080C8DC0B5a
// proxy address: 0xE2C89b43724ed4932626e937742347280A5dcb4D
// call proxy.GetFileDidNonce
// nonce:  3
// user sk: 9db5e51e62c438bc32e0137bab95d73892d057faeea15d9868eb71c983945a80
// user addr: 0x1E571f8a8Ad450A9453975B4207D40B25B16741b
// did:  f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f20
// fdid:  bafkreibkkylda7ub52hkpl4ulbysvazjm2mcs2zjbjvfvy4hdaa2qnk4ne
// Price Bytes: 0000000000000000000000000000000000000000000000000000000000000003
// buf len:  32
// message:  registerMfileDidbafkreibkkylda7ub52hkpl4ulbysvazjm2mcs2zjbjvfvy4hdaa2qnk4nemidf3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f20pictureanimaldogwhitebig
// unprefixed message bytes: 72656769737465724d66696c654469646261666b726569626b6b796c64613775623532686b706c34756c62797376617a6a6d326d6373327a6a626a76667679346864616132716e6b346e656d696400663330353339343664376663623735653338306638653431353164656431343536616265363764643736303731303166646439636331396330643162336632300000000000000000000000000000000000000000000000000000000000000003000000000000000370696374757265616e696d616c646f677768697465626967
// prefixed message:  Ethereum Signed Message:
// 207registerMfileDidbafkreibkkylda7ub52hkpl4ulbysvazjm2mcs2zjbjvfvy4hdaa2qnk4nemidf3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f20pictureanimaldogwhitebig
// prefixed message bytes: 19457468657265756d205369676e6564204d6573736167653a0a32303772656769737465724d66696c654469646261666b726569626b6b796c64613775623532686b706c34756c62797376617a6a6d326d6373327a6a626a76667679346864616132716e6b346e656d696400663330353339343664376663623735653338306638653431353164656431343536616265363764643736303731303166646439636331396330643162336632300000000000000000000000000000000000000000000000000000000000000003000000000000000370696374757265616e696d616c646f677768697465626967
// hash: 867a02a41ee9cadc3ebe45a92812de92a7747e101bcbd151181d5be544c80b0e
// signature: cb403f12f22244fb649a6e7e588f5a2c24eca0c92202becb914817274fe03e3e2859c40e2fb26ee2ec4cd97e78ddc70e4eb426175ddbb8d7a4ff152d62b19b2b00
// call proxy.RegisterMfileDid
// get gas used
// check tx
// endpoint:
// waiting tx
// get encode
// file encode: mid
// file type: 0
// file's controller: f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f20
// file read permission price: 3
// mfile keywords: [picture animal dog white big]
// file did deactivated: false
// has read permission: 0
