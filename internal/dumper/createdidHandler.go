package dumper

import (
	"github.com/data-market/internal/database"
	"github.com/ethereum/go-ethereum/core/types"
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

	// make object for db store
	memoDID := database.MemoDID{
		MemoDID:     out.DID,
		UserAddress: "test",
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
