package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/interfaces"
	helloroute "github.com/ocsen-hoc-code/go-auj/routers/hellorouter"
)

type Routers struct {
	routes []interfaces.IRoute
}

func NewRouters(helloRoute *helloroute.HelloRouter) *Routers {
	return &Routers{routes: []interfaces.IRoute{helloRoute}}
}

func (r *Routers) Register(route *gin.Engine) {
	r.routes[0].RouteRegister(route)
}
