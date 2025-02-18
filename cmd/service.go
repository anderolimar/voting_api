package cmd

import (
	"fmt"
	"votingapi/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "service",
	Short: "Voting API Service",
	Run: func(cmd *cobra.Command, args []string) {
		service.Start()
		fmt.Println("Runnig Voting API Service")
	},
}
