package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/builder"
)

func SetupServer() {
	container := builder.BuildContainer()
	container.Invoke(func(serv *gin.Engine) {
		serv.Run(":8888")
	})
}

func main() {
	SetupServer()
}
