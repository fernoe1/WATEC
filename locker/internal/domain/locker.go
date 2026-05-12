package domain

type Locker struct {
	Number     int64 `gorm:"primaryKey"`
	BlockFloor int64
	MeshId     int64
}
