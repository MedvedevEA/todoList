package apiserver

import (
	"os"
	"os/signal"
	"syscall"
	"todolist/internal/controller"
	"todolist/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type ApiServer struct {
	server      *fiber.App
	bindAddress string
}

func New(bindAddress string, service *service.Service) *ApiServer {
	app := fiber.New()
	app.Use(recover.New())
	controller.Init(app, service)

	return &ApiServer{
		server:      app,
		bindAddress: bindAddress,
	}
}
func (a *ApiServer) Run() error {

	chError := make(chan error, 1)
	go func() {
		if err := a.server.Listen(a.bindAddress); err != nil {
			chError <- err
		}
	}()
	go func() {
		chQuit := make(chan os.Signal, 1)
		signal.Notify(chQuit, syscall.SIGINT, syscall.SIGTERM)
		<-chQuit
		chError <- a.server.Shutdown()
	}()

	return <-chError
}
