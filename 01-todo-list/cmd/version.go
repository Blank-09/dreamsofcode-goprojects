package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of Tasks",
	Long:  "A longer description that spans multiple lines and likely contains examples and usage of using your application.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tasks v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
