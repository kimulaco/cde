package dir

type Dir struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

func IsNotDir(d Dir) bool {
	return d.Name == "" || d.Path == ""
}
