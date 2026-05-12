package grpc

import (
	"context"

	"github.com/fernoe1/WATEC/teacher/internal/application"
	tchersvc "github.com/fernoe1/protogen/watec/service/teacher"
)

type TeacherService struct {
	tchersvc.UnimplementedTeacherServiceServer
	uc application.TeacherUsecase
}

func NewTeacherService(uc application.TeacherUsecase) *TeacherService {
	return &TeacherService{uc: uc}
}

func (t TeacherService) Create(ctx context.Context, request *tchersvc.CreateRequest) (*tchersvc.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TeacherService) Read(ctx context.Context, request *tchersvc.ReadRequest) (*tchersvc.ReadResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TeacherService) Update(ctx context.Context, request *tchersvc.UpdateRequest) (*tchersvc.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TeacherService) Delete(ctx context.Context, request *tchersvc.DeleteRequest) (*tchersvc.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TeacherService) mustEmbedUnimplementedTeacherServiceServer() {
}
