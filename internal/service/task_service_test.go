package service

import (
	"errors"
	"testing"

	"task-tracker/internal/domain"
)

// MockRepository — простая имитация хранилища
type MockRepository struct {
	tasks       map[int]*domain.Task
	nextID      int
	shouldFail  bool
	notFoundErr error
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		tasks:       make(map[int]*domain.Task),
		nextID:      1,
		notFoundErr: errors.New("not found"),
	}
}

func (m *MockRepository) Save(t *domain.Task) error {
	if m.shouldFail {
		return errors.New("db error")
	}
	if t.ID == 0 {
		t.ID = m.nextID
		m.nextID++
	}
	m.tasks[t.ID] = t
	return nil
}

func (m *MockRepository) FindByID(id int) (*domain.Task, error) {
	t, ok := m.tasks[id]
	if !ok {
		return nil, m.notFoundErr
	}
	return t, nil
}

func (m *MockRepository) Delete(id int) error {
	delete(m.tasks, id)
	return nil
}

func (m *MockRepository) FindAll() ([]*domain.Task, error) {
	if m.shouldFail {
		return nil, errors.New("db error")
	}
	var list []*domain.Task
	for _, t := range m.tasks {
		list = append(list, t)
	}
	return list, nil
}

func (m *MockRepository) FindByStatus(status domain.TaskStatus) ([]*domain.Task, error) {
	return nil, nil // Для данного ТЗ не требуется
}

// ТЕСТЫ

func TestTaskService(t *testing.T) {
	repo := NewMockRepository()
	svc := NewTaskService(repo)

	t.Run("AddTask creates ID and saves", func(t *testing.T) {
		task, err := svc.AddTask("Test task")
		if err != nil || task.ID == 0 {
			t.Errorf("Expected task with ID, got error: %v", err)
		}
	})

	t.Run("FindAll/Save error handling", func(t *testing.T) {
		repo.shouldFail = true
		_, err := svc.ListAllTasks()
		if err == nil {
			t.Error("Expected error on ListAllTasks, got nil")
		}

		_, err = svc.AddTask("Fail task")
		if err == nil {
			t.Error("Expected error on Save, got nil")
		}
		repo.shouldFail = false
	})

	t.Run("Update actions - Not Found", func(t *testing.T) {
		err := svc.UpdateTaskStatus(999, domain.StatusDone)
		if err == nil {
			t.Error("Expected error for non-existent task status update")
		}

		err = svc.UpdateTaskDescription(999, "New desc")
		if err == nil {
			t.Error("Expected error for non-existent task description update")
		}
	})

	t.Run("DeleteTask - Success and Not Found", func(t *testing.T) {
		// Создаем задачу для удаления
		task, _ := svc.AddTask("To be deleted")

		// Успешное удаление
		err := svc.DeleteTask(task.ID)
		if err != nil {
			t.Errorf("Expected successful delete, got: %v", err)
		}

		// Удаление несуществующей
		err = svc.DeleteTask(999)
		if err == nil {
			t.Error("Expected error when deleting non-existent task")
		}
	})
}
