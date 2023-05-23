package task_admin

import (
	"awesome-start/task_admin/db"
	_ "awesome-start/task_admin/db"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

		err = db.DeleteTask(taskID)
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

		task, err := db.GetTask(taskID)
		if err != nil {
			fmt.Println("Failed to retrieve the task:", err)
			break
		}

		task.Status = "Completed"
		err = db.UpdateTask(taskID, task)
		if err != nil {
			fmt.Println("Failed to update the task:", err)
			break
		}

		fmt.Println("\nTask marked as completed successfully!\n")
		break
	}
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

		task := db.Task{
			Title:       strings.TrimSpace(title),
			Description: strings.TrimSpace(desc),
			Status:      "Pending",
		}

		err = db.AddTask(&task)
		if err != nil {
			fmt.Println("Failed to add the task:", err)
			break
		}

		fmt.Println("\nTask added successfully!\n")
		break
	}
}

func ViewTask() {
	tasks, err := db.GetTasks()
	if err != nil {
		fmt.Println("Failed to retrieve tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("\nNo tasks found!\n")
		return
	}

	fmt.Println("\nTasks:\n")

	for _, task := range tasks {
		fmt.Printf("Task ID: %d\n", task.TaskID)
		fmt.Printf("Title: %s\n", task.Title)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Status: %s\n\n", task.Status)
	}
}
