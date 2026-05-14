package grpc

import (
	"context"

	"github.com/fernoe1/WATEC/locker/internal/application"
	"github.com/fernoe1/WATEC/locker/internal/domain"
	lokrsvc "github.com/fernoe1/protogen/watec/service/locker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LockerService struct {
	lokrsvc.UnimplementedLockerServiceServer
	uc application.LockerUsecase
}

func NewLockerService(uc application.LockerUsecase) *LockerService {
	return &LockerService{uc: uc}
}

func (l *LockerService) Create(ctx context.Context, req *lokrsvc.CreateRequest) (*lokrsvc.CreateResponse, error) {
	locker := &domain.Locker{
		Number:     req.Number,
		BlockFloor: req.BlockFloor,
		MeshId:     req.MeshId,
	}

	if err := l.uc.Create(ctx, locker); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create locker: %v", err)
	}

	return &lokrsvc.CreateResponse{}, nil
}

func (l *LockerService) Read(ctx context.Context, req *lokrsvc.ReadRequest) (*lokrsvc.ReadResponse, error) {
	locker, err := l.uc.Read(ctx, req.Number)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read locker: %v", err)
	}
	if locker == nil {
		return nil, status.Errorf(codes.NotFound, "locker not found")
	}

	return &lokrsvc.ReadResponse{
		Number:     locker.Number,
		BlockFloor: locker.BlockFloor,
		MeshId:     locker.MeshId,
	}, nil
}

func (l *LockerService) Update(ctx context.Context, req *lokrsvc.UpdateRequest) (*lokrsvc.UpdateResponse, error) {
	locker := &domain.Locker{
		Number:     req.Number,
		BlockFloor: req.BlockFloor,
		MeshId:     req.MeshId,
	}

	if _, err := l.uc.Update(ctx, locker); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update locker: %v", err)
	}

	return &lokrsvc.UpdateResponse{}, nil
}

func (l *LockerService) Delete(ctx context.Context, req *lokrsvc.DeleteRequest) (*lokrsvc.DeleteResponse, error) {
	if err := l.uc.Delete(ctx, req.Number); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete locker: %v", err)
	}

	return &lokrsvc.DeleteResponse{}, nil
}

func (l *LockerService) mustEmbedUnimplementedLockerServiceServer() {}
