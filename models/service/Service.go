package service

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/models/config"
	"github.com/ocsen-hoc-code/go-auj/routers/routers"
)

type Service struct {
	Port   int
	Server *gin.Engine
	Config *config.DbConfig
}

var once sync.Once
var serv Service

func NewService(info Service) *Service {
	serv = info
	return &serv
}

func GetService(router *routers.Routers) *Service {
	once.Do(func() {
		router.Register(serv.Server)
	})
	return &serv
}

func (s Service) Run() {
	s.Server.Run(fmt.Sprintf(":%d", serv.Port))
}
