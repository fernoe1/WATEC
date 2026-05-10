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

func (c *ClassroomUsecase) Create(ctx context.Context, classroom *domain.Classroom) error {
	return c.r.Create(ctx, classroom)
}

func (c *ClassroomUsecase) Read(ctx context.Context) ([]int64, error) {
	res, err := c.imr.Get(ctx)
	if res != nil {
		c.log.InfoContext(ctx, "read from cache")
		return res, nil
	}

	if err == nil {
		free, err := c.r.Read(ctx)
		c.log.InfoContext(ctx, "read from db")

		if err != nil {
			return nil, err
		}

		c.imr.Set(ctx, free)

		return free, nil
	}

	return nil, err
}

func (c *ClassroomUsecase) Update(ctx context.Context, classroom *domain.Classroom) (*domain.Classroom, error) {
	return c.r.Update(ctx, classroom)
}

func (c *ClassroomUsecase) Delete(ctx context.Context, roomNumber int64) error {
	return c.r.Delete(ctx, roomNumber)
}
