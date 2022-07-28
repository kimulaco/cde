package config

import (
	"fmt"
)

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
