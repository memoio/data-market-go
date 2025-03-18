package database

import "log"

var blockNumberKey = "block_number_key"

type BlockNumber struct {
	BlockNumberKey string `gorm:"primarykey;column:block_number_key"`
	BlockNumber    int64
}

func SetBlockNumber(blockNumber int64) error {
	var daBlockNumber = BlockNumber{
		BlockNumberKey: blockNumberKey,
		BlockNumber:    blockNumber,
	}

	if err := G_DB.Save(&daBlockNumber).Error; err != nil {
		log.Printf("写入失败: %v (SQL: %s)", err, G_DB.Dialector.Explain(G_DB.Statement.SQL.String(), G_DB.Statement.Vars...))
		return err
	}

	return nil
}

func GetBlockNumber() (int64, error) {
	var blockNumber BlockNumber
	err := G_DB.Model(&BlockNumber{}).First(&blockNumber).Error

	return blockNumber.BlockNumber, err
}
