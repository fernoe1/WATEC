package application

import (
	"context"

	"github.com/fernoe1/WATEC/teacher/internal/domain"
)

type TeacherUsecase interface {
	Create(ctx context.Context, teacher *domain.Teacher) error
	Read(ctx context.Context, name string) (error, *domain.Free)
	Update(ctx context.Context, teacher *domain.Teacher) (error, *domain.Teacher)
	Delete(ctx context.Context, name string) error
}
