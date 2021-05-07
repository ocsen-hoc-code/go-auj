package helloroute

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/ocsen-hoc-code/go-auj/controllers/hellocontroller"
)

type HelloRouter struct {
	crtl *controllers.HelloController
}

func NewHelloRouter(crtl *controllers.HelloController) *HelloRouter {
	return &HelloRouter{crtl: crtl}
}

func (r *HelloRouter) RouteRegister(route *gin.Engine) {
	route.GET("/", r.crtl.Index)
}
