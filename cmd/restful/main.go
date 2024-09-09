package main

import (
	"log"

	restful_adapter "github.com/samverrall/hex-structure/internal/adapters/primary/restful"
	sqlite_adapter "github.com/samverrall/hex-structure/internal/adapters/secondary/sqlite"
	user_service "github.com/samverrall/hex-structure/internal/core/services/user"

	"gorm.io/driver/sqlite"
)

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Init database dialector
	dialector := sqlite.Open("./database/sqlite3.db")

	// Init secondary adapters
	userRepo, err := sqlite_adapter.NewUserRepository(dialector)
	handleError(err)

	// Init business logic
	userService := user_service.NewService(userRepo)

	// Init primary adapters
	app := restful_adapter.NewApp(restful_adapter.ServiceApi{
		UserApi: userService,
	})
	err = app.Run()
	handleError(err)
}
