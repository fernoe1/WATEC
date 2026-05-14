package handler

import (
	"github.com/fernoe1/WATEC/gateway/internal/application/adapter/grpc/client"
	"github.com/gin-gonic/gin"
)

type LockerHandler struct {
	group  *gin.RouterGroup
	client *client.LockerClient
}

func NewLockerHandler(
	rg *gin.RouterGroup,
	c *client.LockerClient,
) *LockerHandler {
	return &LockerHandler{group: rg, client: c}
}

func (l *LockerHandler) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerHandler) Read(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerHandler) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (l *LockerHandler) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
