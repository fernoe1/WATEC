package grpc

import (
	"context"

	"github.com/fernoe1/WATEC/teacher/internal/application"
	"github.com/fernoe1/WATEC/teacher/internal/domain"
	tchersvc "github.com/fernoe1/protogen/watec/service/teacher"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TeacherService struct {
	tchersvc.UnimplementedTeacherServiceServer
	uc application.TeacherUsecase
}

func NewTeacherService(uc application.TeacherUsecase) *TeacherService {
	return &TeacherService{uc: uc}
}

func toDomainTeacher(name string, free []*tchersvc.Free) *domain.Teacher {
	teacher := &domain.Teacher{Name: name}
	for _, v := range free {
		if v == nil {
			continue
		}
		teacher.Free = append(teacher.Free, domain.Free{
			TeacherName: name,
			RoomNumber:  v.GetRoomNumber(),
			From:        v.GetFrom(),
			To:          v.GetTo(),
		})
	}
	return teacher
}

func toProtoFree(free *domain.Free) *tchersvc.Free {
	if free == nil {
		return nil
	}
	return &tchersvc.Free{
		RoomNumber: free.RoomNumber,
		From:       free.From,
		To:         free.To,
	}
}

func (t *TeacherService) Create(ctx context.Context, request *tchersvc.CreateRequest) (*tchersvc.CreateResponse, error) {
	teacher := toDomainTeacher(request.GetName(), request.GetFree())
	if err := t.uc.Create(ctx, teacher); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &tchersvc.CreateResponse{}, nil
}

func (t *TeacherService) Read(ctx context.Context, request *tchersvc.ReadRequest) (*tchersvc.ReadResponse, error) {
	err, free := t.uc.Read(ctx, request.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &tchersvc.ReadResponse{IsFree: free != nil, Free: toProtoFree(free)}, nil
}

func (t *TeacherService) Update(ctx context.Context, request *tchersvc.UpdateRequest) (*tchersvc.UpdateResponse, error) {
	teacher := toDomainTeacher(request.GetName(), request.GetFree())
	if err, _ := t.uc.Update(ctx, teacher); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &tchersvc.UpdateResponse{}, nil
}

func (t *TeacherService) Delete(ctx context.Context, request *tchersvc.DeleteRequest) (*tchersvc.DeleteResponse, error) {
	if err := t.uc.Delete(ctx, request.GetName()); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &tchersvc.DeleteResponse{}, nil
}

func (t *TeacherService) mustEmbedUnimplementedTeacherServiceServer() {}
