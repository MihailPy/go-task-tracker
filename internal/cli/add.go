package cli

import "fmt"
import "github.com/spf13/cobra"

func (a *App) TaskAddCmd() *cobra.Command {
	var desc string
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Добавление задачи. Описание задачи (обязательно)",
		RunE: func(cmd *cobra.Command, args []string) error {
			task, err := a.taskService.AddTask(desc)
			if err != nil {
				return fmt.Errorf("Не удалось добавить задачу: %w", err)
			}

			fmt.Printf("✅ Добавлена задача #%d: %s\n", task.ID, task.Description)
			return nil
		},
	}
	cmd.Flags().StringVarP(&desc, "description", "d", "", "Описание задачи (обязательно)")
	cmd.MarkFlagRequired("description")
	return cmd
}
