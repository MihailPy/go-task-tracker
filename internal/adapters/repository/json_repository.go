package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
		return fmt.Errorf("repository cannot load tasks for saving: %w", err)
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

	err = r.save(tasks)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

func (r *JSONTaskRepository) FindByID(id int) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks, err := r.load()
	if err != nil {
		return nil, fmt.Errorf("failed to load tasks: %w", err)
	}

	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, fmt.Errorf("task %d: %w", id, domain.ErrTaskNotFound)
}

func (r *JSONTaskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.load()
	if err != nil {
		return fmt.Errorf("failed to load tasks: %w", err)
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			err := r.save(tasks)
			if err != nil {
				return fmt.Errorf("failed to save tasks after deletion: %w", err)
			}

			return nil
		}
	}

	return fmt.Errorf("task %d: %w", id, domain.ErrTaskNotFound)
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
		return nil, fmt.Errorf("failed to load tasks: %w", err)
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
		return nil, fmt.Errorf("failed to open storage file: %w", err)
	}
	defer file.Close()

	var tasks []*domain.Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err == io.EOF {
		return []*domain.Task{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to decode tasks json: %w", err)
	}
	return tasks, nil
}

func (r *JSONTaskRepository) save(tasks []*domain.Task) error {
	dir := filepath.Dir(r.filePath)
	tempFile, err := os.CreateTemp(dir, "tmp-*.json.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp file in %s: %w", dir, err)
	}
	tempPath := tempFile.Name()

	defer os.Remove(tempPath)

	defer tempFile.Close()

	encoder := json.NewEncoder(tempFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("failed to encode tasks to json: %w", err)
	}

	if err := tempFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync temp file: %w", err)
	}

	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	if err := os.Rename(tempPath, r.filePath); err != nil {
		return fmt.Errorf("failed to rename temp file to %s: %w", r.filePath, err)
	}
	return nil
}
