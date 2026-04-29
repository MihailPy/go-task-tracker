package service

import (
	"fmt"
	"task-tracker/internal/domain"
	"task-tracker/internal/ports"
)

type TaskService struct {
	repo ports.TaskRepository
}

func NewTaskService(repo ports.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) AddTask(description string) (*domain.Task, error) {
	task, err := domain.NewTask(0, description)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(task); err != nil {
		return nil, fmt.Errorf("failed to save task: %w", err)
	}

	return task, nil
}

func (s *TaskService) UpdateTaskStatus(id int, newStatus domain.TaskStatus) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("action failed for task %d: %w", id, err)
	}

	if err := task.UpdateStatus(newStatus); err != nil {
		return err
	}

	return s.repo.Save(task)
}

func (s *TaskService) UpdateTaskDescription(id int, newDescription string) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("action failed for task %d: %w", id, err)
	}

	if err := task.UpdateDescription(newDescription); err != nil {
		return err
	}

	return s.repo.Save(task)
}

func (s *TaskService) DeleteTask(id int) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("action failed for task %d: %w", id, err)
	}

	return s.repo.Delete(id)
}

func (s *TaskService) ListAllTasks() ([]*domain.Task, error) {
	return s.repo.FindAll()
}

func (s *TaskService) ListTasksByStatus(status domain.TaskStatus) ([]*domain.Task, error) {
	if err := status.Validate(); err != nil {
		return nil, err
	}
	return s.repo.FindByStatus(status)
}
