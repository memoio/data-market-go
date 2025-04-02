package dumper

import (
	"github.com/data-market/internal/database"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type GrantReadEvent struct {
	MfileDid common.Hash `json:"mfileDid"`
	MemoDid  string      `json:"memoDid"`
}

// unpack log data and store into db
func (d *Dumper) HandleGrantRead(log types.Log) error {
	var out GrantReadEvent

	// unpack createdid
	err := d.unpack(log, d.filedid_ABI, &out)
	if err != nil {
		return err
	}

	// hash to bytes to string
	fileDidTopic := hexutil.Encode(out.MfileDid.Bytes())

	logger.Debug("memodid:", out.MemoDid)
	logger.Debug("filedid topic:", fileDidTopic)

	logger.Debug("query filedid with topic")
	// todo: get filedid from file table with filedidTopic
	fileDid, err := d.GetFileDIDByTopic(fileDidTopic)
	if err != nil {
		return err
	}

	logger.Debug("filedid:", fileDid)

	// get the controller of this filedid
	controllerAddr, err := d.getController(fileDid)
	if err != nil {
		logger.Debug("get controller of this filedid failed: ", err)
		return err
	}

	logger.Debug("controller:", controllerAddr)

	owner, err := d.getOnwer(controllerAddr)
	if err != nil {
		logger.Debug("get owner of this controller failed: ", err)
		return err
	}

	// get the buyer address
	addressHex, err := d.getAddrWithDID(out.MemoDid)
	if err != nil {
		logger.Debug("get address with memodid failed: ", err)
		return err
	}
	logger.Debug("get user address from memodid:", addressHex)

	// get the block time(buy time)
	buyTime, err := d.getBuyTime(log.BlockNumber)
	if err != nil {
		logger.Debug("get buy time error: ", err.Error())
		return err
	}

	// make object for db store
	fileMemo := database.Access{
		FileDID:      fileDid,
		MemoDID:      out.MemoDid,
		UserAddress:  addressHex,
		OwnerAddress: owner,
		AddTime:      buyTime,
		AddType:      2, // 1 for buyRead, 2 for grantRead
	}

	// store db
	err = fileMemo.CreateFileMemo()
	if err != nil {
		logger.Debug("store AddNode error: ", err.Error())
		return err
	}

	return nil
}
