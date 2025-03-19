package dumper

import (
	"context"
	"log"
	"math/big"
	"testing"

	"did-solidity/go-contracts/filedid"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestBuyRead(t *testing.T) {
	// 配置以太坊客户端（连接到测试网络或本地节点）
	client, err := ethclient.Dial("https://testchain.metamemo.one:24180")
	if err != nil {
		t.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	// make auth to send transaction
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 管理员私钥（替换为实际的私钥，通常从安全的地方获取，不要硬编码）
	adminPrivateKey := "your_admin_private_key_hex_without_0x_prefix"
	privateKey, err := crypto.HexToECDSA(adminPrivateKey)
	if err != nil {
		t.Fatalf("Failed to parse private key: %v", err)
	}

	// 获取管理员地址（验证是否与adminAddr一致）
	derivedAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	adminAddr := common.HexToAddress("your_admin_address_here") // 替换为你的adminAddr
	if derivedAddr != adminAddr {
		t.Fatalf("Derived address %s does not match adminAddr %s", derivedAddr.Hex(), adminAddr.Hex())
	}

	// 创建交易选项
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		t.Fatalf("Failed to create transactor: %v", err)
	}

	// 配置交易参数
	opts.GasLimit = uint64(300000) // 根据合约需求调整
	opts.Context = context.Background()

	// 获取建议的Gas价格
	gasPrice, err := client.SuggestGasPrice(opts.Context)
	if err != nil {
		t.Fatalf("Failed to suggest gas price: %v", err)
	}
	opts.GasPrice = gasPrice

	// 获取当前Nonce
	nonce, err := client.PendingNonceAt(opts.Context, adminAddr)
	if err != nil {
		t.Fatalf("Failed to get nonce: %v", err)
	}
	opts.Nonce = big.NewInt(int64(nonce))

	// 假设已初始化合约实例fileIns

	// filedid contract address on test chain
	filedidAddr := common.HexToAddress("0x9d80DBcC46052B598107b3760E5EbDcd0DDDA858")
	// get filedid instance
	fileIns, err := filedid.NewFileDid(filedidAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	// 调用buyRead方法
	mfileDid := "example_mfileDid"
	memoDid := "example_memoDid"
	tx, err := fileIns.BuyRead(opts, mfileDid, memoDid)
	if err != nil {
		t.Fatalf("Failed to send transaction: %v", err)
	}

	t.Logf("Transaction submitted! Tx Hash: %s", tx.Hash().Hex())

	// 等待交易确认（可选，测试中可能不需要）
	receipt, err := bind.WaitMined(opts.Context, client, tx)
	if err != nil {
		t.Fatalf("Failed to mine transaction: %v", err)
	}
	if receipt.Status != 1 {
		t.Fatalf("Transaction failed! Status: %d", receipt.Status)
	}

	t.Log("Transaction confirmed successfully!")
}
