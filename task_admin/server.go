package task_admin

import (
	"github.com/gofiber/fiber/v2"
)

type TaskBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Routers(app *fiber.App) {
	app.Get("/tasks/view", ViewTask)
	app.Post("/tasks/add", ApiAddTask)
	app.Put("/tasks/completeTask/:id", ApiMarkTask)
	app.Delete("/tasks/deleteTask/:id", ApiDeleteTask)
}

func Init(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the API")
	})

	Routers(app)

	err := app.Listen(":3162")
	if err != nil {
		panic("uh!")
		return
	}

}
