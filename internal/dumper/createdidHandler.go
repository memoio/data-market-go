package dumper

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"

	did "did-solidity/go-contracts/did"

	"github.com/data-market/internal/database"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CreateDIDEvent struct {
	DID string
}

// unpack log data and store into db
func (d *Dumper) HandleCreateDID(log types.Log) error {
	var out CreateDIDEvent

	// unpack createdid
	err := d.unpack(log, d.accountdid_ABI, &out)
	if err != nil {
		return err
	}

	logger.Info("memodid:", out.DID)
	logger.Info("out: ", out)

	addressHex, err := d.getAddrWithDID(out.DID)
	if err != nil {
		logger.Debug("get address with memodid failed: ", err)
		return err
	}

	logger.Debug("user address:", addressHex)

	// make object for db store
	memoDID := database.MemoDID{
		MemoDID:     out.DID,
		UserAddress: addressHex,
	}

	// store db
	err = memoDID.CreateMemoDID()
	if err != nil {
		logger.Debug("store AddNode error: ", err.Error())
		return err
	}

	return nil
}

// pubKeyToAddress 从公钥生成以太坊地址
func pubKeyToAddress(pubKey *ecdsa.PublicKey) common.Address {
	// 将公钥转换为未压缩格式（65字节）
	uncompressedPubKey := crypto.FromECDSAPub(pubKey)

	// 提取X和Y坐标（跳过04前缀）
	pubKeyBytes := uncompressedPubKey[1:]

	// 计算Keccak-256哈希
	hash := crypto.Keccak256(pubKeyBytes)

	// 取最后20字节作为地址
	return common.BytesToAddress(hash[len(hash)-20:])
}

// get user address with memodid
func (d *Dumper) getAddrWithDID(memodid string) (addr string, err error) {

	// get user address with memodid

	// connect chain
	client, err := ethclient.DialContext(context.Background(), d.endpoint)
	if err != nil {
		return "", err
	}

	// get instance
	didIns, err := did.NewAccountDid(d.accountdid_ADDR, client)
	if err != nil {
		return "", err
	}

	// get pubkey with memodid
	pubkey, err := didIns.GetMasterVerification(&bind.CallOpts{}, memodid)
	if err != nil {
		return "", err
	}
	logger.Debug("pubkey data:", pubkey.PubKeyData)

	var addressHex string

	// parse address from pubkey
	switch pubkey.MethodType {
	// for type 2020, pubkey data is the address
	case "EcdsaSecp256k1RecoveryMethod2020":
		// 将字节转换为小写的十六进制字符串（无0x前缀）
		addressHex = hex.EncodeToString(pubkey.PubKeyData)

	// for type 2019, the pubkey data is the compressed pubkey
	case "EcdsaSecp256k1VerificationKey2019":
		// 1. 解压缩公钥
		pubKey, err := crypto.DecompressPubkey(pubkey.PubKeyData)
		if err != nil {
			logger.Debug("解压缩公钥失败: " + err.Error())
			return "", err
		}

		// 2. 生成以太坊地址
		address := pubKeyToAddress(pubKey)

		// to string
		addressHex = address.String()

	default:
		logger.Debug("error pubkey.methodType, not EcdsaSecp256k1VerificationKey2019 or EcdsaSecp256k1RecoveryMethod2020")
	}

	return addressHex, nil
}
