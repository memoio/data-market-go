package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"did-solidity/go-contracts/proxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	com "github.com/memoio/contractsv2/common"
	"github.com/memoio/contractsv2/go_contracts/erc"
	inst "github.com/memoio/contractsv2/go_contracts/instance"
)

var (
	eth     string
	adminSk string

	l     = "1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed"
	lBase = "14def9dea2f79cd65812631a5cf5d3ed"

	// // AdminAddr admin address
	// AdminAddr = common.HexToAddress("0x1c111472F298E4119150850c198C657DA1F8a368")

	//address      = common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97")
	instanceAddr common.Address

	// // params
	// scalar = big.NewInt(12)
)

//go run buyread.go -eth=dev -sk=0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4

func main() {
	chain := flag.String("eth", "test", "eth api Address;") //dev test or product
	// TODO: read sk from local config
	sk := flag.String("sk", "", "signature for sending transaction")

	flag.Parse()

	// get instance address and endpoint with chain
	instanceAddr, eth = com.GetInsEndPointByChain(*chain)
	adminSk = *sk

	fmt.Println("instance address: ", instanceAddr)

	// tx sk is a must
	if len(adminSk) == 0 {
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
	// make auth for admin
	adminAuth, err := com.MakeAuth(chainId, adminSk)
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

	// get controleFileDid address
	controlFileDidAddr, _ := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeFileDidControl)
	fmt.Println("controlFileDidAddr :", controlFileDidAddr.Hex())

	// file did
	fdid := "did:mfile:mid:bafkreih6n5g5w4y6u7uvc4mh7jhjm7gidmkrbbpi7phyiyg54gplvngcpm"
	fmt.Println("file did: ", fdid)

	// user sk and did
	user_sk := "9db5e51e62c438bc32e0137bab95d73892d057faeea15d9868eb71c983945a80"
	user_addr := "0x1E571f8a8Ad450A9453975B4207D40B25B16741b"
	did := "f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f20"
	fmt.Println("user sk: ", user_sk)
	fmt.Println("user address: ", user_addr)
	fmt.Println("did: ", did)

	// user auth
	userAuth, err := com.MakeAuth(chainId, user_sk)
	if err != nil {
		log.Fatal(err)
	}

	// get erc20 address from instance
	erc20Addr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeERC20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("erc20Addr:", erc20Addr.Hex())
	// new erc20 instance
	erc20Ins, err := erc.NewERC20(erc20Addr, client)
	if err != nil {
		log.Fatal(err)
	}

	// check filedid price
	// price
	mfilePrice, err := proxyIns.GetPrice(&bind.CallOpts{From: com.AdminAddr}, fdid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file read permission price:", mfilePrice)

	fmt.Println("transfering price to user")
	// transfer money to user
	tx, err := erc20Ins.Transfer(adminAuth, common.HexToAddress(user_sk), mfilePrice)
	if err != nil {
		log.Fatal(err)
	}
	//go com.PrintGasUsed(eth, tx.Hash(), "transfer memo gas:", nil)
	err = com.CheckTx(eth, tx.Hash(), "admin transfer memo to user")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	// approve to ControlFileDid
	userAuth, err = com.MakeAuth(chainId, user_sk)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("approve to controlFileDid")
	// approve
	tx, err = erc20Ins.Approve(userAuth, controlFileDidAddr, mfilePrice)
	if err != nil {
		log.Fatal(err)
	}
	//go com.PrintGasUsed(eth, tx.Hash(), "user approve ControlFileDid gas:", &allGas)
	err = com.CheckTx(eth, tx.Hash(), "user approve ControlFileDid")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("call proxy.BuyRead")

	// call buyRead
	tx, err = proxyIns.BuyRead(userAuth, fdid, did)
	if err != nil {
		log.Fatal(err)
	}
	err = com.CheckTx(eth, tx.Hash(), "buyRead")
	if err != nil {
		log.Fatal(err)
	}

	// check read status about this filedid and did
	r, err := proxyIns.Read(&bind.CallOpts{}, fdid, did)
	if r != 1 {
		log.Fatal("read status must be 1")
	}

	fmt.Println("buy read ok")
}
