package database

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
	return G_DB.Save(&daBlockNumber).Error
}

func GetBlockNumber() (int64, error) {
	var blockNumber BlockNumber
	err := G_DB.Model(&BlockNumber{}).First(&blockNumber).Error

	return blockNumber.BlockNumber, err
}
