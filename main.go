package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func SaveTasks(fileName string, tasks []Task) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}
func LoadTasks(fileName string) ([]Task, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []Task{}, nil
	}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
func main() {
	fileName := "tasks.json"
	now := time.Now()

	tasks, err := LoadTasks(fileName)
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}

	// Данные для записи
	newTask := Task{Id: len(tasks) + 1, Description: "My first task", Status: "todo", CreatedAt: now, UpdatedAt: now}

	tasks = append(tasks, newTask)

	if err := SaveTasks(fileName, tasks); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	fmt.Printf("Успешно открыт/создан файл.\nУспешно создана задача: %d\nУспешно сохранена задача в файле.\n", len(tasks))
}
