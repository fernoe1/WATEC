package repository

import (
	"context"

	"github.com/fernoe1/WATEC/teacher/internal/domain"
	"gorm.io/gorm"
)

type GormRepository struct {
	g *gorm.DB
}

func NewGormRepository(g *gorm.DB) *GormRepository {
	return &GormRepository{g: g}
}

func (g GormRepository) Create(ctx context.Context, teacher *domain.Teacher) error {
	//TODO implement me
	panic("implement me")
}

func (g GormRepository) Read(ctx context.Context, name string) (error, *domain.Free) {
	//TODO implement me
	panic("implement me")
}

func (g GormRepository) Update(ctx context.Context, teacher *domain.Teacher) (error, *domain.Teacher) {
	//TODO implement me
	panic("implement me")
}

func (g GormRepository) Delete(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}
