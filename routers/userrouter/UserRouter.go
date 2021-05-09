package userroute

import (
	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/controllers/usercontroller"
)

type UserRouter struct {
	crtl *usercontroller.UserController
}

func NewUserRouter(crtlInject *usercontroller.UserController) *UserRouter {
	return &UserRouter{crtl: crtlInject}
}

func (r *UserRouter) RouteRegister(route *gin.Engine) {
	route.POST("/login", r.crtl.Login)
}
