package client

import (
	"context"

	tchersvc "github.com/fernoe1/protogen/watec/service/teacher"
	"google.golang.org/grpc"
)

type TeacherClient struct {
	c tchersvc.TeacherServiceClient
}

func NewTeacherClient(cc tchersvc.TeacherServiceClient) *TeacherClient {
	return &TeacherClient{c: cc}
}

func (t *TeacherClient) Create(ctx context.Context, in *tchersvc.CreateRequest, opts ...grpc.CallOption) (*tchersvc.CreateResponse, error) {
	return t.c.Create(ctx, in, opts...)
}

func (t *TeacherClient) Read(ctx context.Context, in *tchersvc.ReadRequest, opts ...grpc.CallOption) (*tchersvc.ReadResponse, error) {
	return t.c.Read(ctx, in, opts...)
}

func (t *TeacherClient) Update(ctx context.Context, in *tchersvc.UpdateRequest, opts ...grpc.CallOption) (*tchersvc.UpdateResponse, error) {
	return t.c.Update(ctx, in, opts...)
}

func (t *TeacherClient) Delete(ctx context.Context, in *tchersvc.DeleteRequest, opts ...grpc.CallOption) (*tchersvc.DeleteResponse, error) {
	return t.c.Delete(ctx, in, opts...)
}
