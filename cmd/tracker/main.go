package main

import (
	"fmt"
	"log"
	"task-tracker/internal/adapters/repository"
	"task-tracker/internal/domain"
	"task-tracker/internal/service"
)

func main() {
	// Инициализация
	repo := repository.NewJSONTaskRepository("tasks.json")
	taskService := service.NewTaskService(repo)

	// === ПРИМЕРЫ ИСПОЛЬЗОВАНИЯ ===

	// 1. Добавить задачу
	task, err := taskService.AddTask("Купить хлеб")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("✅ Добавлена задача #%d: %s\n", task.ID, task.Description)

	// 2. Добавить ещё одну
	taskService.AddTask("Сделать домашку")

	// 3. Вывести все задачи
	allTasks, _ := taskService.ListAllTasks()
	fmt.Println("\n📋 Все задачи:")
	for _, t := range allTasks {
		fmt.Printf("  #%d: %s [%s]\n", t.ID, t.Description, t.Status)
	}

	// 4. Обновить статус
	err = taskService.UpdateTaskStatus(1, domain.StatusInProgress)
	if err != nil {
		fmt.Printf("❌ Ошибка: %v\n", err)
	} else {
		fmt.Printf("\n🔄 Статус задачи #1 обновлён на %s\n", domain.StatusInProgress)
	}

	// 5. Вывести задачи со статусом "todo"
	todoTasks, _ := taskService.ListTasksByStatus(domain.StatusTodo)
	fmt.Println("\n📌 Задачи в статусе 'todo':")
	for _, t := range todoTasks {
		fmt.Printf("  #%d: %s\n", t.ID, t.Description)
	}

	// 6. Обновить описание
	err = taskService.UpdateTaskDescription(2, "Сделать домашку по математике")
	if err != nil {
		fmt.Printf("❌ Ошибка: %v\n", err)
	} else {
		fmt.Printf("\n📝 Описание задачи #2 обновлено\n")
	}

	// 7. Удалить задачу
	err = taskService.DeleteTask(2)
	if err != nil {
		fmt.Printf("❌ Ошибка: %v\n", err)
	} else {
		fmt.Printf("\n🗑️ Задача #2 удалена\n")
	}

	// 8. Финальный список
	finalTasks, _ := taskService.ListAllTasks()
	fmt.Println("\n🏁 Финальный список задач:")
	for _, t := range finalTasks {
		fmt.Printf("  #%d: %s [%s] (создана: %s)\n",
			t.ID, t.Description, t.Status,
			t.CreatedAt.Format("02.01.2006 15:04"))
	}
}
