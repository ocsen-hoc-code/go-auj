package hello2route

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/ocsen-hoc-code/go-auj/controllers/hello2controller"
)

type Hello2Router struct {
	crtl *controllers.Hello2Controller
}

func NewHello2Router(crtl *controllers.Hello2Controller) *Hello2Router {
	return &Hello2Router{crtl: crtl}
}

func (r *Hello2Router) RouteRegister(route *gin.Engine) {
	route.GET("/h2", r.crtl.Index)
}
