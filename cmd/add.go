package cmd

import (
	"fmt"
	"os"

	"github.com/kimulaco/go-dir/pkg/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Command to add a dir.",
	Long:  "Command to add a dir.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("Required dirName argument.")
			os.Exit(1)
		}

		c, err := config.Init()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		dirName := args[0]
		if !config.IsNotDir(c.GetDir(dirName)) {
			fmt.Println(dirName + " already exists.")
			os.Exit(1)
		}

		wd, wdErr := os.Getwd()
		if wdErr != nil {
			fmt.Println(wdErr.Error())
			os.Exit(1)
		}

		c.Dirs = append(c.Dirs, config.ConfigDir{
			Name: dirName,
			Path: wd,
		})

		configPath, configPathErr := config.GetConfigPath()
		if configPathErr != nil {
			fmt.Println(configPathErr.Error())
			os.Exit(1)
		}

		writeErr := config.WriteConfig(configPath, c)
		if writeErr != nil {
			fmt.Println(writeErr.Error())
			os.Exit(1)
		}

		fmt.Println("Success add " + dirName + ".")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
