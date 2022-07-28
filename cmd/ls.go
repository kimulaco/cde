package cmd

import (
	"fmt"
	"os"

	"github.com/kimulaco/go-dir/pkg/config"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Command to list dirs.",
	Long:  "Command to list dirs.",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.Init()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		config.PrintDirs(c)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
