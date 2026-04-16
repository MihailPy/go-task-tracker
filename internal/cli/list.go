package cli

import (
	"fmt"
	"task-tracker/internal/domain"

	"github.com/spf13/cobra"
)

func (a *App) TaskListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Возможен фильтр по статусу (todo, in-progress, done)",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, _ := a.taskService.ListAllTasks()
			fmt.Println("\n📋 Все задачи:")
			for _, t := range tasks {
				fmt.Printf("  #%d: %s [%s]\n", t.ID, t.Description, t.Status)
			}
		},
	}
	cmd.AddCommand(a.TaskListDoneCmd())
	cmd.AddCommand(a.TaskListTodoCmd())
	cmd.AddCommand(a.TaskListInProgressCmd())
	return cmd
}

func (a *App) TaskListTodoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "todo",
		Short: "Показать только не начатые задачи",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, _ := a.taskService.ListTasksByStatus(domain.StatusTodo)
			fmt.Println("\n📋 Задачи в статусе 'todo':")
			for _, t := range tasks {
				fmt.Printf("  #%d: %s [%s]\n", t.ID, t.Description, t.Status)
			}
		},
	}
	return cmd
}
func (a *App) TaskListInProgressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "in-progress",
		Short: "Показать только выполняемые задачи",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, _ := a.taskService.ListTasksByStatus(domain.StatusInProgress)
			fmt.Println("\n📋 Задачи в статусе 'in-progress':")
			for _, t := range tasks {
				fmt.Printf("  #%d: %s [%s]\n", t.ID, t.Description, t.Status)
			}
		},
	}
	return cmd
}
func (a *App) TaskListDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "done",
		Short: "Показать только выполненные задачи",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, _ := a.taskService.ListTasksByStatus(domain.StatusDone)
			fmt.Println("\n📋 Задачи в статусе 'done':")
			for _, t := range tasks {
				fmt.Printf("  #%d: %s [%s]\n", t.ID, t.Description, t.Status)
			}
		},
	}
	return cmd
}
