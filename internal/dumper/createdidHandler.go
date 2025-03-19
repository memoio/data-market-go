package dumper

import (
	"context"
	"encoding/hex"

	did "did-solidity/go-contracts/did"

	"github.com/data-market/internal/database"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	// get user address with memodid

	// connect chain
	client, err := ethclient.DialContext(context.Background(), d.endpoint)
	if err != nil {
		return err
	}

	// get instance
	didIns, err := did.NewAccountDid(d.accountdid_ADDR, client)
	if err != nil {
		return err
	}

	// get pubkey with memodid
	pubkey, err := didIns.GetMasterVerification(&bind.CallOpts{}, out.DID)
	if err != nil {
		return err
	}

	logger.Debug("pubkey data:", pubkey.PubKeyData)
	// 将字节转换为小写的十六进制字符串（无0x前缀）
	addressHex := hex.EncodeToString(pubkey.PubKeyData)
	addr := common.HexToAddress(addressHex)

	logger.Debug("user address:", addr)

	// make object for db store
	memoDID := database.MemoDID{
		MemoDID:     out.DID,
		UserAddress: addr.String(),
	}

	//logger.Info("============= store MemoDID..", memoDID)
	// store db
	err = memoDID.CreateMemoDID()
	if err != nil {
		logger.Debug("store AddNode error: ", err.Error())
		return err
	}

	// // test set online
	// database.SetOnline(nodeInfo.Address, nodeInfo.Id, true)

	return nil
}
