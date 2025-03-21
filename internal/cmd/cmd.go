package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/utkarshkrsingh/ripple/internal/command"
	"github.com/utkarshkrsingh/ripple/internal/config"
	"github.com/utkarshkrsingh/ripple/internal/log"
)

var rootCmd = &cobra.Command{
	Use:   "ripple",
	Short: "Ripple is a lightweight cli automation tool",
	Long:  `Ripple is a lightweight cli automation tool written in go`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Run the given task",
	Long: `Run the given task if specified in ripple.yaml
    Usage:
    ripple task -r <task-name>      Runs a task by it's name
    ripple task -s                  Shows all the task available

    For example:
    ripple task -r build
    ripple task -r deploy`,

	Run: tasks,
}

func init() {
	rootCmd.AddCommand(taskCmd)
	taskCmd.Flags().StringP("name", "n", "", "Run task by name")
	taskCmd.Flags().BoolP("show", "s", false, "Shows the task available")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func tasks(cmd *cobra.Command, argv []string) {

	if err := config.InitReadConfig(); err != nil {
		log.Logger.Error(err)
		log.Logger.Errorf("terminating execution...")
		os.Exit(1)
	}

	show, err := cmd.Flags().GetBool("show")
	taskName, err := cmd.Flags().GetString("name")

	if err != nil {
		log.Logger.Error(err)
	}

	log.Logger.Info("Ripple started successfully...")
	if show {
		command.ShowAllTasks()
	}

	if len(taskName) > 0 {
		err = command.RunByName(taskName)
		if err != nil {
			log.Logger.Error(err)
		}
	}
}
