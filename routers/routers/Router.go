package routers

import (
	"github.com/gin-gonic/gin"
	helloroute "github.com/ocsen-hoc-code/go-auj/routers/HelloRouter"
)

type IRouter interface {
	RouteRegister(route *gin.Engine)
}

type Routers struct {
	routes []IRouter
}

func NewRouters(helloRoute *helloroute.HelloRouter) *Routers {
	return &Routers{routes: []IRouter{helloRoute}}
}

func (r *Routers) Register(route *gin.Engine) {
	r.routes[0].RouteRegister(route)
}
