package restful_adapter

import (
	user_controller "github.com/tahadostifam/go-hexagonal-architecture/internal/adapters/primary/restful/controller/user"
)

func (a *App) initAppRoutes() {
	a.fiber.Post("/users/create_account", user_controller.CreateAccount(a.servicesApi.UserApi))
}
