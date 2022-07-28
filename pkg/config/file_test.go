package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/kimulaco/cde/pkg/dir"
	"github.com/kimulaco/cde/pkg/fs"
)

func TestReadConfig(t *testing.T) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "../../test/config/data/config.yml")
	c, readErr := ReadConfig(path)

	if readErr != nil {
		t.Error("Error: " + readErr.Error())
	}
	if c.UpdatedAt != "2022-01-01 01:01:00" || len(c.Dirs) != 2 {
		v, _ := json.Marshal(c)
		t.Error("Recieve:\n" + string(v))
	}
}

func TestReadConfigNotfound(t *testing.T) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "../../test/config/notfound/config.yml")
	c, readErr := ReadConfig(path)

	if readErr == nil {
		t.Error("Error: nil")
	}
	if c.UpdatedAt != "" || len(c.Dirs) != 0 {
		v, _ := json.Marshal(c)
		t.Error("Recieve:\n" + string(v))
	}
}

func TestWriteConfig(t *testing.T) {
	wd, _ := os.Getwd()
	fp := filepath.Join(wd, "../../test/config/data/dist/config.yml")
	c := Config{
		UpdatedAt: "2022-01-01 01:01:00",
		Dirs:      []dir.Dir{},
	}
	writeErr := WriteConfig(fp, c)
	if writeErr != nil {
		t.Error("Error: " + writeErr.Error())
	}
	isFile, fileErr := fs.StatFile(fp)
	if fileErr != nil {
		t.Error("Error: " + fileErr.Error())
	}
	if !isFile {
		t.Error("Recieve: false")
	}
}

func TestGetConfigDirPath(t *testing.T) {
	configDirPath, err := GetConfigDirPath()
	if err != nil {
		t.Error("Error: " + err.Error())
	}
	if configDirPath == "" {
		t.Error("Recieve: " + configDirPath)
	}
}

func TestGetConfigFilePath(t *testing.T) {
	configFilePath, err := GetConfigFilePath()
	if err != nil {
		t.Error("Error: " + err.Error())
	}
	if configFilePath == "" {
		t.Error("Recieve: " + configFilePath)
	}
}
