package restful_adapter

import user_controller "github.com/samverrall/hex-structure/internal/adapters/primary/restful/controller/user"

func (a *App) initAppRoutes() {
	a.fiber.Post("/users/create_account", user_controller.CreateAccount(a.serviceApi.UserApi))
}
