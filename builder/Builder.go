package builder

import (
	"github.com/ocsen-hoc-code/go-auj/controllers/hellocontroller"
	"github.com/ocsen-hoc-code/go-auj/models/service"
	helloroute "github.com/ocsen-hoc-code/go-auj/routers/hellorouter"
	"github.com/ocsen-hoc-code/go-auj/routers/routers"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(hellocontroller.NewHelloController)
	container.Provide(helloroute.NewHelloRouter)
	container.Provide(routers.NewRouters)
	container.Provide(service.NewService)
	return container
}
