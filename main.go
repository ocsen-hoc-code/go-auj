package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/builder"
	"github.com/ocsen-hoc-code/go-auj/models/service"
)

func SetupServer(serv service.Service) service.Service {
	service.NewService(serv)
	container := builder.BuildContainer()
	container.Invoke(func(serv *service.Service) {})
	return serv
}

func main() {
	serv := SetupServer(
		service.Service{
			Server: gin.Default(),
			Port:   8888})
	serv.Run()
}
