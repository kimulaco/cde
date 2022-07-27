package config

import (
	"fmt"
)

func IsNotDir(d ConfigDir) bool {
	return d.Name == "" || d.Path == ""
}

func GetDir(c Config, name string) ConfigDir {
	for _, dir := range c.Dirs {
		fmt.Printf("%+v", dir)
		if dir.Name == name {
			fmt.Printf("%+v", name)
			return dir
		}
	}
	return ConfigDir{}
}
