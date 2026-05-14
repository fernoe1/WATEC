package usecase

import (
	"context"
	"log/slog"

	"github.com/fernoe1/WATEC/classroom/internal/application"
	"github.com/fernoe1/WATEC/classroom/internal/domain"
)

type ClassroomUsecase struct {
	log *slog.Logger
	r   application.ClassroomRepository
	imr application.InMemoryRepository
}

func NewClassroomUsecase(
	log *slog.Logger,
	r application.ClassroomRepository,
	imr application.InMemoryRepository,
) *ClassroomUsecase {
	return &ClassroomUsecase{
		log: log,
		r:   r,
		imr: imr,
	}
}

func (c *ClassroomUsecase) Create(ctx context.Context, classroom *domain.Classroom) error {
	if err := c.r.Create(ctx, classroom); err != nil {
		return err
	}

	return c.imr.Del(ctx)
}

func (c *ClassroomUsecase) Read(ctx context.Context) ([]int64, error) {
	res, err := c.imr.Get(ctx)
	if err == nil && res != nil {
		c.log.InfoContext(ctx, "usecase.classroom.read.cache")
		return res, nil
	}

	free, err := c.r.Read(ctx)
	if err != nil {
		return nil, err
	}

	c.log.InfoContext(ctx, "usecase.classroom.read.db")

	if err := c.imr.Set(ctx, free); err != nil {
		c.log.ErrorContext(ctx, "usecase.classroom.read.cache", "error", err)
	}

	return free, nil
}

func (c *ClassroomUsecase) Update(ctx context.Context, classroom *domain.Classroom) (*domain.Classroom, error) {
	classroom, err := c.r.Update(ctx, classroom)
	if err != nil {
		return nil, err
	}

	if err := c.imr.Del(ctx); err != nil {
		c.log.ErrorContext(ctx, "usecase.classroom.update.cache", "error", err)
	}

	return classroom, nil
}

func (c *ClassroomUsecase) Delete(ctx context.Context, roomNumber int64) error {
	if err := c.r.Delete(ctx, roomNumber); err != nil {
		return err
	}

	return c.imr.Del(ctx)
}
