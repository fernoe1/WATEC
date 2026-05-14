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
	return l.c.Create(ctx, in)
}

func (l *LockerClient) Read(ctx context.Context, in *lokrsvc.ReadRequest, opts ...grpc.CallOption) (*lokrsvc.ReadResponse, error) {
	return l.c.Read(ctx, in)
}

func (l *LockerClient) Update(ctx context.Context, in *lokrsvc.UpdateRequest, opts ...grpc.CallOption) (*lokrsvc.UpdateResponse, error) {
	return l.c.Update(ctx, in)
}

func (l *LockerClient) Delete(ctx context.Context, in *lokrsvc.DeleteRequest, opts ...grpc.CallOption) (*lokrsvc.DeleteResponse, error) {
	return l.c.Delete(ctx, in)
}
