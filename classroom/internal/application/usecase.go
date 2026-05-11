package application

import (
	"context"

	"github.com/fernoe1/WATEC/classroom/internal/domain"
)

type ClassroomUsecase interface {
	Create(ctx context.Context, classroom *domain.Classroom) error
	Read(ctx context.Context) ([]int64, error)
	Update(ctx context.Context, classroom *domain.Classroom) (*domain.Classroom, error)
	Delete(ctx context.Context, roomNumber int64) error
}
