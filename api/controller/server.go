package controller

import (
	"os"

	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/controller/handler"
	"github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/controller/middleware"
	"golang.org/x/exp/slog"

	"github.com/gin-gonic/gin"
)

const (
	apiVersion = "/api/v1"
)

func NewServer() (*gin.Engine, error) {
	r := gin.Default()
	opts := middleware.ServerLogJsonOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelInfo,
		},
		Indent: 4,
	}
	loghandler := middleware.NewServerLogJsonHandler(os.Stdout,opts)
	logger := slog.New(loghandler)

	// setting a CORS
	// setting a logger
	r.Use(middleware.Cors(),middleware.Logger(logger))

	tag := r.Group(apiVersion)

	{
		systemHandler := handler.NewSystemHandler()
		tag.GET("/system/health",systemHandler.Health)
	}

	return r,nil
}