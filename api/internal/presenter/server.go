package presenter

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/o-ga09/api/internal/controller/system"
	"github.com/o-ga09/api/internal/controller/user"
	"github.com/o-ga09/api/internal/middleware"
)

const latest = "/v1"

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()
	v1 := r.Group(latest)

	// ロガーを設定
	logger := middleware.New()
	httpLogger := middleware.RequestLogger(logger)

	// ginを使用してリクエスト情報を取得する
	r.Use(httpLogger)

	// 死活監視用
	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health", systemHandler.Health)
	}

	// ユーザー管理機能
	{
		userHandler := user.NewUserHandler()
		v1.GET("", userHandler.GetUsers)
		v1.GET("/:id", userHandler.GetUserById)
		v1.POST("", userHandler.EditUser)
		v1.DELETE("/:id", userHandler.DeleteUser)
	}

	err := r.Run()
	if err != nil {
		return err
	}

	return nil
}

func NewServer() *Server {
	return &Server{}
}
