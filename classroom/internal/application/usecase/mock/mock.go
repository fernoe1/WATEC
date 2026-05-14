package mock

import (
	"context"

	"github.com/fernoe1/WATEC/classroom/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ClassroomRepository struct{ mock.Mock }

func (c *ClassroomRepository) Create(ctx context.Context, classroom *domain.Classroom) error {
	args := c.Called(ctx, classroom)
	return args.Error(0)
}

func (c *ClassroomRepository) Read(ctx context.Context) ([]int64, error) {
	args := c.Called(ctx)
	free, _ := args.Get(0).([]int64)
	return free, args.Error(1)
}

func (c *ClassroomRepository) Update(ctx context.Context, classroom *domain.Classroom) (*domain.Classroom, error) {
	args := c.Called(ctx, classroom)
	newClassroom, _ := args.Get(0).(*domain.Classroom)
	return newClassroom, args.Error(1)
}

func (c *ClassroomRepository) Delete(ctx context.Context, roomNumber int64) error {
	args := c.Called(ctx, roomNumber)
	return args.Error(0)
}

type InMemoryRepository struct{ mock.Mock }

func (i *InMemoryRepository) Del(ctx context.Context) error {
	args := i.Called(ctx)
	return args.Error(0)
}

func (i *InMemoryRepository) Set(ctx context.Context, free []int64) error {
	args := i.Called(ctx, free)
	return args.Error(0)
}

func (i *InMemoryRepository) Get(ctx context.Context) ([]int64, error) {
	args := i.Called(ctx)
	free, _ := args.Get(0).([]int64)
	return free, args.Error(1)
}
