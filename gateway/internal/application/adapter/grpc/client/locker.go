package client

import (
	"context"

	lokrsvc "github.com/fernoe1/protogen/watec/service/locker"
	"google.golang.org/grpc"
)

type LockerClient struct {
	c lokrsvc.LockerServiceClient
}

func NewLockerClient(c lokrsvc.LockerServiceClient) *LockerClient {
	return &LockerClient{c: c}
}

func (l *LockerClient) Create(ctx context.Context, in *lokrsvc.CreateRequest, opts ...grpc.CallOption) (*lokrsvc.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerClient) Read(ctx context.Context, in *lokrsvc.ReadRequest, opts ...grpc.CallOption) (*lokrsvc.ReadResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerClient) Update(ctx context.Context, in *lokrsvc.UpdateRequest, opts ...grpc.CallOption) (*lokrsvc.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerClient) Delete(ctx context.Context, in *lokrsvc.DeleteRequest, opts ...grpc.CallOption) (*lokrsvc.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}
