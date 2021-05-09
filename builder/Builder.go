package builder

import (
	"github.com/ocsen-hoc-code/go-auj/controllers/hellocontroller"
	"github.com/ocsen-hoc-code/go-auj/controllers/taskcontroller"
	"github.com/ocsen-hoc-code/go-auj/controllers/usercontroller"
	"github.com/ocsen-hoc-code/go-auj/models/config"
	"github.com/ocsen-hoc-code/go-auj/models/service"
	taskrepository "github.com/ocsen-hoc-code/go-auj/reponsitories/task"
	userrepository "github.com/ocsen-hoc-code/go-auj/reponsitories/user"
	helloroute "github.com/ocsen-hoc-code/go-auj/routers/hellorouter"
	"github.com/ocsen-hoc-code/go-auj/routers/routers"
	taskroute "github.com/ocsen-hoc-code/go-auj/routers/taskrouter"
	userroute "github.com/ocsen-hoc-code/go-auj/routers/userrouter"
	"github.com/ocsen-hoc-code/go-auj/services/taskservice"
	userservice "github.com/ocsen-hoc-code/go-auj/services/usersevice"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	//Register database config
	container.Provide(config.GetDbConfg)
	//Register JWT config
	container.Provide(config.GetJWTConfig)

	//Register Repositories
	container.Provide(userrepository.NewUserRepository)
	container.Provide(taskrepository.NewTaskRepository)
	container.Provide(service.GetService)

	//Register Services
	container.Provide(userservice.NewUserService)
	container.Provide(taskservice.NewTaskService)

	//Register controllers
	container.Provide(hellocontroller.NewHelloController)
	container.Provide(usercontroller.NewUserController)
	container.Provide(taskcontroller.NewTaskController)

	//Register Routers
	container.Provide(helloroute.NewHelloRouter)
	container.Provide(userroute.NewUserRouter)
	container.Provide(taskroute.NewTaskRouter)
	container.Provide(routers.NewRouters)

	return container
}
