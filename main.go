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

func FilterTasksByStatus(tasks []Task, s TaskStatus) {
	for i := range tasks {
		t := tasks[i]
		if t.Status == s {
			fmt.Printf("№ %d | Задача: %s | Статус: %s | Время создания: %v | Последние время редоктирования: %v\n", t.Id, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
		}
	}
}
func UpdateTaskStatus(t []Task, id int, s TaskStatus) ([]Task, error) {
	if len(t) == 0 {
		return nil, fmt.Errorf("list index out of range")
	}
	idx := FindTaskById(t, id)
	if idx != -1 {
		t[idx].Status = s
		t[idx].UpdatedAt = time.Now()
		return t, nil
	} else {
		return nil, fmt.Errorf("Element with ID '%d' not found in the list.", id)
	}
}
func UpdateTask(tasks []Task, idTask int, description string) ([]Task, error) {
	if len(tasks) == 0 {
		return nil, fmt.Errorf("list index out of range")
	}
	idx := FindTaskById(tasks, idTask)
	if idx != -1 {
		tasks[idx].Description = description
		tasks[idx].UpdatedAt = time.Now()
		return tasks, nil
	} else {
		return nil, fmt.Errorf("Element with ID '%d' not found in the list.", idTask)
	}

}
func DeleteTask(tasks []Task, idTask int) ([]Task, error) {
	if len(tasks) == 0 {
		return nil, fmt.Errorf("list index out of range")
	}
	idx := FindTaskById(tasks, idTask)
	if idx != -1 {
		return append(tasks[:idx], tasks[idx+1:]...), nil
	} else {
		return nil, fmt.Errorf("Element with ID '%d' not found in the list.", idTask)
	}
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
	result, err := AddTask(description, tasks)
	if err != nil {
		fmt.Println("Ошибка дбавление задачи: ", err)
	} else {
		tasks = result
	}
	// result, err = DelTask(tasks, 1)
	// if err != nil {
	// 	fmt.Println("Ошибка удаления задачи: ", err)
	// } else {
	// 	tasks = result
	// }
	// result, err = UpdateTask(tasks, 8, "New description task")
	// if err != nil {
	// 	fmt.Println("Ошибка удаления: ", err)
	// } else {
	// 	tasks = result
	// }

	// ListTasks(tasks)

	FilterTasksByStatus(tasks, StatusTodo)

	result, err = UpdateTaskStatus(tasks, 8, StatusInProgress)
	if err != nil {
		fmt.Println("Ошибка изменения статуса задачи: ", err)
	} else {
		tasks = result
	}
	ListTasks(tasks)

	if err := SaveTasks(fileName, tasks); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	fmt.Printf("Успешно открыт/создан файл.\nУспешно создана задача: %d\nУспешно сохранена задача в файле.\n", len(tasks))
}
