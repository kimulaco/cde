package config

import (
	"testing"
)

func TestPrint(t *testing.T) {
	c := Config{
		UpdatedAt: "2022-01-01 12:00:00",
		Dirs: []ConfigDir{
			{Name: "test-1", Path: "/test/dir-1"},
			{Name: "test-2", Path: "/test/dir-2"},
		},
	}
	result := Print(c)
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
	result := Print(c)
	expect := `go-dir

Dirs:

Updated At: 2022-01-01 12:00:00`

	if result != expect {
		t.Error("Recieve: \n" + result)
	}
}
