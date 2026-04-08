package ports

import "task-tracker/internal/domain"

type TaskRepository interface {
	Save(task *domain.Task) error          // Добавить или обновить
	FindByID(id int) (*domain.Task, error) // Вернуть nil, если не найден
	Delete(id int) error
	FindAll() ([]*domain.Task, error)
	FindByStatus(status domain.TaskStatus) ([]*domain.Task, error)
}
