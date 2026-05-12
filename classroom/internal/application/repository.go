package application

import (
	"context"

	"github.com/fernoe1/WATEC/classroom/internal/domain"
)

type ClassroomRepository interface {
	Create(ctx context.Context, classroom *domain.Classroom) error
	Read(ctx context.Context) ([]int64, error)
	Update(ctx context.Context, classroom *domain.Classroom) (*domain.Classroom, error)
	Delete(ctx context.Context, roomNumber int64) error
}

type InMemoryRepository interface {
	Set(ctx context.Context, free []int64) error
	Get(ctx context.Context) ([]int64, error)
}
