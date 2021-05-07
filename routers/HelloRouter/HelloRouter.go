package helloroute

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/ocsen-hoc-code/go-auj/controllers/hellocontroller"
)

type HelloRouter struct {
	helloCrtl *controllers.HelloController
}

func NewHelloRouter(helloCrtl *controllers.HelloController) *HelloRouter {
	return &HelloRouter{helloCrtl: helloCrtl}
}

func (r *HelloRouter) RouteRegister(route *gin.Engine) {
	route.GET("/", r.helloCrtl.Index)
}
