package handler

import (
	"net/http"

	tchersvc "github.com/fernoe1/protogen/watec/service/teacher"
	"github.com/gin-gonic/gin"
)

type TeacherHandler struct {
	client tchersvc.TeacherServiceClient
}

type teacherFreeRequest struct {
	RoomNumber int64 `json:"roomNumber"`
	From       int64 `json:"from"`
	To         int64 `json:"to"`
}

type teacherWriteRequest struct {
	Name string               `json:"name" binding:"required"`
	Free []teacherFreeRequest `json:"free"`
}

func NewTeacherHandler(client tchersvc.TeacherServiceClient) *TeacherHandler {
	return &TeacherHandler{client: client}
}

func toProtoTeacherFree(free []teacherFreeRequest) []*tchersvc.Free {
	result := make([]*tchersvc.Free, 0, len(free))
	for _, v := range free {
		result = append(result, &tchersvc.Free{
			RoomNumber: v.RoomNumber,
			From:       v.From,
			To:         v.To,
		})
	}
	return result
}

func teacherName(ctx *gin.Context) string {
	if name := ctx.Param("name"); name != "" {
		return name
	}
	return ctx.Query("name")
}

func (t *TeacherHandler) Create(ctx *gin.Context) {
	var req teacherWriteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := t.client.Create(ctx.Request.Context(), &tchersvc.CreateRequest{
		Name: req.Name,
		Free: toProtoTeacherFree(req.Free),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "teacher created"})
}

func (t *TeacherHandler) Read(ctx *gin.Context) {
	name := teacherName(ctx)
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	res, err := t.client.Read(ctx.Request.Context(), &tchersvc.ReadRequest{Name: name})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (t *TeacherHandler) Update(ctx *gin.Context) {
	var req teacherWriteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Name == "" {
		req.Name = teacherName(ctx)
	}
	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	_, err := t.client.Update(ctx.Request.Context(), &tchersvc.UpdateRequest{
		Name: req.Name,
		Free: toProtoTeacherFree(req.Free),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "teacher updated"})
}

func (t *TeacherHandler) Delete(ctx *gin.Context) {
	name := teacherName(ctx)
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	_, err := t.client.Delete(ctx.Request.Context(), &tchersvc.DeleteRequest{Name: name})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "teacher deleted"})
}
