package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "service",
	Short: "Voting API Service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Runnig Voting API Service")
	},
}
