package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestStatFile(t *testing.T) {
	wd, _ := os.Getwd()

	fp := filepath.Join(wd, "../../test/config/data/config.yml")
	isFile, fileErr := StatFile(fp)
	if fileErr != nil {
		t.Error("Error: " + fileErr.Error())
	}
	if !isFile {
		t.Error("Recieve: false")
	}

	dp := filepath.Join(wd, "../../test/config/data")
	isFile, dirErr := StatFile(dp)
	if dirErr != nil {
		t.Error("Error: " + dirErr.Error())
	}
	if isFile {
		t.Error("Recieve: true")
	}

	// TODO: Non-existent path
}

func TestStatDir(t *testing.T) {
	wd, _ := os.Getwd()

	dp := filepath.Join(wd, "../../test/config/data")
	isDir, fileErr := StatDir(dp)
	if fileErr != nil {
		t.Error("Error: " + fileErr.Error())
	}
	if !isDir {
		t.Error("Recieve: false")
	}

	fp := filepath.Join(wd, "../../test/config/data/config.yml")
	isDir, dirErr := StatDir(fp)
	if dirErr != nil {
		t.Error("Error: " + dirErr.Error())
	}
	if isDir {
		t.Error("Recieve: true")
	}

	// TODO: Non-existent path
}

func TestReadConfig(t *testing.T) {
	wd, _ := os.Getwd()
	fp := filepath.Join(wd, "../../test/config/data/config.yml")
	c, readErr := ReadConfig(fp)

	if readErr != nil {
		t.Error("Error: " + readErr.Error())
	}
	if c.UpdatedAt != "2022-01-01 01:01:00" || len(c.Dirs) != 2 {
		v, _ := json.Marshal(c)
		t.Error("Recieve:\n" + string(v))
	}

	// TODO: Non-existent path
}

func TestWriteConfig(t *testing.T) {
	wd, _ := os.Getwd()
	fp := filepath.Join(wd, "../../test/config/data/dist/config.yml")
	c := Config{
		UpdatedAt: "2022-01-01 01:01:00",
		Dirs:      make([]ConfigDir, 0),
	}
	writeErr := WriteConfig(fp, c)
	if writeErr != nil {
		t.Error("Error: " + writeErr.Error())
	}
	isFile, fileErr := StatFile(fp)
	if fileErr != nil {
		t.Error("Error: " + fileErr.Error())
	}
	if !isFile {
		t.Error("Recieve: false")
	}
}
