package main

import (
	Auth "awesome-start/auth"
	"awesome-start/db"
	"awesome-start/operator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

func Routers(app *fiber.App) {
	app.Post("/login", Auth.Login)
	app.Post("/register", Auth.Register)

	app.Get("/tasks/view", RequiredAuth(), operator.ViewTask)
	app.Post("/tasks/add", RequiredAuth(), operator.AddTask)
	app.Put("/tasks/completeTask/:id", RequiredAuth(), operator.MarkTask)
	app.Delete("/tasks/deleteTask/:id", RequiredAuth(), operator.DeleteTask)
}

func RequiredAuth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(Auth.JwtSecret),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errorCode": err.Error(),
				"error":     "Unauthorized",
			})
			return nil
		},
	})
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
