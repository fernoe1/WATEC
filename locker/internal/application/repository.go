package application

import (
	"context"

	"github.com/fernoe1/WATEC/locker/internal/domain"
)

type LockerRepository interface {
	Create(ctx context.Context, locker *domain.Locker) error
	Read(ctx context.Context, number int64) (*domain.Locker, error)
	Update(ctx context.Context, locker *domain.Locker) (*domain.Locker, error)
	Delete(ctx context.Context, number int64) error
}

type InMemoryLockerRepository interface {
	Set(ctx context.Context, locker *domain.Locker) error
	Get(ctx context.Context, number int64) (*domain.Locker, error)
	Delete(ctx context.Context, number int64) error
}
