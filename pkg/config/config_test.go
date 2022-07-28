package config

import (
	"encoding/json"
	"testing"
)

func TestGetDir(t *testing.T) {
	c := Config{
		UpdatedAt: "",
		Dirs: []ConfigDir{
			{Name: "test-1", Path: "/test/dir-1"},
			{Name: "test-2", Path: "/test/dir-2"},
		},
	}

	dir1 := c.GetDir("test-1")
	if dir1.Name != "test-1" || dir1.Path != "/test/dir-1" {
		v, _ := json.Marshal(dir1)
		t.Error("Recieve:\n" + string(v))
	}

	dir2 := c.GetDir("test-2")
	if dir2.Name != "test-2" || dir2.Path != "/test/dir-2" {
		v, _ := json.Marshal(dir2)
		t.Error("Recieve:\n" + string(v))
	}

	dir3 := c.GetDir("test-3")
	if dir3.Name != "" || dir3.Path != "" {
		v, _ := json.Marshal(dir3)
		t.Error("Recieve:\n" + string(v))
	}
}

func TestGetDirNoDir(t *testing.T) {
	c := Config{
		UpdatedAt: "",
		Dirs:      make([]ConfigDir, 0),
	}
	dir := c.GetDir("test")
	if dir.Name != "" || dir.Path != "" {
		v, _ := json.Marshal(dir)
		t.Error("Recieve:\n" + string(v))
	}
}

func TestConfigPrint(t *testing.T) {
	c := Config{
		UpdatedAt: "2022-01-01 12:00:00",
		Dirs: []ConfigDir{
			{Name: "test-1", Path: "/test/dir-1"},
			{Name: "test-2", Path: "/test/dir-2"},
		},
	}
	result := c.Print()
	expect := `go-dir

Dirs:
test-1 /test/dir-1
test-2 /test/dir-2

Updated At: 2022-01-01 12:00:00`

	if result != expect {
		t.Error("Recieve: \n" + result)
	}
}

func TestPrintNoDir(t *testing.T) {
	c := Config{
		UpdatedAt: "2022-01-01 12:00:00",
		Dirs:      []ConfigDir{},
	}
	result := c.Print()
	expect := `go-dir

Dirs:

Updated At: 2022-01-01 12:00:00`

	if result != expect {
		t.Error("Recieve: \n" + result)
	}
}

func TestConfigPrintDirs(t *testing.T) {
	c := Config{
		UpdatedAt: "2022-01-01 12:00:00",
		Dirs: []ConfigDir{
			{Name: "test-1", Path: "/test/dir-1"},
			{Name: "test-2", Path: "/test/dir-2"},
		},
	}
	result := c.PrintDirs()
	expect := `
test-1 /test/dir-1
test-2 /test/dir-2`

	if result != expect {
		t.Error("Recieve: \n" + result)
	}
}
