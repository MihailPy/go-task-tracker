package cli

import (
	"errors"
	"fmt"
	"task-tracker/internal/domain"

	"github.com/spf13/cobra"
)

func (a *App) TaskAddCmd() *cobra.Command {
	var desc string
	cmd := &cobra.Command{
		Use:   "add [task description]",
		Short: "Добавление задачи. Описание задачи (обязательно)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			desc = args[0]
			task, err := a.taskService.AddTask(desc)
			if err != nil {
				if errors.Is(err, domain.ErrEmptyDescription) {
					return errors.New("ошибка: описание задачи не может быть пустым")
				}
				return fmt.Errorf("критическая ошибка при добавлении: %w", err)
			}

			fmt.Printf("✅ Добавлена задача #%d: %s\n", task.ID, task.Description)
			return nil
		},
	}
	return cmd
}
