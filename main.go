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

func main() {
	fileName := "tasks.json"

	// Открываем файл:
	// os.O_RDWR — чтение/запись
	// os.O_CREATE — создать, если нет
	// os.O_APPEND — (опционально) дописывать в конец
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	now := time.Now()
	// Данные для записи
	data := Task{Id: 1, Description: "My first task", Status: "todo", CreatedAt: now, UpdatedAt: now}

	// Запись в формате JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // для красивого форматирования
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Ошибка при записи JSON:", err)
		return
	}

	fmt.Println("Файл успешно открыт/создан и обновлен.")
	fmt.Println("Трекер задач запущен!")
}
