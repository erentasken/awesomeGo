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

	// Group routes that require authentication
	authenticated := app.Group("", RequiredAuth())

	authenticated.Get("/tasks/view", operator.ViewTask)
	authenticated.Post("/tasks/add", operator.AddTask)
	authenticated.Put("/tasks/completeTask/:id", operator.MarkTask)
	authenticated.Delete("/tasks/deleteTask/:id", operator.DeleteTask)
}

func RequiredAuth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SuccessHandler: func(c *fiber.Ctx) error {
			err := c.Status(200).JSON(fiber.Map{
				"Authorization": "Success",
			})
			if err != nil {
				return err
			}
			return c.Next()
		},
		SigningKey: []byte(Auth.JwtSecret),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			err = ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errorCode": err.Error(),
				"error":     "Unauthorized",
			})
			if err != nil {
				return err
			}
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
