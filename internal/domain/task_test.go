package domain

import (
	"errors"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		description string
		wantErr     error
	}{
		{
			name:        "Valid description",
			id:          1,
			description: "Buy milk",
			wantErr:     nil,
		},
		{
			name:        "Empty description",
			id:          2,
			description: "  ",
			wantErr:     ErrEmptyDescription,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := NewTask(tt.id, tt.description)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if task.ID != tt.id {
					t.Errorf("expected ID %d, got %d", tt.id, task.ID)
				}
				if task.Status != StatusTodo {
					t.Errorf("expected Status %s, got %s", StatusTodo, task.Status)
				}
				if task.CreatedAt.IsZero() {
					t.Error("expected CreatedAt to be set, got zero time")
				}
			}
		})
	}
}

func TestUpdateDescription(t *testing.T) {
	tests := []struct {
		name    string
		newDesc string
		wantErr error
	}{
		{
			name:    "Valid update",
			newDesc: "Update task description",
			wantErr: nil,
		},
		{
			name:    "Empty description",
			newDesc: "   ",
			wantErr: ErrEmptyDescription,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, _ := NewTask(1, "Original description")
			initialTime := task.UpdatedAt

			time.Sleep(1 * time.Millisecond)

			err := task.UpdateDescription(tt.newDesc)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UpdateDescription() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				if task.Description != tt.newDesc {
					t.Errorf("expected description %s, got %s", tt.newDesc, task.Description)
				}
				if !task.UpdatedAt.After(initialTime) {
					t.Error("expected UpdatedAt to be updated")
				}
			}
		})
	}
}

func TestUpdateStatus(t *testing.T) {
	tests := []struct {
		name      string
		newStatus TaskStatus
		wantErr   error
	}{
		{
			name:      "Valid update",
			newStatus: StatusInProgress,
			wantErr:   nil,
		},
		{
			name:      "Invalid Status",
			newStatus: TaskStatus("ololo"),
			wantErr:   ErrInvalidStatus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, _ := NewTask(1, "Test task")
			initialTime := task.UpdatedAt

			time.Sleep(1 * time.Millisecond)

			err := task.UpdateStatus(tt.newStatus)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				if task.Status != tt.newStatus {
					t.Errorf("expected status %s, got %s", tt.newStatus, task.Status)
				}
				if !task.UpdatedAt.After(initialTime) {
					t.Error("expected UpdatedAt to be updated")
				}
			}
		})
	}
}
