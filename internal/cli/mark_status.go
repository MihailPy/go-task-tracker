package cli

import (
	"errors"
	"fmt"
	"task-tracker/internal/domain"

	"github.com/spf13/cobra"
)

func (a *App) TaskMarkInProgressCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mark-in-progress [id]",
		Short: "Перевести задачу в статус \"in-progress\"",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseTaskID(args[0])
			if err != nil {
				return err
			}
			err = a.taskService.UpdateTaskStatus(id, domain.StatusInProgress)
			if err != nil {
				if errors.Is(err, domain.ErrTaskNotFound) {
					return fmt.Errorf("ошибка: задача #%d не существует", id)
				}
				return fmt.Errorf("не удалось обновить статус задачи #%d : %w", id, err)
			}
			fmt.Printf("Статус задачи #%d изменён на %s\n", id, domain.StatusInProgress)
			return nil
		},
	}
}

func (a *App) TaskMarkDoneCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mark-done [id]",
		Short: "Перевести задачу в статус \"done\"",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseTaskID(args[0])
			if err != nil {
				return err
			}
			err = a.taskService.UpdateTaskStatus(id, domain.StatusDone)
			if err != nil {
				if errors.Is(err, domain.ErrTaskNotFound) {
					return fmt.Errorf("ошибка: задача #%d не существует", id)
				}
				return fmt.Errorf("не удалось обновить статус задачи #%d : %w", id, err)
			}
			fmt.Printf("Статус задачи #%d изменён на %s\n", id, domain.StatusDone)
			return nil
		},
	}
}
