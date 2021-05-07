package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/routers/routers"
)

func NewService(router *routers.Routers) *gin.Engine {
	service := gin.Default()
	router.Register(service)
	return service
}
