package handler

import (
	"github.com/fernoe1/WATEC/gateway/internal/application/adapter/grpc/client"
	"github.com/gin-gonic/gin"
)

type ClassroomHandler struct {
	group  *gin.RouterGroup
	client *client.ClassroomClient
}

func NewClassroomHandler(
	rg *gin.RouterGroup,
	c *client.ClassroomClient,
) *ClassroomHandler {
	return &ClassroomHandler{group: rg, client: c}
}

func (c *ClassroomHandler) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *ClassroomHandler) Read(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *ClassroomHandler) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *ClassroomHandler) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
