package repository

import (
	"context"

	"github.com/fernoe1/WATEC/locker/internal/domain"
	"gorm.io/gorm"
)

type GormRepository struct {
	g *gorm.DB
}

func NewGormRepository(g *gorm.DB) *GormRepository {
	return &GormRepository{g: g}
}

func (g *GormRepository) Create(ctx context.Context, locker *domain.Locker) error {
	return g.g.WithContext(ctx).Create(locker).Error
}

func (g *GormRepository) Read(ctx context.Context, number int64) (*domain.Locker, error) {
	var locker domain.Locker

	err := g.g.WithContext(ctx).
		Where("number = ?", number).
		First(&locker).Error

	if err != nil {
		return nil, err
	}

	return &locker, nil
}

func (g *GormRepository) Update(ctx context.Context, locker *domain.Locker) (*domain.Locker, error) {
	err := g.g.WithContext(ctx).
		Save(locker).Error

	if err != nil {
		return nil, err
	}

	return locker, nil
}

func (g *GormRepository) Delete(ctx context.Context, number int64) error {
	return g.g.WithContext(ctx).
		Where("number = ?", number).
		Delete(&domain.Locker{}).Error
}
