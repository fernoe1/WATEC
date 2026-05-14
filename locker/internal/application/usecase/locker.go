package usecase

import (
	"context"

	"github.com/fernoe1/WATEC/locker/internal/application"
	"github.com/fernoe1/WATEC/locker/internal/domain"
)

type LockerUsecase struct {
	r   application.LockerRepository
	imr application.InMemoryLockerRepository
}

func NewLockerUsecase(
	r application.LockerRepository,
	imr application.InMemoryLockerRepository,
) *LockerUsecase {
	return &LockerUsecase{r: r, imr: imr}
}

func (l *LockerUsecase) Create(ctx context.Context, locker *domain.Locker) error {
	
	if err := l.r.Create(ctx, locker); err != nil {
		return err
	}

	_ = l.imr.Set(ctx, locker)

	return nil
}

func (l *LockerUsecase) Read(ctx context.Context, number int64) (*domain.Locker, error) {

	locker, err := l.imr.Get(ctx, number)
	if err == nil && locker != nil {
		return locker, nil
	}

	locker, err = l.r.Read(ctx, number)
	if err != nil {
		return nil, err
	}

	_ = l.imr.Set(ctx, locker)

	return locker, nil
}

func (l *LockerUsecase) Update(ctx context.Context, locker *domain.Locker) (*domain.Locker, error) {
	updated, err := l.r.Update(ctx, locker)
	if err != nil {
		return nil, err
	}

	_ = l.imr.Set(ctx, updated)

	return updated, nil
}

func (l *LockerUsecase) Delete(ctx context.Context, number int64) error {

	if err := l.r.Delete(ctx, number); err != nil {
		return err
	}

	return nil
}
