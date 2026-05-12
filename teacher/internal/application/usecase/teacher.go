package usecase

import (
	"context"

	"github.com/fernoe1/WATEC/teacher/internal/application"
	"github.com/fernoe1/WATEC/teacher/internal/domain"
)

type TeacherUsecase struct {
	r   application.TeacherRepository
	imr application.InMemoryTeacherRepository
}

func NewTeacherUsecase(
	r application.TeacherRepository,
	imr application.InMemoryTeacherRepository,
) *TeacherUsecase {
	return &TeacherUsecase{r: r, imr: imr}
}

func (t *TeacherUsecase) Create(ctx context.Context, teacher *domain.Teacher) error {
	//TODO implement me
	panic("implement me")
}

func (t *TeacherUsecase) Read(ctx context.Context, name string) (error, *domain.Free) {
	//TODO implement me
	panic("implement me")
}

func (t *TeacherUsecase) Update(ctx context.Context, teacher *domain.Teacher) (error, *domain.Teacher) {
	//TODO implement me
	panic("implement me")
}

func (t *TeacherUsecase) Delete(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}
