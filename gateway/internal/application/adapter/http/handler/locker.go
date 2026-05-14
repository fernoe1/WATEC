package handler

import (
	"net/http"
	"strconv"

	"github.com/fernoe1/WATEC/gateway/internal/application/adapter/grpc/client"
	lokrsvc "github.com/fernoe1/protogen/watec/service/locker"
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
	var req lokrsvc.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := l.client.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (l *LockerHandler) Read(ctx *gin.Context) {
	number, err := strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid number"})
		return
	}

	resp, err := l.client.Read(ctx, &lokrsvc.ReadRequest{
		Number: int64(number),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (l *LockerHandler) Update(ctx *gin.Context) {
	var req lokrsvc.UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := l.client.Update(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (l *LockerHandler) Delete(ctx *gin.Context) {
	number, err := strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid number"})
		return
	}

	resp, err := l.client.Delete(ctx, &lokrsvc.DeleteRequest{
		Number: int64(number),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
