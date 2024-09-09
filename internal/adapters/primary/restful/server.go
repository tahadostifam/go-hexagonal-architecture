package restful_adapter

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	user_service "github.com/samverrall/hex-structure/internal/core/services/user"
)

type ServiceApi struct {
	UserApi user_service.Api
}

type App struct {
	fiber      *fiber.App
	port       int
	serviceApi ServiceApi
}

func NewApp(serviceApi ServiceApi, opts ...AppOption) *App {
	s := &App{
		serviceApi: serviceApi,
		port:       8000,
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
