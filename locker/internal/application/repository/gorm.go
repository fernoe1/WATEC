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
	//TODO implement me
	panic("implement me")
}

func (g *GormRepository) Read(ctx context.Context, number int64) (*domain.Locker, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormRepository) Update(ctx context.Context, locker *domain.Locker) (*domain.Locker, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormRepository) Delete(ctx context.Context, number int64) error {
	//TODO implement me
	panic("implement me")
}
