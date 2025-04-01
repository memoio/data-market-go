package dumper

import (
	"context"
	"did-solidity/go-contracts/proxy"
	"log"
	"math/big"
	"time"

	"github.com/data-market/internal/database"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type RegisterFileDIDEvent struct {
	FileDID string
}

// unpack log data and store into db
func (d *Dumper) HandleRegisterFileDID(log types.Log) error {
	var out RegisterFileDIDEvent

	// unpack
	err := d.unpack(log, d.filedid_ABI, &out)
	if err != nil {
		return err
	}

	logger.Debug("filedid: ", out.FileDID)
	logger.Debug("out: ", out)

	// 获取当前时间
	now := time.Now()

	// get controller and price
	controller, price, err := d.getControllerAndPrice(out.FileDID)
	if err != nil {
		return err
	}

	logger.Debug("controller:", controller)

	// update file info, locate the file with etag
	result := database.G_DB.Model(&database.File{}).
		Where("e_tag = ?", out.FileDID). // filedid is the same as the etag of this file
		Updates(database.File{
			FileDID:       out.FileDID,
			ControllerDID: controller,
			PublishState:  1,    // 设置为已上架
			PublishTime:   &now, // 设置为当前时间
			Price:         price.String(),
		})

	if result.Error != nil {
		return result.Error
	}

	// 检查是否确实更新了记录
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// get controller and price with file did
func (d *Dumper) getControllerAndPrice(fid string) (string, *big.Int, error) {

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

	controller, err := proxyIns.GetController(&bind.CallOpts{}, fid)
	if err != nil {
		return "", nil, err
	}

	price, err := proxyIns.GetPrice(&bind.CallOpts{}, fid)
	if err != nil {
		return "", nil, err
	}

	return controller, price, nil
}

// // get owner of this controller
// func (d *Dumper) getOnwer(controller string) (string, error) {

// 	eth := d.endpoint
// 	proxyAddr := d.proxy_ADDR

// 	// connect endpoint
// 	client, err := ethclient.DialContext(context.Background(), eth)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// get proxy instance
// 	proxyIns, err := proxy.NewProxy(proxyAddr, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// get owner address from proxy
// 	owner, err := proxyIns.GetMasterKeyAddr(&bind.CallOpts{}, controller)
// 	if err != nil {
// 		return "", err
// 	}

// 	return owner.String(), nil
// }
