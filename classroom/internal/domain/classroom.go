package domain

type Classroom struct {
	RoomNumber int64  `gorm:"primaryKey"`
	Free       []Free `gorm:"foreignKey:RoomNumber;references:RoomNumber;constraint:OnDelete:CASCADE"`
}

type Free struct {
	ID         int64 `gorm:"primaryKey"`
	RoomNumber int64 `gorm:"index"`
	From       int64
	To         int64
}
