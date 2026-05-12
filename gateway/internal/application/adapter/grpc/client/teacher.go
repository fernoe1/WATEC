package client

import (
	"context"

	tchersvc "github.com/fernoe1/protogen/watec/service/teacher"
	"google.golang.org/grpc"
)

type TeacherClient struct {
	c tchersvc.TeacherServiceClient
}

func (t *TeacherClient) Create(ctx context.Context, in *tchersvc.CreateRequest, opts ...grpc.CallOption) (*tchersvc.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeacherClient) Read(ctx context.Context, in *tchersvc.ReadRequest, opts ...grpc.CallOption) (*tchersvc.ReadResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeacherClient) Update(ctx context.Context, in *tchersvc.UpdateRequest, opts ...grpc.CallOption) (*tchersvc.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TeacherClient) Delete(ctx context.Context, in *tchersvc.DeleteRequest, opts ...grpc.CallOption) (*tchersvc.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}
