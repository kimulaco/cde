package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kimulaco/cde/pkg/config"
	"github.com/kimulaco/cde/pkg/dir"
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
		configDir := c.GetDir(dirName)
		if dir.IsNotDir(configDir) {
			fmt.Println(dirName + " not found.")
			os.Exit(1)
		}

		cdErr := exec.Command("cd " + configDir.Path).Run()
		if cdErr != nil {
			fmt.Println(cdErr.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
}
