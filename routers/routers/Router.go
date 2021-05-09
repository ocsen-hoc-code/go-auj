package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/interfaces"
	helloroute "github.com/ocsen-hoc-code/go-auj/routers/hellorouter"
	taskroute "github.com/ocsen-hoc-code/go-auj/routers/taskrouter"
	userroute "github.com/ocsen-hoc-code/go-auj/routers/userrouter"
)

type Routers struct {
	routes []interfaces.IRoute
}

func NewRouters(helloRoute *helloroute.HelloRouter, userRoute *userroute.UserRouter, taskRoute *taskroute.TaskRouter) *Routers {
	return &Routers{routes: []interfaces.IRoute{helloRoute, userRoute, taskRoute}}
}

func (r *Routers) Register(route *gin.Engine) {
	for _, v := range r.routes {
		v.RouteRegister(route)
	}
}
