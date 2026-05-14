package handler

import (
	"net/http"

	"github.com/fernoe1/WATEC/gateway/internal/application/adapter/grpc/client"
	clsrmsvc "github.com/fernoe1/protogen/watec/service/classroom"
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
	var createReq clsrmsvc.CreateRequest
	if err := ctx.ShouldBindJSON(&createReq); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createResp, err := c.client.Create(ctx, &createReq)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createResp)
}

func (c *ClassroomHandler) Read(ctx *gin.Context) {
	var readReq clsrmsvc.ReadRequest
	if err := ctx.ShouldBindJSON(&readReq); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	readResp, err := c.client.Read(ctx, &readReq)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, readResp)
}

func (c *ClassroomHandler) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *ClassroomHandler) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
