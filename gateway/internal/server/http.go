package server

import (
	"net/http"
	"time"

	"github.com/fernoe1/WATEC/gateway/internal/application/adapter/http/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func (s *Server) runHTTPServer() error {
	// to do: run this in another endpoint with cors configured.
	s.gin.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "up")
	})

	s.gin.Use(otelgin.Middleware(s.cfg.Telemetry.Name))
	s.gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		MaxAge:           7 * time.Hour,
	}))
	s.gin.Use(middleware.RateLimit(s.redis))

	return nil
}
