package repository

import (
	"context"
	"errors"
	"time"

	"github.com/fernoe1/WATEC/teacher/internal/domain"
	"gorm.io/gorm"
)

type GormRepository struct {
	g *gorm.DB
}

func NewGormRepository(g *gorm.DB) *GormRepository {
	return &GormRepository{g: g}
}

func currentHour() int64 {
	now := time.Now().In(time.FixedZone("UTC+6", 6*3600))
	return int64(now.Hour())
}

func setTeacherName(teacher *domain.Teacher) {
	for i := range teacher.Free {
		teacher.Free[i].TeacherName = teacher.Name
	}
}

func (g *GormRepository) Create(ctx context.Context, teacher *domain.Teacher) error {
	setTeacherName(teacher)
	return g.g.WithContext(ctx).Create(teacher).Error
}

func (g *GormRepository) Read(ctx context.Context, name string) (error, *domain.Free) {
	hour := currentHour()
	free := &domain.Free{}

	err := g.g.WithContext(ctx).
		Where("teacher_name = ? AND \"from\" <= ? AND \"to\" > ?", name, hour, hour).
		First(free).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return err, nil
	}

	return nil, free
}

func (g *GormRepository) Update(ctx context.Context, teacher *domain.Teacher) (error, *domain.Teacher) {
	setTeacherName(teacher)

	err := g.g.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&domain.Teacher{Name: teacher.Name}).Error; err != nil {
			return err
		}

		if err := tx.Where("teacher_name = ?", teacher.Name).Delete(&domain.Free{}).Error; err != nil {
			return err
		}

		if len(teacher.Free) > 0 {
			if err := tx.Create(&teacher.Free).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err, nil
	}

	return nil, teacher
}

func (g *GormRepository) Delete(ctx context.Context, name string) error {
	return g.g.WithContext(ctx).Where("name = ?", name).Delete(&domain.Teacher{}).Error
}
