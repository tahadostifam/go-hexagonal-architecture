package user_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samverrall/hex-structure/internal/core/domain/user"
	user_service "github.com/samverrall/hex-structure/internal/core/services/user"
)

func CreateAccount(userApi user_service.Api) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		var input struct {
			Name     string `json:"name"`
			Username string `json:"username"`
			Age      int    `json:"age"`
		}
		if err := c.BodyParser(&input); err != nil {
			return c.SendStatus(400)
		}

		resp, err := userApi.CreateAccount(ctx, input.Name, input.Username, input.Age)
		if err != nil {
			return c.SendStatus(500)
		}

		var out struct {
			Code   int         `json:"code"`
			Msg    string      `json:"msg"`
			Detail interface{} `json:"detail"`
		}

		out.Code = 200
		out.Msg = "user created"
		out.Detail = struct {
			User *user.User `json:"user"`
		}{
			User: resp,
		}

		return c.JSON(out)
	}
}
