package handlers

import (
	"fmt"
	"my-go-project/models"
)

var tasks []models.Task

func AddTask(title string) {
	id := len(tasks) + 1
	task := models.NewTask(id, title)
	tasks = append(tasks, task)
	fmt.Println("Added task:", task)
}

func ListTasks() {
	for _, t := range tasks {
		fmt.Println(t.ID, "-", t.Title, "Done:", t.Done)
	}
}
