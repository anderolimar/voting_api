package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(workerCmd)
}

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Vote Worker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Runnig Vote Worker")
	},
}
