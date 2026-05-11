package grpc

import (
	"context"
	"log/slog"

	"github.com/fernoe1/WATEC/classroom/internal/application"
	"github.com/fernoe1/WATEC/classroom/internal/domain"
	clsrmsvc "github.com/fernoe1/protogen/watec/service/classroom"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClassroomService struct {
	clsrmsvc.UnimplementedClassroomServiceServer
	log *slog.Logger
	uc  application.ClassroomUsecase
}

func NewClassroomService(
	log *slog.Logger,
	uc application.ClassroomUsecase,
) *ClassroomService {
	return &ClassroomService{
		log: log,
		uc:  uc,
	}
}

func (c ClassroomService) Create(ctx context.Context, request *clsrmsvc.CreateRequest) (*clsrmsvc.CreateResponse, error) {
	classroom := &domain.Classroom{}
	roomNumber := request.GetRoomNumber()
	free := request.GetFree()

	classroom.RoomNumber = roomNumber
	for _, v := range free {
		classroom.Free = append(classroom.Free, domain.Free{
			RoomNumber: roomNumber,
			From:       v.From,
			To:         v.To,
		})
	}

	if err := c.uc.Create(ctx, classroom); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return nil, nil
}

func (c ClassroomService) Read(ctx context.Context, request *clsrmsvc.ReadRequest) (*clsrmsvc.ReadResponse, error) {
	free, err := c.uc.Read(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &clsrmsvc.ReadResponse{Free: free}, nil
}

func (c ClassroomService) Update(ctx context.Context, request *clsrmsvc.UpdateRequest) (*clsrmsvc.UpdateResponse, error) {
	classroom := &domain.Classroom{}
	roomNumber := request.GetRoomNumber()
	free := request.GetFree()

	classroom.RoomNumber = roomNumber
	for _, v := range free {
		classroom.Free = append(classroom.Free, domain.Free{
			RoomNumber: roomNumber,
			From:       v.From,
			To:         v.To,
		})
	}

	_, err := c.uc.Update(ctx, classroom)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return nil, nil
}

func (c ClassroomService) Delete(ctx context.Context, request *clsrmsvc.DeleteRequest) (*clsrmsvc.DeleteResponse, error) {
	roomNumber := request.GetRoomNumber()

	if err := c.uc.Delete(ctx, roomNumber); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return nil, nil
}

func (c ClassroomService) mustEmbedUnimplementedClassroomServiceServer() {
	//TODO implement me
	panic("implement me")
}
