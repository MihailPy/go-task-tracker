package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"task-tracker/internal/domain"
)

type JSONTaskRepository struct {
	filePath string
	mu       sync.RWMutex
}

func NewJSONTaskRepository(filePath string) *JSONTaskRepository {
	return &JSONTaskRepository{filePath: filePath}
}

func (r *JSONTaskRepository) Save(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.load()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			found = true
			break
		}
	}
	if !found {
		tasks = append(tasks, task)
	}

	return r.save(tasks)
}

func (r *JSONTaskRepository) FindByID(id int) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks, err := r.load()
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, nil
}

func (r *JSONTaskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.load()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return r.save(tasks)
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func (r *JSONTaskRepository) FindAll() ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.load()
}

func (r *JSONTaskRepository) FindByStatus(status domain.TaskStatus) ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks, err := r.load()
	if err != nil {
		return nil, err
	}

	var result []*domain.Task
	for _, t := range tasks {
		if t.Status == status {
			result = append(result, t)
		}
	}
	return result, nil
}

func (r *JSONTaskRepository) load() ([]*domain.Task, error) {
	file, err := os.Open(r.filePath)
	if os.IsNotExist(err) {
		return []*domain.Task{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []*domain.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *JSONTaskRepository) save(tasks []*domain.Task) error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}
