package config

import (
	"encoding/json"
	"testing"
)

func TestIsNotDir(t *testing.T) {
	noDir := ConfigDir{}
	if !IsNotDir(noDir) {
		v, _ := json.Marshal(noDir)
		t.Error("Recieve:\n" + string(v))
	}

	testDir := ConfigDir{Name: "test", Path: "/test/dir"}
	if IsNotDir(testDir) {
		v, _ := json.Marshal(testDir)
		t.Error("Recieve:\n" + string(v))
	}
}

func TestGetDir(t *testing.T) {
	c := Config{
		UpdatedAt: "",
		Dirs: []ConfigDir{
			{Name: "test-1", Path: "/test/dir-1"},
			{Name: "test-2", Path: "/test/dir-2"},
		},
	}

	dir1 := GetDir(c, "test-1")
	if dir1.Name != "test-1" || dir1.Path != "/test/dir-1" {
		v, _ := json.Marshal(dir1)
		t.Error("Recieve:\n" + string(v))
	}

	dir2 := GetDir(c, "test-2")
	if dir2.Name != "test-2" || dir2.Path != "/test/dir-2" {
		v, _ := json.Marshal(dir2)
		t.Error("Recieve:\n" + string(v))
	}

	dir3 := GetDir(c, "test-3")
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
	dir := GetDir(c, "test")
	if dir.Name != "" || dir.Path != "" {
		v, _ := json.Marshal(dir)
		t.Error("Recieve:\n" + string(v))
	}
}
