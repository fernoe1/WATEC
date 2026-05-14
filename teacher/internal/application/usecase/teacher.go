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
	if err := t.r.Create(ctx, teacher); err != nil {
		return err
	}

	return t.imr.Set(ctx, teacher)
}

func (t *TeacherUsecase) Read(ctx context.Context, name string) (error, *domain.Free) {
	err, free := t.imr.Get(ctx, name)
	if err != nil {
		return err, nil
	}
	if free != nil {
		return nil, free
	}

	err, free = t.r.Read(ctx, name)
	if err != nil {
		return err, nil
	}
	if free != nil {
		_ = t.imr.Set(ctx, &domain.Teacher{Name: name, Free: []domain.Free{*free}})
	}

	return nil, free
}

func (t *TeacherUsecase) Update(ctx context.Context, teacher *domain.Teacher) (error, *domain.Teacher) {
	err, updated := t.r.Update(ctx, teacher)
	if err != nil {
		return err, nil
	}

	if err := t.imr.Set(ctx, updated); err != nil {
		return err, nil
	}

	return nil, updated
}

func (t *TeacherUsecase) Delete(ctx context.Context, name string) error {
	if err := t.r.Delete(ctx, name); err != nil {
		return err
	}

	return t.imr.Set(ctx, &domain.Teacher{Name: name})
}
