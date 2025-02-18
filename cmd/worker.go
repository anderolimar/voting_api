package cmd

import (
	//"context"
	"context"
	"fmt"
	"votingapi/worker"
	//"votingapi/worker"

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
		worker.Run(context.Background())

	},
}
