package dumper

import (
	"context"
	"math/big"
	"time"

	"github.com/data-market/internal/database"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	proxy "did-solidity/go-contracts/proxy"

	com "github.com/memoio/contractsv2/common"
	inst "github.com/memoio/contractsv2/go_contracts/instance"
)

type BuyReadEvent struct {
	FileDID string
	MemoDID string
}

// unpack log data and store into db
func (d *Dumper) HandleBuyRead(log types.Log) error {
	var out BuyReadEvent

	// unpack createdid
	err := d.unpack(log, d.filedid_ABI, &out)
	if err != nil {
		return err
	}

	logger.Debug("memodid:", out.MemoDID)
	logger.Debug("filedid:", out.FileDID)

	// get the owner address with filedid
	ownerAddr, err := d.getOwner(out.FileDID)
	if err != nil {
		logger.Debug("get owner address with filedid failed: ", err)
		return err
	}

	// get the buyer address
	addressHex, err := d.getAddrWithDID(out.MemoDID)
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
		FileDID:      out.FileDID,
		MemoDID:      out.MemoDID,
		UserAddress:  addressHex,
		OwnerAddress: ownerAddr,
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

// get owner with filedid
func (d *Dumper) getOwner(filedid string) (string, error) {
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
	owner, err := proxyIns.GetController(&bind.CallOpts{}, filedid)
	if err != nil {
		return "", err
	}

	return owner, nil
}

// get the buy time(buyRead onchain)
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
