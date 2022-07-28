package config

func IsNotDir(d ConfigDir) bool {
	return d.Name == "" || d.Path == ""
}

func GetDir(c Config, name string) ConfigDir {
	for _, dir := range c.Dirs {
		if dir.Name == name {
			return dir
		}
	}
	return ConfigDir{}
}
