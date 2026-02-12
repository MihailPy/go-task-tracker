package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type TaskStatus string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func FindTaskById(tasks []Task, idTask int) int {
	for i := range tasks {
		if tasks[i].Id == idTask {
			return i
		}
	}
	return -1
}
func ListTasks(tasks []Task) {
	for _, t := range tasks {
		fmt.Printf("№ %d | Задача: %s | Статус: %s | Время создания: %v | Последние время редоктирования: %v\n", t.Id, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
	}
}
func AddTask(description string, tasks []Task) ([]Task, error) {
	if strings.TrimSpace(description) == "" {
		return tasks, fmt.Errorf("description cannot be empty")
	}
	maxId := 0
	for _, t := range tasks {
		if t.Id > maxId {
			maxId = t.Id
		}
	}
	now := time.Now()
	newTask := Task{Id: maxId + 1, Description: description, Status: StatusTodo, CreatedAt: now, UpdatedAt: now}
	return append(tasks, newTask), nil
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

	tasks, err := LoadTasks(fileName)
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}

	description := "My description task"
	tasks, err = AddTask(description, tasks)

	ListTasks(tasks)

	if err := SaveTasks(fileName, tasks); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	fmt.Printf("Успешно открыт/создан файл.\nУспешно создана задача: %d\nУспешно сохранена задача в файле.\n", len(tasks))
}
