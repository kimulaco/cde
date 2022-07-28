package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kimulaco/go-dir/pkg/config"
	"github.com/spf13/cobra"
)

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Command to go to dir.",
	Long:  "Command to go to dir.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("Required dirName argument.")
			os.Exit(1)
		}

		c, initErr := config.Init()
		if initErr != nil {
			fmt.Println(initErr.Error())
			os.Exit(1)
		}

		dirName := args[0]
		dir := config.GetDir(c, dirName)
		if config.IsNotDir(dir) {
			fmt.Println(dirName + " not found.")
			os.Exit(1)
		}

		cdErr := exec.Command("cd " + dir.Path).Run()
		if cdErr != nil {
			fmt.Println(cdErr.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
}
