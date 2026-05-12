package handler

import "github.com/gin-gonic/gin"

type CRUD interface {
	Create(ctx *gin.Context)
	Read(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
