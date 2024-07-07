package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Taskname  string
	completed bool
}

var tasks []Task

func addTask(task string) {
	newTask := Task{Taskname: task, completed: false}
	tasks = append(tasks, newTask)
	fmt.Println("Task Added")

}

func listTasks() {
	for i, task := range tasks {
		status := "n"
		if task.completed {
			status = "d"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Taskname, status)
	}
}

func markCompleted(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].completed = true
		fmt.Println("Task Marked as completed")

	} else {
		fmt.Println("Invalid Index")
	}
}

func editTask(index int, newString string) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].Taskname = newString
		fmt.Println("Task Edited Successfully")

	} else {
		fmt.Println("Invalid Index")
	}
}

func deleteTask(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks = append(tasks[:index-1], tasks[index:]...)
		fmt.Println("Task deleted")

	} else {
		fmt.Println("Invalid Index")
	}
}

func main() {

	var indexInput int
	var taskInput, newTaskInput string

	fmt.Println("Options")
	fmt.Println("1. Add Tasks")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Mark as completed")
	fmt.Println("4. Edit Task")
	fmt.Println("5. Delete Task")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter Choice (1, 2, 3, 4, 5, 6): ")
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalide Choice")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("ENter task: ")
			scanner.Scan()
			taskInput = scanner.Text()
			addTask(taskInput)

		case 2:
			listTasks()

		case 3:
			fmt.Println("Enter Index: ")
			scanner.Scan()

			indexInput, _ = strconv.Atoi(scanner.Text())
			markCompleted(indexInput)

		case 4:
			fmt.Println("Enter Index")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			fmt.Println("Enter task: ")

			scanner.Scan()
			newTaskInput = scanner.Text()

			editTask(indexInput, newTaskInput)

		case 5:
			fmt.Println("Enter Index")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			deleteTask(indexInput)

		case 6:
			os.Exit(0)

		default:
			fmt.Println("Invalid Choice")

		}

	}

}
