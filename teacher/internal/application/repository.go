package application

import (
	"context"

	"github.com/fernoe1/WATEC/teacher/internal/domain"
)

type TeacherRepository interface {
	Create(ctx context.Context, teacher *domain.Teacher) error
	Read(ctx context.Context, name string) (error, *domain.Free)
	Update(ctx context.Context, teacher *domain.Teacher) (error, *domain.Teacher)
	Delete(ctx context.Context, name string) error
}

type InMemoryTeacherRepository interface {
	Get(ctx context.Context, name string) (error, *domain.Free)
	Set(ctx context.Context, teacher *domain.Teacher) error
}
