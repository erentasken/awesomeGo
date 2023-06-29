package main

import (
	Auth "awesome-start/auth"
	"awesome-start/db"
	"awesome-start/operator"
	"awesome-start/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Routers(app *fiber.App) {
	app.Post("/login", user.Login)
	app.Post("/register", user.Register)

	// Group routes that require authentication
	authenticated := app.Group("", Auth.RequiredAuth())

	authenticated.Get("/tasks/view", operator.ViewTask)
	authenticated.Post("/tasks/add", operator.AddTask)
	authenticated.Put("/tasks/completeTask/:id", operator.MarkTask)
	authenticated.Delete("/tasks/deleteTask/:id", operator.DeleteTask)
}

func main() {
	app := fiber.New()
	db.InitDB()

	// Register the middleware using the Use method
	app.Use(logger.New())

	Routers(app)

	err := app.Listen(":3000")
	if err != nil {
		panic("Failed to start the server.")
	}
}
