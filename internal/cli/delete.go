package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (a *App) TaskDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Удалить задачу, по ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseTaskID(args[0])
			if err != nil {
				return err
			}
			err = a.taskService.DeleteTask(id)
			if err != nil {
				return fmt.Errorf("Не удалось удалить задачу #%d : %w", id, err)
			}
			fmt.Printf("\n🔄 Задача #%d удалена \n", id)
			return nil
		},
	}
}
