package cmd

import (
	"github.com/MSLacerda/ES2-Project-Backend/internal"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs ES2-Project-Backend server",
	Run: func(cmd *cobra.Command, args []string) {
		internal.BuildApp()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
