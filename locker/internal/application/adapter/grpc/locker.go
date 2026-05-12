package grpc

import (
	"context"

	"github.com/fernoe1/WATEC/locker/internal/application"
	lokrsvc "github.com/fernoe1/protogen/watec/service/locker"
)

type LockerService struct {
	lokrsvc.UnimplementedLockerServiceServer
	uc application.LockerUsecase
}

func NewLockerService(uc application.LockerUsecase) *LockerService {
	return &LockerService{uc: uc}
}

func (l *LockerService) Create(ctx context.Context, request *lokrsvc.CreateRequest) (*lokrsvc.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerService) Read(ctx context.Context, request *lokrsvc.ReadRequest) (*lokrsvc.ReadResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerService) Update(ctx context.Context, request *lokrsvc.UpdateRequest) (*lokrsvc.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerService) Delete(ctx context.Context, request *lokrsvc.DeleteRequest) (*lokrsvc.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerService) mustEmbedUnimplementedLockerServiceServer() {
}
