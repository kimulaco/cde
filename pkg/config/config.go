package config

import (
	"fmt"
)

const CONFIG_DIR_NAME string = ".go-dir"
const CONFIG_FILE_NAME string = "config.yml"

type Config struct {
	UpdatedAt string      `yaml:"updatedAt"`
	Dirs      []ConfigDir `yaml:"dirs"`
}

func (c Config) GetDir(name string) ConfigDir {
	for _, dir := range c.Dirs {
		if dir.Name == name {
			return dir
		}
	}
	return ConfigDir{}
}

func (c Config) Print() string {
	value := "go-dir"

	value += "\n\n" + "Dirs:"
	for _, dir := range c.Dirs {
		value += "\n" + dir.Name + " " + dir.Path
	}

	value += "\n\n" + "Updated At: " + c.UpdatedAt

	fmt.Println(value)
	return value
}

func (c Config) PrintDirs() string {
	var value string
	for _, dir := range c.Dirs {
		value += "\n" + dir.Name + " " + dir.Path
	}

	fmt.Println(value)
	return value
}
