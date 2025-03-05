package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func displayToDoHelp() {
	fmt.Println("\nCLI based TO-DO App Task:")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Task")
	fmt.Println("3. Mark a Task as Done")
	fmt.Println("4. Delete a Task")
	fmt.Println("5. Exit")
	fmt.Print("\nEnter your choice: ")
}

func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, inputError := reader.ReadString('\n')
	return strings.TrimSpace(input), inputError
}

func getInputNumber() (int, bool) {
	taskNumber, _ := getInput()
	taskIndex, taskIndexError := strconv.Atoi(taskNumber)

	if taskIndexError != nil {
		fmt.Println("Invalid input number, please choice a number", taskIndexError)
		return taskIndex, false
	}

	return taskIndex, true
}

func validateInputTaskNumber(taskNumber int, tasks []string) int {
	if taskNumber <= 0 || taskNumber > len(tasks) {
		fmt.Println("Invalid task number", taskNumber)
		return -1
	}

	return taskNumber
}

func addTask(task string, tasks *[]string) {
	*tasks = append(*tasks, task)
	fmt.Println("Task added successfully")
}

func listTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show")
		return
	}

	fmt.Println("\nTasks:")
	for index, task := range tasks {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func markTaskDone(taskNumber int, tasks *[]string) bool {
	taskNumber = validateInputTaskNumber(taskNumber, *tasks)
	if taskNumber == -1 {
		return false
	}
	(*tasks)[taskNumber-1] = "[+]" + (*tasks)[taskNumber-1]
	fmt.Println("Task marked as done successfully")
	return true
}

func deleteTask(taskNumber int, tasks *[]string) bool {
	taskNumber = validateInputTaskNumber(taskNumber, *tasks)
	if taskNumber == -1 {
		return false
	}
	*tasks = append((*tasks)[:taskNumber-1], (*tasks)[taskNumber:]...)
	fmt.Println("Task deleted successfully")
	return true
}

func main() {
	var tasks []string

	for {
		displayToDoHelp()
		choice, isValid := getInputNumber()
		if !isValid {
			return
		}

		switch choice {
		case 1:
			fmt.Print("\nEnter a task: ")
			task, _ := getInput()
			addTask(task, &tasks)
		case 2:
			listTasks(tasks)
		case 3:
			if len(tasks) == 0 {
				fmt.Println("No tasks is available to mark as done")
				break
			}

			fmt.Print("Enter the task number to mark as done: ")
			taskNumber, isValid := getInputNumber()
			if !isValid {
				break
			}
			markTaskDone(taskNumber, &tasks)
		case 4:
			if len(tasks) == 0 {
				fmt.Println("No tasks is available to delete")
				break
			}

			fmt.Print("Enter the task number to delete: ")
			taskNumber, isValid := getInputNumber()
			if !isValid {
				break
			}
			deleteTask(taskNumber, &tasks)
		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
			displayToDoHelp()
		}
	}
}
