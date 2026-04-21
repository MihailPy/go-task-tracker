package cli

import (
	"fmt"
	"task-tracker/internal/domain"

	"io"

	"github.com/spf13/cobra"
)

func printTasks(w io.Writer, header string, tasks []*domain.Task) {
	if len(tasks) == 0 {
		fmt.Fprintln(w, "Список задач пуст")
		return
	}
	fmt.Fprintf(w, "\n📋 %s:\n", header)
	for _, t := range tasks {
		fmt.Fprintf(w, "  #%d: %s [%s]\n", t.ID, t.Description, t.Status)
	}
}

func (a *App) TaskListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Возможен фильтр по статусу (todo, in-progress, done)",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := a.taskService.ListAllTasks()
			if err != nil {
				return err
			}
			printTasks(cmd.OutOrStdout(), "Все задачи", tasks)
			return nil
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
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := a.taskService.ListTasksByStatus(domain.StatusTodo)
			if err != nil {
				return err
			}
			printTasks(cmd.OutOrStdout(), "Задачи в статусе 'todo'", tasks)
			return nil
		},
	}
	return cmd
}
func (a *App) TaskListInProgressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "in-progress",
		Short: "Показать только выполняемые задачи",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := a.taskService.ListTasksByStatus(domain.StatusInProgress)
			if err != nil {
				return err
			}
			printTasks(cmd.OutOrStdout(), "Задачи в статусе 'in-progress'", tasks)
			return nil
		},
	}
	return cmd
}
func (a *App) TaskListDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "done",
		Short: "Показать только выполненные задачи",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := a.taskService.ListTasksByStatus(domain.StatusDone)
			if err != nil {
				return err
			}
			printTasks(cmd.OutOrStdout(), "Задачи в статусе 'done'", tasks)
			return nil
		},
	}
	return cmd
}
