package cli

import (
	"fmt"
	"task-tracker/internal/service"

	"github.com/spf13/cobra"
)

type App struct {
	taskService service.TaskService
}

func NewApp(service service.TaskService) *App {
	return &App{taskService: service}
}

func (a *App) Execute() error {
	rootCmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task-tracker - это менеджер задач",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
████████  █████  ███████ ██   ██       ████████ ██████   █████   ██████ ██   ██ ███████ ██████  
   ██    ██   ██ ██      ██  ██           ██    ██   ██ ██   ██ ██      ██  ██  ██      ██   ██ 
   ██    ███████ ███████ █████   █████    ██    ██████  ███████ ██      █████   █████   ██████  
   ██    ██   ██      ██ ██  ██           ██    ██   ██ ██   ██ ██      ██  ██  ██      ██   ██ 
   ██    ██   ██ ███████ ██   ██          ██    ██   ██ ██   ██  ██████ ██   ██ ███████ ██   ██ 
                                                                                                
			`)
			fmt.Println("Enert help for help about command.")

		},
	}
	rootCmd.AddCommand(a.newListCmd())
	return rootCmd.Execute()
}
