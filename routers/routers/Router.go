package routers

import (
	"github.com/gin-gonic/gin"
	routers "github.com/ocsen-hoc-code/go-auj/routers/HelloRouter"
)

type Routers struct {
	helloRoute *routers.HelloRouter
}

func NewRouters(helloRoute *routers.HelloRouter) *Routers {
	return &Routers{helloRoute: helloRoute}
}

func (r *Routers) RouteRegister(route *gin.Engine) {
	r.helloRoute.RouteRegister(route)
}
