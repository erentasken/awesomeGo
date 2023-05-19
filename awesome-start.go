package main

import (
	"awesome-start/task_admin"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Task Manager!\n")
	var choice int = -1
	menu := [5]string{"Add a Task", "View Tasks", "Mark a Task as Completed", "Delete a Task", "Exit"}

	for choice != len(menu) {
		fmt.Println("Menu:")
		for i, option := range menu {
			fmt.Printf("%d. %s\n", i+1, option)
		}

		fmt.Print("\nEnter your choice: ")
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid choice.\n")
			continue
		}

		switch choice {
		case 1:
			task_admin.AddTask()
			break
		case 2:
			task_admin.ViewTask()
			break
		case 3:
			task_admin.MarkTheTask()
			break
		case 4:
			task_admin.DeleteTheTask()
			break
		case 5:
			fmt.Println("\nI take my horse and leaving the town with the resolute farewells embracing the boundless horizons cowboy, adios fellas")
			os.Exit(0)
		default:
			fmt.Println("\nEnter proper choice\n")
			continue
		}
	}

	task_admin.AddTask()

}
