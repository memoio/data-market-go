package main

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"encoding/binary"
// 	"encoding/hex"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	// cfdid "did-solidity/go-contracts/controlfiledid"
// 	// "did-solidity/go-contracts/proxy"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	com "github.com/memoio/contractsv2/common"
// 	inst "github.com/memoio/contractsv2/go_contracts/instance"
// 	"github.com/memoio/did-solidity/go-contracts/proxy"
// )

// var (
// 	eth      string
// 	hexSk    string
// 	sks      [5]string
// 	as       [5]common.Address
// 	endPoint string

// 	hash  []byte
// 	signs [5][]byte

// 	txfee = big.NewInt(1e12)

// 	bigone       = big.NewInt(1)
// 	defaultNonce = big.NewInt(0)
// 	set          = true

// 	allGas = make([]uint64, 0)

// 	methodType = "EcdsaSecp256k1VerificationKey2019"

// 	// user sk and did, as the controller of this mfiledid
// 	// user_sk = "11f797550cd4d77d08fd160047f9d55c8f468260c87e53a1f74505de4d9454be"
// 	// did = "f3053946d7fcb75e380f8e4151ded1456abe67dd7607101fdd9cc19c0d1b3f18"

// 	// file did
// 	fdid        = "bafkreidfepr4j4wf6ooinzhncx6qucqlvaxspl3idyuite2aqlncyh25t4"
// 	encode      = "mid"
// 	ftype       = uint8(0)        // 0:private; 1:public
// 	price       = big.NewInt(100) // attomemo
// 	keywords    = []string{"picture", "animal", "dog", "white", "big"}
// 	newPrice    = big.NewInt(200)
// 	newFtype    = uint8(1)
// 	newKeywords = []string{"picture", "animal", "dog", "white", "big", "smile"}

// 	instanceAddr common.Address
// )

// //go run register.go -eth=test -sk=0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4

// func main() {
// 	chain := flag.String("eth", "dev", "eth api Address;") //dev test or product
// 	// TODO: read sk from local config
// 	sk := flag.String("sk", "", "signature for sending transaction")

// 	flag.Parse()

// 	// get instance address and endpoint with chain
// 	instanceAddr, eth = com.GetInsEndPointByChain(*chain)
// 	hexSk = *sk
// 	fmt.Println("instance address: ", instanceAddr)

// 	// tx sk is a must
// 	if len(hexSk) == 0 {
// 		log.Fatal("please input sk")
// 	}

// 	fmt.Println()

// 	// connect endpoint
// 	client, err := ethclient.DialContext(context.Background(), eth)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// get chain id
// 	chainId, err := client.ChainID(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// make auth
// 	txAuth, err := com.MakeAuth(chainId, hexSk)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_ = txAuth

// 	// new instanceIns
// 	instanceIns, err := inst.NewInstance(instanceAddr, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// get proxy address
// 	proxyAddr, _ := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeDidProxy)
// 	fmt.Println("proxy address:", proxyAddr.Hex())

// 	// proxyIns, err := proxy.NewProxy(common.HexToAddress("0x98b10003642E517C8eCc1cee052E31Bb1C4d1d29"), client)
// 	proxyIns, err := proxy.NewProxy(proxyAddr, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_ = proxyIns

// 	sk1, err := crypto.GenerateKey()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	sk2, err := crypto.GenerateKey()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	did1, err := createDID(sk1, proxyIns, txAuth)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(did1)

// 	did2, err := createDID(sk2, proxyIns, txAuth)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(did2)

// 	err = registerMfileDID(sk1, fdid, encode, ftype, did1, price, keywords, proxyIns, txAuth)
// 	if err != nil {
// 		log.Fatal("register mfile did", err)
// 	}

// 	// sk1, err := crypto.HexToECDSA("585300dd6882289845f7065fa7b3783b838c263d0a3eb95b8171a1f973e21f62")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// sk2, err := crypto.HexToECDSA("99e961aa181ca67c4800754e4c6bbafbe36ed652fa3ba53fb6ce98ee3797c19f")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// did1 := "a9388bebec991366053dd60b30fc7139b8a102fc3125c116bfcbbd692c0b9614"
// 	// did2 := "7fd5490de7902fe19f1eb11cdb8175d4b3ea0415edb49b12a45e4daf82e957e5"

// 	log.Println("sk1: ", hex.EncodeToString(crypto.FromECDSA(sk1)))
// 	log.Println("sk2: ", hex.EncodeToString(crypto.FromECDSA(sk2)))
// 	log.Println("DID1: ", did1)
// 	log.Println("DID2: ", did2)
// 	log.Println("file did: ", fdid)

