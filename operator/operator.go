package operator

import (
	"awesome-start/db"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func DeleteTask(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return nil
	}

	err = db.DbDeleteTask(id)

	if err != nil {
		return nil
	}

	return nil
}

func MarkTask(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return nil
	}

	task, err := db.DbGetTask(id)
	if err != nil {
		return nil
	}

	task.Status = "Completed"
	err = db.DbUpdateTask(id, task)
	if err != nil {
		return nil
	}

	return nil
}

func AddTask(c *fiber.Ctx) error {
	var body db.TaskBody
	err := c.BodyParser(&body)
	if err != nil {
		return nil
	}
	task := db.CreateTask(body.Title, body.Description)

	err = db.DbAddTask(task)

	if err != nil {
		return nil
	}

	return nil
}

type TaskResponse struct {
	Tasks []db.Task `json:"tasks"`
}

func ViewTask(c *fiber.Ctx) error {

	tasks, err := db.DbGetTasks()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	if len(tasks) == 0 {
		err := c.Status(200).JSON("There is no task to be viewed.")
		if err != nil {
			return err
		}
		return nil
	}

	var taskResponse TaskResponse

	for _, task := range tasks {
		taskResponse.Tasks = append(taskResponse.Tasks, task)
	}

	return c.JSON(taskResponse)
}
