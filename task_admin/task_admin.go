package task_admin

import (
	"bufio"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var reader = bufio.NewReader(os.Stdin)

func DeleteTheTask() {
	for {
		fmt.Print("\nEnter the ID of the task to delete: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			fmt.Println("Invalid input. Please enter a valid task ID.")
			continue
		}

		taskID, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid task ID.")
			continue
		}

		err = DbDeleteTask(taskID)
		if err != nil {
			fmt.Println("Failed to delete the task:", err)
			break
		}

		fmt.Println("\nTask deleted successfully!\n")
		break
	}
}

func MarkTheTask() {
	for {
		fmt.Print("\nEnter the ID of the task to mark as completed: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			fmt.Println("Invalid input. Please enter a valid task ID.")
			continue
		}

		taskID, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid task ID.")
			continue
		}

		task, err := DbGetTask(taskID)
		if err != nil {
			fmt.Println("Failed to retrieve the task:", err)
			break
		}

		task.Status = "Completed"
		err = DbUpdateTask(taskID, task)
		if err != nil {
			fmt.Println("Failed to update the task:", err)
			break
		}

		fmt.Println("\nTask marked as completed successfully!\n")
		break
	}
}

func ApiDeleteTask(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.Status(400).SendString("Invalid id. ")
	}

	err = DbDeleteTask(id)

	if err != nil {
		return c.Status(400).SendString("Task can't be deleted")
	}

	return c.Status(200).SendString("Task is deleted.")
}

func ApiMarkTask(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.Status(400).SendString("Invalid id. ")
	}

	task, err := DbGetTask(id)
	if err != nil {
		fmt.Println("Failed to retrieve the task:", err)
	}

	task.Status = "Completed"
	err = DbUpdateTask(id, task)
	if err != nil {
		return c.Status(400).SendString("Could not mark the task")
	}

	return c.Status(200).SendString("Task is marked.")
}

func ApiAddTask(c *fiber.Ctx) error {
	var body TaskBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).SendString("Error while parsing the body")
	}

	task := DbCreateTask(body.Title, body.Description)

	err = DbAddTask(task)

	if err != nil {
		return c.Status(400).SendString("Could not add the task.")
	}

	return c.Status(200).JSON(task)
}

func AddTask() {
	for {
		var title string
		var desc string
		fmt.Print("\nEnter the title of the task: ")
		title, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Enter title properly.")
			continue
		}
		fmt.Print("Enter the description of the task: ")
		desc, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Enter desc properly.")
			continue
		}

		task := DbCreateTask(title, desc)

		err = DbAddTask(task)
		if err != nil {
			fmt.Println("Failed to add the task:", err)
			break
		}

		fmt.Println("\nTask added successfully!\n")
		break
	}
}

func ViewTask(*fiber.Ctx) error {
	tasks, err := DbGetTasks()
	if err != nil {
		fmt.Println("Failed to retrieve tasks:", err)
		return nil
	}

	if len(tasks) == 0 {
		fmt.Println("\nNo tasks found!\n")
		return nil
	}

	fmt.Println("\nTasks:\n")

	for _, task := range tasks {
		fmt.Printf("Task ID: %d\n", task.TaskID)
		fmt.Printf("Title: %s\n", task.Title)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Status: %s\n\n", task.Status)
	}
	return nil
}
