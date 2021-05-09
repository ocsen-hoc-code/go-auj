package taskroute

import (
	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/controllers/taskcontroller"
	"github.com/ocsen-hoc-code/go-auj/middlewares/authencation"
	"github.com/ocsen-hoc-code/go-auj/models/config"
)

type TaskRouter struct {
	crtl      *taskcontroller.TaskController
	jwtConfig *config.JWTConfig
}

func NewTaskRouter(jwtConfigInject *config.JWTConfig, crtlInject *taskcontroller.TaskController) *TaskRouter {
	return &TaskRouter{crtl: crtlInject, jwtConfig: jwtConfigInject}
}

func (r *TaskRouter) RouteRegister(route *gin.Engine) {
	route.GET("/tasks", authencation.Authentication(r.jwtConfig.SecretKey), r.crtl.View)
	route.POST("/tasks", authencation.Authentication(r.jwtConfig.SecretKey), r.crtl.Create)
}
