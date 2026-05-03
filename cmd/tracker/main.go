package main

import (
	"log"
	"task-tracker/internal/adapters/repository"
	"task-tracker/internal/cli"
	"task-tracker/internal/service"
)

func main() {
	// Инициализация
	repo := repository.NewJSONTaskRepository("tasks.json")
	taskService := service.NewTaskService(repo)

	app := cli.NewApp(taskService)

	if err := app.Execute(); err != nil {
		log.Fatal(err)
	}
}
