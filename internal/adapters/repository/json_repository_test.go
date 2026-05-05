package repository

import (
	"os"
	"path/filepath"
	"task-tracker/internal/domain"
	"testing"
)

func TestJSONTaskRepository(t *testing.T) {
	// Создаем временную директорию для тестов
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "tasks_test.json")

	t.Run("FindAll on non-existing file returns empty list", func(t *testing.T) {
		// Файла еще нет
		repo := NewJSONTaskRepository(dbPath)
		tasks, err := repo.FindAll() // Убедись, что метод FindAll реализован
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(tasks) != 0 {
			t.Errorf("expected 0 tasks, got %d", len(tasks))
		}
	})

	t.Run("Save and FindByID", func(t *testing.T) {
		repo := NewJSONTaskRepository(dbPath)
		task := &domain.Task{Description: "Test Task"}

		err := repo.Save(task)
		if err != nil {
			t.Fatalf("failed to save task: %v", err)
		}

		if task.ID == 0 {
			t.Error("expected task ID to be assigned, got 0")
		}

		found, err := repo.FindByID(task.ID)
		if err != nil {
			t.Fatalf("failed to find task: %v", err)
		}
		if found.Description != task.Description {
			t.Errorf("expected title %s, got %s", task.Description, found.Description)
		}
	})

	t.Run("Delete existing and non-existing task", func(t *testing.T) {
		repo := NewJSONTaskRepository(dbPath)
		task := &domain.Task{Description: "To be deleted"}
		_ = repo.Save(task)

		// Удаляем существующую
		err := repo.Delete(task.ID)
		if err != nil {
			t.Errorf("failed to delete task: %v", err)
		}

		// Проверяем, что ее больше нет
		_, err = repo.FindByID(task.ID)
		if err == nil {
			t.Error("expected error finding deleted task, got nil")
		}

		// Удаляем несуществующую (обычно возвращает ошибку или nil, зависит от реализации)
		err = repo.Delete(9999)
		if err == nil {
			t.Log("Note: Delete non-existing task didn't return error (this is often acceptable)")
		}
	})

	t.Run("Broken JSON returns error", func(t *testing.T) {
		repo := NewJSONTaskRepository(dbPath)

		// Записываем мусор в файл
		err := os.WriteFile(dbPath, []byte("{ broken json ..."), 0644)
		if err != nil {
			t.Fatal(err)
		}

		_, err = repo.FindByID(1)
		if err == nil {
			t.Error("expected error for broken JSON, got nil")
		}
	})
}
