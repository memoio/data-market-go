package database

type MemoDID struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	MemoDID     string `gorm:"column:memo_did;type:TEXT;not null"`
	UserAddress string `gorm:"column:user_address;type:TEXT;not null"`
}

func (MemoDID) TableName() string {
	return "memodid"
}
