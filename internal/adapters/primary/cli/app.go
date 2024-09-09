package cli_adapter

import (
	"fmt"
	"log"
	"os"
	"strconv"

	user_service "github.com/tahadostifam/go-hexagonal-architecture/internal/core/services/user"
	"github.com/tahadostifam/go-hexagonal-architecture/utils"
	"github.com/urfave/cli/v2"
)

type ServicesApi struct {
	UserApi user_service.Api
}

func NewCLIAdapter(servicesApi *ServicesApi) {
	app := &cli.App{
		Name:  "Go Hexagonal CLI",
		Usage: "Simple implementation of CLI primary adapter with Go.",
		Action: func(*cli.Context) error {
			fmt.Println("Use --help to find out how this works.")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:      "create_account",
				Aliases:   []string{"ac"},
				Usage:     "Create account",
				Args:      true,
				ArgsUsage: "[name] [username] [age]",
				Action: func(ctx *cli.Context) error {
					name := ctx.Args().Get(0)
					username := ctx.Args().Get(1)
					age := ctx.Args().Get(2)
					ageInt, err := strconv.Atoi(age)
					utils.HandleError(err)

					user, err := servicesApi.UserApi.CreateAccount(ctx.Context, name, username, ageInt)
					utils.HandleError(err)

					fmt.Println("=== User created! ===")
					fmt.Println("ID", user.ID)
					fmt.Println("Name", user.Name)
					fmt.Println("Username", user.Username)
					fmt.Println("Age", user.Age)

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
