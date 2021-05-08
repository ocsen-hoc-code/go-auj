package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/interfaces"
	hello2route "github.com/ocsen-hoc-code/go-auj/routers/hello2router"
	helloroute "github.com/ocsen-hoc-code/go-auj/routers/hellorouter"
)

type Routers struct {
	routes []interfaces.IRoute
}

func NewRouters(helloRoute *helloroute.HelloRouter, hello2Route *hello2route.Hello2Router) *Routers {
	return &Routers{routes: []interfaces.IRoute{helloRoute, hello2Route}}
}

func (r *Routers) Register(route *gin.Engine) {
	fmt.Print(len(r.routes))
	for _, v := range r.routes {
		v.RouteRegister(route)
	}
}
