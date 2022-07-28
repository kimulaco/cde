package config

type ConfigDir struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

func IsNotDir(d ConfigDir) bool {
	return d.Name == "" || d.Path == ""
}
