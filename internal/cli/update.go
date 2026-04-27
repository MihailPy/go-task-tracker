package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (a *App) TaskUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update [id] [description]",
		Short: "Обновление описания задачи. Id задачи и описание задачи (обязательно)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseTaskID(args[0])
			if err != nil {
				return err
			}
			desc := args[1]
			err = a.taskService.UpdateTaskDescription(id, desc)
			if err != nil {
				return fmt.Errorf("Не удалось обновить описание задачи: %w", err)
			}
			fmt.Printf("\n📝 Описание задачи #%d обновлено, на %s\n", id, desc)
			return nil
		},
	}
}
