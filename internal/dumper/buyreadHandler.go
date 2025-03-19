package dumper

import (
	"github.com/data-market/internal/database"
	"github.com/ethereum/go-ethereum/core/types"
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

	addressHex, err := d.getAddrWithDID(out.MemoDID)
	if err != nil {
		logger.Debug("get address with memodid failed: ", err)
		return err
	}
	logger.Debug("get user address from memodid:", addressHex)

	// make object for db store
	fileMemo := database.FileMemo{
		FileDID:     out.FileDID,
		MemoDID:     out.MemoDID,
		UserAddress: addressHex,
		AddType:     1, // 1 for buyRead, 2 for grantRead
	}

	// store db
	err = fileMemo.CreateFileMemo()
	if err != nil {
		logger.Debug("store AddNode error: ", err.Error())
		return err
	}

	return nil
}
