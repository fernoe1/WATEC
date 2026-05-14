package handler

import (
	"net/http"

	"github.com/fernoe1/WATEC/gateway/internal/application/adapter/nats/notification"
	notifpb "github.com/fernoe1/protogen/watec/shares/notification"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	group     *gin.RouterGroup
	publisher *notification.Publisher
}

func NewNotificationHandler(
	rg *gin.RouterGroup,
	p *notification.Publisher,
) *NotificationHandler {
	return &NotificationHandler{group: rg, publisher: p}
}

func (n *NotificationHandler) Post(ctx *gin.Context) {
	var notif notifpb.Notification
	if err := ctx.ShouldBindJSON(&notif); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := n.publisher.Publish(ctx, "mailjet", &notif)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
