package server

import (
	"net/http"
	"time"

	"github.com/fernoe1/WATEC/gateway/internal/application/adapter/grpc/client"
	"github.com/fernoe1/WATEC/gateway/internal/application/adapter/http/handler"
	"github.com/fernoe1/WATEC/gateway/internal/application/adapter/http/middleware"
	"github.com/fernoe1/WATEC/gateway/pkg/grpc"
	clsrmsvc "github.com/fernoe1/protogen/watec/service/classroom"
	lokrsvc "github.com/fernoe1/protogen/watec/service/locker"
	tchersvc "github.com/fernoe1/protogen/watec/service/teacher"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func (s *Server) runHTTPServer() error {
	// to do: run this in another endpoint with cors configured.
	s.gin.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "up")
	})

	s.gin.Use(gin.Logger())
	s.gin.Use(gin.Recovery())
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

	classroom := s.gin.Group("/classroom")
	classroomClient := client.NewClassroomClient(
		clsrmsvc.NewClassroomServiceClient(grpc.NewGRPCClient(s.cfg.Server.ClassroomGrpcSrvAddr)),
	)
	classroomHandler := handler.NewClassroomHandler(classroom, classroomClient)
	classroomHandler.MapClassroomRoutes()

	locker := s.gin.Group("/locker")
	lockerClient := client.NewLockerClient(
		lokrsvc.NewLockerServiceClient(grpc.NewGRPCClient(s.cfg.Server.LockerGrpcSrvAddr)),
	)
	lockerHandler := handler.NewLockerHandler(locker, lockerClient)
	lockerHandler.MapLockerRoutes()

	teacher := s.gin.Group("/teacher")
	teacherClient := client.NewTeacherClient(
		tchersvc.NewTeacherServiceClient(grpc.NewGRPCClient(s.cfg.Server.TeacherGrpcSrvAddr)),
	)
	teacherHandler := handler.NewTeacherHandler(teacher, teacherClient)
	teacherHandler.MapTeacherRoutes()

	s.srv = &http.Server{
		Addr:              s.cfg.Http.Port,
		Handler:           s.gin,
		ReadTimeout:       s.cfg.Http.ReadTimeout,
		ReadHeaderTimeout: s.cfg.Http.ReadHeaderTimeout,
		WriteTimeout:      s.cfg.Http.WriteTimeout,
		IdleTimeout:       s.cfg.Http.IdleTimeout,
		MaxHeaderBytes:    s.cfg.Http.MaxHeaderBytes,
	}

	return s.srv.ListenAndServe()
}
