package dumper

import (
	"time"

	"github.com/data-market/internal/database"
	"github.com/ethereum/go-ethereum/core/types"
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

	// todo:
	// get etag from filedid
	// controller did with filedid
	// get price with filedid

	// 更新操作
	result := database.G_DB.Model(&database.File{}).
		Where("e_tag = ?", "eTag"). // 指定ETag条件
		Updates(database.File{
			FileDID:       "newFileDID",
			ControllerDID: "newControllerDID",
			PublishState:  1,    // 设置为已上架
			PublishTime:   &now, // 设置为当前时间
			Price:         "newPrice",
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
