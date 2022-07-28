package cmd

import (
	"fmt"
	"os"

	"github.com/kimulaco/go-dir/pkg/array"
	"github.com/kimulaco/go-dir/pkg/config"
	"github.com/kimulaco/go-dir/pkg/dir"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Command to remove a dir.",
	Long:  "Command to remove a dir.",
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
		configDir := c.GetDir(dirName)
		if dir.IsNotDir(configDir) {
			fmt.Println(dirName + " not found.")
			os.Exit(1)
		}

		c.Dirs = array.Filter(c.Dirs, func(_dir dir.Dir) bool {
			return _dir.Name != dirName
		})

		configPath, configPathErr := config.GetConfigFilePath()
		if configPathErr != nil {
			fmt.Println(configPathErr.Error())
			os.Exit(1)
		}

		writeErr := config.WriteConfig(configPath, c)
		if writeErr != nil {
			fmt.Println(writeErr.Error())
			os.Exit(1)
		}

		fmt.Println("Success remove " + dirName + ".")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
