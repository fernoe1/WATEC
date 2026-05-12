package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/fernoe1/WATEC/classroom/internal/domain"
	"gorm.io/gorm"
)

type GormRepository struct {
	log *slog.Logger
	g   *gorm.DB
}

func NewGormRepository(log *slog.Logger, g *gorm.DB) *GormRepository {
	return &GormRepository{log: log, g: g}
}

func (g *GormRepository) Create(ctx context.Context, classroom *domain.Classroom) error {
	return g.g.WithContext(ctx).Create(classroom).Error
}

func (g *GormRepository) Read(ctx context.Context) ([]int64, error) {
	now := time.Now().In(time.FixedZone("UTC+6", 6*3600))
	hour := int64(now.Hour())
	g.log.InfoContext(ctx, "repository.gorm.read", "now", now, "hour", hour)

	var roomNumbers []int64

	err := g.g.WithContext(ctx).
		Table("classrooms").
		Select("DISTINCT classrooms.room_number").
		Where(`
			EXISTS (
				SELECT 1 FROM frees
				WHERE frees.room_number = classrooms.room_number
				AND frees.from <= ?
				AND frees.to > ?
			)
		`, hour, hour).
		Scan(&roomNumbers).Error

	return roomNumbers, err
}

func (g *GormRepository) Update(ctx context.Context, classroom *domain.Classroom) (*domain.Classroom, error) {
	err := g.g.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("room_number = ?", classroom.RoomNumber).
			Delete(&domain.Free{}).Error; err != nil {
			return err
		}

		for i := range classroom.Free {
			classroom.Free[i].RoomNumber = classroom.RoomNumber
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
	return g.g.WithContext(ctx).
		Where("room_number = ?", roomNumber).
		Delete(&domain.Classroom{}).Error
}
