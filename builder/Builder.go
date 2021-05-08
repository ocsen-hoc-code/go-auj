package builder

import (
	"github.com/ocsen-hoc-code/go-auj/controllers/hello2controller"
	"github.com/ocsen-hoc-code/go-auj/controllers/hellocontroller"
	"github.com/ocsen-hoc-code/go-auj/models/service"
	hello2route "github.com/ocsen-hoc-code/go-auj/routers/hello2router"
	helloroute "github.com/ocsen-hoc-code/go-auj/routers/hellorouter"
	"github.com/ocsen-hoc-code/go-auj/routers/routers"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	//Register controllers
	container.Provide(hellocontroller.NewHelloController)
	container.Provide(hello2controller.NewHello2Controller)

	//Register Routers
	container.Provide(helloroute.NewHelloRouter)
	container.Provide(hello2route.NewHello2Router)

	container.Provide(routers.NewRouters)
	container.Provide(service.GetService)

	return container
}
