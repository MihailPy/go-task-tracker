package cli

import (
	"fmt"
	"task-tracker/internal/domain"

	"github.com/spf13/cobra"
)

type TaskUseCase interface {
	AddTask(description string) (*domain.Task, error)
	UpdateTaskStatus(id int, status domain.TaskStatus) error
	UpdateTaskDescription(id int, desc string) error
	DeleteTask(id int) error
	ListAllTasks() ([]*domain.Task, error)
	ListTasksByStatus(status domain.TaskStatus) ([]*domain.Task, error)
}
type App struct {
	taskService TaskUseCase
}

func NewApp(svc TaskUseCase) *App {
	return &App{taskService: svc}
}

func (a *App) Execute() error {
	rootCmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task Tracker - менеджер задач",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`
████████  █████  ███████ ██   ██       ████████ ██████   █████   ██████ ██   ██ ███████ ██████  
   ██    ██   ██ ██      ██  ██           ██    ██   ██ ██   ██ ██      ██  ██  ██      ██   ██ 
   ██    ███████ ███████ █████   █████    ██    ██████  ███████ ██      █████   █████   ██████  
   ██    ██   ██      ██ ██  ██           ██    ██   ██ ██   ██ ██      ██  ██  ██      ██   ██ 
   ██    ██   ██ ███████ ██   ██          ██    ██   ██ ██   ██  ██████ ██   ██ ███████ ██   ██ 
                                                                                                
			`)
			fmt.Printf("Введите \"help\" для списка команд.\n")

		},
	}
	rootCmd.AddCommand(a.TaskListCmd())
	rootCmd.AddCommand(a.TaskAddCmd())
	rootCmd.AddCommand(a.TaskUpdateCmd())
	rootCmd.AddCommand(a.TaskMarkInProgressCmd())
	rootCmd.AddCommand(a.TaskMarkDoneCmd())
	rootCmd.AddCommand(a.TaskDeleteCmd())
	return rootCmd.Execute()
}
