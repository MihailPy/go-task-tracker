package cli

import (
	"fmt"
	"task-tracker/internal/domain"

	"github.com/spf13/cobra"
)

func (a *App) TaskMarkInProgressCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mark-in-progress [id]",
		Short: "Отметить задачу in-progress (в прогрессе)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseTaskID(args[0])
			if err != nil {
				return err
			}
			err = a.taskService.UpdateTaskStatus(id, domain.StatusInProgress)
			if err != nil {
				return fmt.Errorf("Не удалось отметить in-progress задачу #%d : %w", id, err)
			}
			fmt.Printf("\n🔄 Статус задачи #%d обновлён на %s\n", id, domain.StatusInProgress)
			return nil
		},
	}
}

func (a *App) TaskMarkDoneCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mark-done [id]",
		Short: "Отметить задачу done (выполнено)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseTaskID(args[0])
			if err != nil {
				return err
			}
			err = a.taskService.UpdateTaskStatus(id, domain.StatusDone)
			if err != nil {
				return fmt.Errorf("Не удалось отметить done задачу #%d : %w", id, err)
			}
			fmt.Printf("\n🔄 Статус задачи #%d обновлён на %s\n", id, domain.StatusDone)
			return nil
		},
	}
}
