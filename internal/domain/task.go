package domain

import (
	"fmt"
	"strings"
	"time"
)

type TaskStatus string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

func (s TaskStatus) Validate() error {
	switch s {
	case StatusTodo, StatusInProgress, StatusDone:
		return nil
	default:
		return fmt.Errorf("invalid status: %s", s)
	}
}

type Task struct {
	ID          int
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(id int, description string) (*Task, error) {
	if strings.TrimSpace(description) == "" {
		return nil, fmt.Errorf("description cannot be empty")
	}
	now := time.Now()
	return &Task{
		ID:          id,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (t *Task) UpdateStatus(newStatus TaskStatus) error {
	if err := newStatus.Validate(); err != nil {
		return err
	}
	t.Status = newStatus
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) UpdateDescription(newDesc string) error {
	if strings.TrimSpace(newDesc) == "" {
		return fmt.Errorf("description cannot be empty")
	}
	t.Description = newDesc
	t.UpdatedAt = time.Now()
	return nil
}
