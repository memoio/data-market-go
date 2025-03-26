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
	eth   string
	hexSk string

	l     = "1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed"
	lBase = "14def9dea2f79cd65812631a5cf5d3ed"

	// // AdminAddr admin address
	// AdminAddr = common.HexToAddress("0x1c111472F298E4119150850c198C657DA1F8a368")

	//address      = common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97")
	instanceAddr common.Address

	// // params
	// scalar = big.NewInt(12)
)

//go run create_did.go -eth=test -sk=0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4 -sk1=9b4fc2a14cbc63a0d338377413ca72949cbb2fd5be1b08844b4b5e332597d91e
//go run create_did.go -eth=test -sk=0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4

func main() {
	chain := flag.String("eth", "test", "eth api Address;") //dev test or product
	// TODO: read sk from local config
	sk := flag.String("sk", "", "signature for sending transaction")
	sk1 := flag.String("sk1", "", "signature for on-chain")

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

	// str to ecdsa
	userSk, err := crypto.HexToECDSA(*sk1)
	if err != nil {
		userSk, err = crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("user sk:", userSk)

	// pubkey to address
	b_user_addr := crypto.PubkeyToAddress(userSk.PublicKey).Bytes()
	methodType := "EcdsaSecp256k1RecoveryMethod2020"
	// publicKeyData := crypto.CompressPubkey(&userSk.PublicKey)
	// methodType := "EcdsaSecp256k1VerificationKey2019"
	// did := hex.EncodeToString(crypto.Keccak256(publicKeyData))
	did := "f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f15"
	var nonceBuf = make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBuf, 0)

	// make msg for sign
	message := string("createDID") + did + methodType + string(b_user_addr) + string(nonceBuf)
	hash := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
	// sign
	signature, err := crypto.Sign(hash, userSk)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("hash:", hash)
	fmt.Println("user addr:", crypto.PubkeyToAddress(userSk.PublicKey))
	fmt.Println("msg:", message)
	fmt.Println("signature:", hex.EncodeToString(signature))

	fmt.Println("call proxy.CreateDID")

	tx, err := proxyIns.CreateDID(txAuth, did, methodType, b_user_addr, signature, big.NewInt(10001))
	// tx, err := didIns.CreateDIDByAdmin(txAuth, did, methodType, publicKeyData, big.NewInt(100002))
	if err != nil {
		log.Fatal(err)
	}

	err = com.CheckTx(eth, tx.Hash(), "createDid")
	if err != nil {
		log.Fatal(err)
	}

	// get did number
	num, err := proxyIns.Number(&bind.CallOpts{From: com.AdminAddr}, did)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("did number:", num)
}

func reverse(s []byte) []byte {

	newS := make([]byte, len(s))
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		newS[i], newS[j] = s[j], s[i]
	}
	return newS
}
