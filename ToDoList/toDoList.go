package main

import (
	"bufio"
	"fmt"
	"os"
)

var taskList []Task

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		func() {
			defer handleRecover()
			showMenu(reader)
		}()
	}
}

func showMenu(reader *bufio.Reader) {
	var choice int
	fmt.Println("\n\n\nChoose operation \n 1)Add task \n 2)View tasks \n 3)Mark selected task as done \n 4)Delete task \n 5)Unfinished Tasks \n 6)Exit")
	fmt.Scanln(&choice)

	if choice < 1 || choice > 6 {
		panic("Invalid choice! Please choose a number between 1 and 6.")
	}

	handleChoice(choice, reader)
}

func handleChoice(choice int, reader *bufio.Reader) {
	fmt.Println("\n")
	switch choice {
	case 1:
		fmt.Println("Insert task text:")
		text, _ := reader.ReadString('\n')
		task := Task{Text: text}
		taskList = append(taskList, task)

	case 2:
		for index, task := range taskList {
			fmt.Println(index, ")", task.Text)
		}
		exitWhenAsked()

	case 3:
		var taskId int
		fmt.Println("Choose task to mark as done by id")
		fmt.Scanln(&taskId)

		if taskId < 0 || taskId >= len(taskList) {
			panic("Invalid task ID! Please enter a valid task ID.")
		}

		taskList[taskId].Done = true

	case 4:
		var index int
		fmt.Println("Choose task to delete by id")
		fmt.Scanln(&index)

		if index < 0 || index >= len(taskList) {
			panic("Invalid task ID! Please enter a valid task ID.")
		}

		taskList = append(taskList[:index], taskList[index+1:]...)

	case 5:
		for index, task := range taskList {
			if !task.Done {
				fmt.Println(index, ")", task.Text)
			}
		}
		exitWhenAsked()

	case 6:
		fmt.Println("Thanks")
		os.Exit(0)
	}
}

func exitWhenAsked() {
	var leave int
	for {
		fmt.Scanln(&leave)
		if leave == 1 {
			break
		}
	}
}

func handleRecover() {
	if r := recover(); r != nil {
		fmt.Printf("Error: %v\n", r)
		fmt.Println("Please try again with valid input.")
	}
}

type Task struct {
	Text string
	Done bool
}
