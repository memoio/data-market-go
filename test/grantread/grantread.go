package main

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/big"

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
	eth     string
	adminSk string
	sks     [5]string
	as      [5]common.Address

	hash  []byte
	signs [5][]byte

	txfee = big.NewInt(1e12)

	bigone       = big.NewInt(1)
	defaultNonce = big.NewInt(0)
	set          = true

	allGas = make([]uint64, 0)

	methodType = "EcdsaSecp256k1VerificationKey2019"

	// user sk and did, as the controller of this mfiledid
	// dev chain
	controller_sk  = "9db5e51e62c438bc32e0137bab95d73892d057faeea15d9868eb71c983945a80"
	controller_did = "f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f20"
	// controller addr: 0x1E571f8a8Ad450A9453975B4207D40B25B16741b

	user_did = "f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f22"

	// file did
	fdid        = "did:mfile:mid:bafkreih6n5g5w4y6u7uvc4mh7jhjm7gidmkrbbpi7phyiyg54gplvngcpm"
	encode      = "mid"
	ftype       = uint8(0)        // 0:private; 1:public
	price       = big.NewInt(100) // attomemo
	keywords    = []string{"picture", "animal", "dog", "white", "big"}
	newPrice    = big.NewInt(200)
	newFtype    = uint8(1)
	newKeywords = []string{"picture", "animal", "dog", "white", "big", "smile"}

	instanceAddr common.Address
)

//go run grantread.go -eth=dev -sk=0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4

func main() {
	chain := flag.String("eth", "test", "eth api Address;") //dev test or product
	// TODO: read sk from local config
	sk := flag.String("sk", "", "signature for sending transaction")
	//sk1 := flag.String("sk1", "", "signature for on-chain")

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
	// make auth
	adminAuth, err := com.MakeAuth(chainId, adminSk)
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
	fmt.Println("controller sk:", hex.EncodeToString(privateKeyBytes))
	fmt.Println("controller addr:", crypto.PubkeyToAddress(ecdsaSk.PublicKey))
	fmt.Println("controller did: ", controller_did)

	//
	fmt.Println("fdid: ", fdid)

	// nonce
	var nonceBuf = make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBuf, 6)

	// sign in controlFileDid
	// bytes memory data = abi.encodePacked(
	// 	"grantRead",
	// 	mfileDid,
	// 	memoDid,
	// 	nonce[controller]
	// );

	// type
	var typeBuf = make([]byte, 8)
	binary.BigEndian.PutUint64(typeBuf, 0)

	// make msg for sign
	message := string("grantRead") + fdid + user_did + string(nonceBuf)
	fmt.Println("message: ", message)
	fmt.Printf("unprefixed message bytes: %x\n", []byte(message))

	// add ethereum prefix to message
	message = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	fmt.Println("prefixed message: ", message)
	fmt.Printf("prefixed message bytes: %x\n", []byte(message))

	// hash
	hash := crypto.Keccak256([]byte(message))
	fmt.Println("hash:", hex.EncodeToString(hash))

	// sign
	signature, err := crypto.Sign(hash, ecdsaSk)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("signature:", hex.EncodeToString(signature))

	fmt.Println("call proxy.grantRead")

	// admin call grantRead
	tx, err := proxyIns.GrantRead(adminAuth, fdid, user_did, signature)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("get gas used")
	go com.PrintGasUsed(eth, tx.Hash(), "admin register file did gas:", &allGas)
	fmt.Println("check tx")
	fmt.Println("endpoint: ", eth)
	err = com.CheckTx(eth, tx.Hash(), "admin register file did")
	if err != nil {
		log.Fatal(err)
	}

	// check read status about this filedid and did
	r, err := proxyIns.Read(&bind.CallOpts{}, fdid, user_did)
	fmt.Println("read status: ", r)
	if r != 2 {
		log.Fatal("read status must be 2 for grant read")
	}

	fmt.Println("buy read ok")

}
