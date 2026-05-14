package grpc

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

const (
	MaxRecvMsgSize = 12 << 20
	timeKeepalive  = 10 * time.Second
)

func NewGRPCClient(target string) *grpc.ClientConn {
	keepaliveClientParameters := keepalive.ClientParameters{
		Time:                timeKeepalive,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepaliveClientParameters),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxRecvMsgSize)),
	}

	client, err := grpc.NewClient(target, opts...)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
