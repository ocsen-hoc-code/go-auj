package routers

import (
	"github.com/gin-gonic/gin"
	helloroute "github.com/ocsen-hoc-code/go-auj/routers/HelloRouter"
)

type Routers struct {
	helloRoute *helloroute.HelloRouter
}

func NewRouters(helloRoute *helloroute.HelloRouter) *Routers {
	return &Routers{helloRoute: helloRoute}
}

func (r *Routers) Register(route *gin.Engine) {
	r.helloRoute.RouteRegister(route)
}
