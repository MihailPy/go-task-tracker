package cli

import (
	"errors"
	"fmt"
	"task-tracker/internal/domain"

	"github.com/spf13/cobra"
)

func (a *App) TaskDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Удалить задачу по ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseTaskID(args[0])
			if err != nil {
				return err
			}
			err = a.taskService.DeleteTask(id)
			if err != nil {
				if errors.Is(err, domain.ErrTaskNotFound) {
					return fmt.Errorf("задача с ID %d не найдена - нечего удалять", id)
				}
				return fmt.Errorf("не удалось выполнить операцию: %w", err)
			}
			fmt.Printf("Задача #%d удалена\n", id)
			return nil
		},
	}
}
