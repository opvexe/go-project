package router

import (
	"github.com/gin-gonic/gin"
	"go-projecct/handler"
)

func Routers() *gin.Engine {

	e := gin.New()

	api := e.Group("/api")

	userHandler := handler.UserHandler{}
	api.GET("/home", userHandler.Get)

	return e
}