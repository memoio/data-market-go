package dumper

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/data-market/internal/database"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"gorm.io/gorm"

	proxy "did-solidity/go-contracts/proxy"

	com "github.com/memoio/contractsv2/common"
	inst "github.com/memoio/contractsv2/go_contracts/instance"
)

type BuyReadEvent struct {
	MfileDid common.Hash `json:"mfileDid"`
	MemoDid  string      `json:"memoDid"`
}

// unpack log data and store into db
func (d *Dumper) HandleBuyRead(log types.Log) error {
	var out BuyReadEvent

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

	var fileID uint

	// get fileid from filedid
	err = d.db.Model(&database.File{}).
		Select("file_id").
		Where("file_did = ?", fileDid).
		First(&fileID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("file with file_did '%s' not found", fileDid)
		}
		return fmt.Errorf("failed to query file_id: %v", err)
	}

	// get the controller of this filedid
	controllerAddr, err := d.getController(fileDid)
	if err != nil {
		logger.Debug("get controller of this filedid failed: ", err)
		return err
	}

	logger.Debug("controller:", controllerAddr)

	// get owner of this controller did
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
		FileID:       fileID,
		MemoDID:      out.MemoDid,
		UserAddress:  "0x" + addressHex,
		OwnerAddress: owner,
		AddTime:      buyTime,
		AddType:      1, // 1 for buyRead, 2 for grantRead
	}

	// store db
	err = fileMemo.CreateFileMemo()
	if err != nil {
		logger.Debug("store AddNode error: ", err.Error())
		return err
	}

	return nil
}

// get controller with filedid
func (d *Dumper) getController(filedid string) (string, error) {
	// connect endpoint
	client, err := ethclient.DialContext(context.Background(), d.endpoint)
	if err != nil {
		return "", err
	}

	// get instance
	instIns, err := inst.NewInstance(d.instance_ADDR, client)
	if err != nil {
		return "", err
	}

	// get proxyAddr
	proxyAddr, err := instIns.Instances(&bind.CallOpts{}, com.TypeDidProxy)
	if err != nil {
		return "", err
	}

	// get proxyInst
	proxyIns, err := proxy.NewProxy(proxyAddr, client)
	if err != nil {
		return "", err
	}

	// call getController in proxy
	controller, err := proxyIns.GetController(&bind.CallOpts{}, filedid)
	if err != nil {
		return "", err
	}

	return controller, nil
}

// get the buy time with block number(buyRead onchain)
func (d *Dumper) getBuyTime(num uint64) (time.Time, error) {
	// 创建 RPC 客户端
	rpcClient, err := rpc.DialContext(context.Background(), d.endpoint)
	if err != nil {
		return time.Time{}, err
	}
	client := ethclient.NewClient(rpcClient)
	defer client.Close()

	// 获取区块信息
	block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(num))
	if err != nil {
		return time.Time{}, err
	}

	// 转换时间戳
	timestamp := time.Unix(int64(block.Time()), 0).UTC()
	// 输出结果
	//buyTime := fmt.Sprintf("Buy Time (UTC): %s\n", timestamp.Format(time.RFC3339))

	return timestamp, nil
}

// query filedid from db with topic
func (d *Dumper) GetFileDIDByTopic(topic string) (string, error) {
	var fileDID string

	// query filedid
	err := d.db.Model(&database.File{}).
		Where("file_did_topic = ?", topic).
		Select("file_did").
		First(&fileDID).
		Error

	if err != nil {
		return "", fmt.Errorf("failed to get fileDID by topic: %w", err)
	}
	return fileDID, nil
}

// get the owner of a controller
func (d *Dumper) getOnwer(controller string) (string, error) {

	eth := d.endpoint
	proxyAddr := d.proxy_ADDR

	// connect endpoint
	client, err := ethclient.DialContext(context.Background(), eth)
	if err != nil {
		log.Fatal(err)
	}

	// get proxy instance
	proxyIns, err := proxy.NewProxy(proxyAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	// get owner address from proxy
	owner, err := proxyIns.GetMasterKeyAddr(&bind.CallOpts{}, controller)
	if err != nil {
		return "", err
	}

	return owner.String(), nil
}
