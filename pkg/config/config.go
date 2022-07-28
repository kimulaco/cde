package config

import (
	"fmt"

	"github.com/kimulaco/go-dir/pkg/dir"
)

const DEFAULT_CONFIG_DIR_NAME string = ".go-dir"
const DEFAULT_CONFIG_FILE_NAME string = "config.yml"

type Config struct {
	UpdatedAt string    `yaml:"updatedAt"`
	Dirs      []dir.Dir `yaml:"dirs"`
}

func (c Config) GetDir(name string) dir.Dir {
	for _, dir := range c.Dirs {
		if dir.Name == name {
			return dir
		}
	}
	return dir.Dir{}
}

func Print(c Config) string {
	value := "go-dir"

	value += "\n\n" + "Dirs:"
	for _, dir := range c.Dirs {
		value += "\n" + dir.Name + " " + dir.Path
	}

	value += "\n\n" + "Updated At: " + c.UpdatedAt

	fmt.Println(value)
	return value
}

func PrintDirs(c Config) string {
	var value string
	for _, dir := range c.Dirs {
		value += "\n" + dir.Name + " " + dir.Path
	}

	fmt.Println(value)
	return value
}
