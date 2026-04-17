package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (a *App) TaskDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Удалить задачу, по ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("неверный формат ID: %s (должно быть число)", args[0])
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
