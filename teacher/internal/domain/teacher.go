package domain

type Teacher struct {
	Name string `gorm:"primaryKey"`
	Free []Free `gorm:"foreignKey:TeacherName;references:Name;constraint:OnDelete:CASCADE"`
}

type Free struct {
	ID          int64  `gorm:"primaryKey"`
	TeacherName string `gorm:"index"`
	RoomNumber  int64
	From        int64
	To          int64
}
