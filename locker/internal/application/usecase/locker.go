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
	//TODO implement me
	panic("implement me")
}

func (l *LockerUsecase) Read(ctx context.Context, number int64) (*domain.Locker, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerUsecase) Update(ctx context.Context, locker *domain.Locker) (*domain.Locker, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerUsecase) Delete(ctx context.Context, number int64) error {
	//TODO implement me
	panic("implement me")
}
