package repository

import (
	"context"
	"time"

	"github.com/fernoe1/WATEC/classroom/internal/domain"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (g *GormRepository) Create(ctx context.Context, classroom *domain.Classroom) error {
	return g.db.WithContext(ctx).Create(classroom).Error
}

func (g *GormRepository) Read(ctx context.Context) ([]int64, error) {
	now := time.Now().In(time.FixedZone("UTC+6", 6*3600))
	hour := int64(now.Hour())

	var roomNumbers []int64

	err := g.db.WithContext(ctx).
		Table("classrooms").
		Select("DISTINCT classrooms.room_number").
		Where(`
			EXISTS (
				SELECT 1 FROM frees
				WHERE frees.classroom_id = classrooms.room_number
				AND frees.from <= ?
				AND frees.to > ?
			)
		`, hour, hour).
		Scan(&roomNumbers).Error

	return roomNumbers, err
}

func (g *GormRepository) Update(ctx context.Context, classroom *domain.Classroom) (*domain.Classroom, error) {
	err := g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("classroom_id = ?", classroom.RoomNumber).
			Delete(&domain.Free{}).Error; err != nil {
			return err
		}

		for i := range classroom.Free {
			classroom.Free[i].ClassroomID = classroom.RoomNumber
		}

		if err := tx.Create(&classroom.Free).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return classroom, nil
}

func (g *GormRepository) Delete(ctx context.Context, roomNumber int64) error {
	return g.db.WithContext(ctx).
		Where("room_number = ?", roomNumber).
		Delete(&domain.Classroom{}).Error
}
