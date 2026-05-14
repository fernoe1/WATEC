package client

import (
	"context"

	clsrmsvc "github.com/fernoe1/protogen/watec/service/classroom"
	"google.golang.org/grpc"
)

type ClassroomClient struct {
	c clsrmsvc.ClassroomServiceClient
}

func NewClassroomClient(c clsrmsvc.ClassroomServiceClient) *ClassroomClient {
	return &ClassroomClient{c: c}
}

func (c *ClassroomClient) Create(ctx context.Context, in *clsrmsvc.CreateRequest, opts ...grpc.CallOption) (*clsrmsvc.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClassroomClient) Read(ctx context.Context, in *clsrmsvc.ReadRequest, opts ...grpc.CallOption) (*clsrmsvc.ReadResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClassroomClient) Update(ctx context.Context, in *clsrmsvc.UpdateRequest, opts ...grpc.CallOption) (*clsrmsvc.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClassroomClient) Delete(ctx context.Context, in *clsrmsvc.DeleteRequest, opts ...grpc.CallOption) (*clsrmsvc.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}
