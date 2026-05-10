package domain

type Classroom struct {
	RoomNumber int64  `gorm:"primaryKey"`
	Free       []Free `gorm:"foreignKey:ClassroomID;constraint:OnDelete:CASCADE"`
}

type Free struct {
	ID          int64 `gorm:"primaryKey"`
	ClassroomID int64 `gorm:"index"`
	From        int64
	To          int64
}
