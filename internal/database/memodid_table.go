package database

type MemoDID struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	MemoDID     string `gorm:"column:memo_did;type:TEXT;not null"`
	UserAddress string `gorm:"column:user_address;type:TEXT;not null"`
}

func (MemoDID) TableName() string {
	return "memodid"
}

// store node info to db
func (did *MemoDID) CreateMemoDID() error {
	// store memodid
	return G_DB.Create(did).Error
}
