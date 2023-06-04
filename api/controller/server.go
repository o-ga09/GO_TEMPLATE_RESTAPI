package controller

import (
	"main/api/controller/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	apiVersion = "/v1"
	langTagJa = "/ja"
	langTagEn = "/en"
)

func NewServer() (*gin.Engine, error) {
	r := gin.Default()

	//setting a CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowCredentials: false,
		MaxAge: 24 * time.Hour,
	}))

	tag := r.Group(apiVersion + langTagJa)

	{
		systemHandler := handler.NewSystemHandler()
		tag.GET("/system/health",systemHandler.Health)
	}

	return r,nil
}

func NewSystemHandler() {
	panic("unimplemented")
}