package restful_adapter

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	user_service "github.com/tahadostifam/go-hexagonal-architecture/internal/core/services/user"
)

type ServicesApi struct {
	UserApi user_service.Api
}

type App struct {
	fiber       *fiber.App
	port        int
	servicesApi ServicesApi
}

func NewApp(servicesApi ServicesApi, opts ...AppOption) *App {
	s := &App{
		fiber:       fiber.New(),
		servicesApi: servicesApi,
		port:        8000,
	}

	for _, applyOption := range opts {
		applyOption(s)
	}

	s.initAppRoutes()

	return s
}

func (a *App) Run() error {
	return a.fiber.Listen(fmt.Sprintf(":%d", a.port))
}
