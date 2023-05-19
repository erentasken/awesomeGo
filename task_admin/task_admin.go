package task_admin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var reader = bufio.NewReader(os.Stdin)
var taskMap = make(map[int]Task)
var taskId int = 1

type Task struct {
	TITLE       string
	DESCRIPTION string
	STATUS      string
	TASKID      int
}

func DeleteTheTask() {
	for {
		fmt.Print("\nEnter the ID of the task to delete: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			fmt.Println("Invalid input. Please enter a valid task ID.")
			continue
		}

		markId, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid task ID.")
			continue
		}

		if _, exist := taskMap[markId]; exist {
			delete(taskMap, markId)
			fmt.Println("\nTask deleted successfully!\n")
			taskId--
			break
		} else {
			fmt.Println("Task ID not found.\n")
			break
		}
	}
}

func MarkTheTask() {
	for {
		var markId int
		fmt.Print("\nEnter the ID of the task to mark as completed: ")
		scanner.Scan()

		markId, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid task ID.")
			continue
		}

		if task, exist := taskMap[markId]; exist {
			task.STATUS = "Completed"
			taskMap[markId] = task
			fmt.Println("\nTask marked as completed successfully!\n")
			break
		} else {
			fmt.Println("Task ID not found.\n")
			break
		}
	}

}

func AddTask() {
	for {
		var title string
		var desc string
		fmt.Print("\nEnter the title of the task: ")
		title, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Enter title properly. ")
			continue
		}
		fmt.Print("Enter the description of the task: ")
		desc, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Enter desc properly. ")
			continue
		}

		task := Task{
			title,
			desc,
			"Pending",
			taskId,
		}

		taskMap[taskId] = task

		taskId++
		fmt.Println("\nTask added successfully!\n")
		break
	}

}

func ViewTask() {
	if taskId == 1 {
		fmt.Println("\nNo tasks found!\n")
		return
	}
	fmt.Println("\nTasks : \n")

	for _, task := range taskMap {
		fmt.Printf("Task ID: %d\n", task.TASKID)
		fmt.Printf("Title: %s", task.TITLE)
		fmt.Printf("Description: %s", task.DESCRIPTION)
		fmt.Printf("Status: %s\n\n", task.STATUS)
	}

}
