package cmd

import (
	"fmt"
	"os"

	"github.com/kimulaco/go-dir/pkg/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dir",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.Init()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		config.Print(c)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