// 	controller, err := proxyIns.GetController(&bind.CallOpts{}, fdid)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("file did's controller:", controller)

// 	nonce, err := proxyIns.GetFileDidNonce(&bind.CallOpts{}, controller)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("file did's nonce:", nonce)

// 	pk, err := proxyIns.GetVeri(&bind.CallOpts{}, did1, big.NewInt(0))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(pk)
// 	log.Println(crypto.PubkeyToAddress(sk1.PublicKey))
// 	log.Println(common.BytesToAddress(pk.PubKeyData))

// 	err = grantRead(sk1, fdid, did2, proxyIns, txAuth)
// 	if err != nil {
// 		log.Fatal("grant read", err)
// 	}
// 	log.Println("all success")
// }

// func createDID(sk *ecdsa.PrivateKey, proxyIns *proxy.Proxy, txAuth *bind.TransactOpts) (string, error) {
// 	publicKeyData := crypto.PubkeyToAddress(sk.PublicKey).Bytes()
// 	methodType := "EcdsaSecp256k1RecoveryMethod2020"
// 	// publicKeyData := crypto.CompressPubkey(&sk.PublicKey)
// 	// methodType := "EcdsaSecp256k1VerificationKey2019"
// 	did := hex.EncodeToString(crypto.Keccak256(append(publicKeyData, []byte("test")...)))
// 	var nonceBuf = make([]byte, 8)
// 	binary.BigEndian.PutUint64(nonceBuf, 0)

// 	message := string("createDID") + did + methodType + string(publicKeyData) + string(nonceBuf)
// 	hash := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
// 	signature, err := crypto.Sign(hash, sk)
// 	if err != nil {
// 		return did, err
// 	}

// 	tx, err := proxyIns.CreateDID(txAuth, did, methodType, publicKeyData, signature, big.NewInt(11000))
// 	// tx, err := didIns.CreateDIDByAdmin(txAuth, did, methodType, publicKeyData, big.NewInt(100002))
// 	if err != nil {
// 		return did, err
// 	}

// 	err = com.CheckTx(eth, tx.Hash(), "createDid")
// 	if err != nil {
// 		return did, err
// 	}

// 	return did, nil
// }

// func registerMfileDID(sk *ecdsa.PrivateKey, fdid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string, proxyIns *proxy.Proxy, txAuth *bind.TransactOpts) error {
// 	// ecdsaSk, err := crypto.HexToECDSA(sk)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	nonce, err := proxyIns.GetFileDidNonce(&bind.CallOpts{}, controller)
// 	if err != nil {
// 		return err
// 	}

// 	var nonceBuf = make([]byte, 8)
// 	binary.BigEndian.PutUint64(nonceBuf, nonce)

// 	priceBuf := make([]byte, 32)
// 	price.FillBytes(priceBuf)

// 	var typeBuf = make([]byte, 8)
// 	binary.BigEndian.PutUint64(typeBuf, 0)

// 	message := string("registerMfileDid") + fdid + encode + string(typeBuf[0]) + controller + string(priceBuf) + string(nonceBuf)
// 	for _, v := range keywords {
// 		message += v
// 	}
// 	message = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)

// 	hash := crypto.Keccak256([]byte(message))
// 	signature, err := crypto.Sign(hash, sk)
// 	if err != nil {
// 		return err
// 	}

// 	tx, err := proxyIns.RegisterMfileDid(txAuth, fdid, encode, ftype, controller, price, keywords, signature)
// 	if err != nil {
// 		return err
// 	}

// 	err = com.CheckTx(eth, tx.Hash(), "admin register file did")
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

// func grantRead(sk *ecdsa.PrivateKey, fdid string, did string, proxyIns *proxy.Proxy, txAuth *bind.TransactOpts) error {
// 	// ecdsaSk, err := crypto.HexToECDSA(sk)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	nonce, err := proxyIns.GetFileDidNonceByMfile(&bind.CallOpts{}, fdid)
// 	if err != nil {
// 		return err
// 	}
// 	log.Println(nonce)

// 	var nonceBuf = make([]byte, 8)
// 	binary.BigEndian.PutUint64(nonceBuf, nonce)

// 	message := string("grantRead") + fdid + did + string(nonceBuf)
// 	message = fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)

// 	hash := crypto.Keccak256([]byte(message))
// 	signature, err := crypto.Sign(hash, sk)
// 	if err != nil {
// 		return err
// 	}

// 	tx, err := proxyIns.GrantRead(txAuth, fdid, did, signature)
// 	if err != nil {
// 		return err
// 	}

// 	err = com.CheckTx(eth, tx.Hash(), "admin grant did read")
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
