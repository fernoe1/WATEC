package handler

import (
	"github.com/fernoe1/WATEC/gateway/internal/application/adapter/grpc/client"
	"github.com/gin-gonic/gin"
)

type TeacherHandler struct {
	group  *gin.RouterGroup
	client *client.TeacherClient
}

func NewTeacherHandler(
	rg *gin.RouterGroup,
	c *client.TeacherClient,
) *TeacherHandler {
	return &TeacherHandler{group: rg, client: c}
}

func (t *TeacherHandler) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *TeacherHandler) Read(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *TeacherHandler) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *TeacherHandler) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
