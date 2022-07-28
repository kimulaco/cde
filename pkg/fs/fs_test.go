package fs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStatFileOK(t *testing.T) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "../../test/config/data/config.yml")
	isFile, statErr := StatFile(path)
	if statErr != nil {
		t.Error("Error: " + statErr.Error())
	}
	if !isFile {
		t.Error("Recieve: false")
	}
}

func TestStatFileDir(t *testing.T) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "../../test/config/data")
	isFile, statErr := StatFile(path)
	if statErr != nil {
		t.Error("Error: " + statErr.Error())
	}
	if isFile {
		t.Error("Recieve: true")
	}
}

func TestStatFileNotFound(t *testing.T) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "../../test/config/notfound/config.yml")
	isFile, statErr := StatFile(path)
	if statErr == nil {
		t.Error("Error: nil")
	}
	if isFile {
		t.Error("Recieve: true")
	}
}

func TestStatDirOK(t *testing.T) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "../../test/config/data")
	isDir, statErr := StatDir(path)
	if statErr != nil {
		t.Error("Error: " + statErr.Error())
	}
	if !isDir {
		t.Error("Recieve: false")
	}
}

func TestStatDirFile(t *testing.T) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "../../test/config/data/config.yml")
	isDir, statErr := StatDir(path)
	if statErr != nil {
		t.Error("Error: " + statErr.Error())
	}
	if isDir {
		t.Error("Recieve: true")
	}
}

func TestStatDirNotFound(t *testing.T) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "../../test/config/notfound")
	isDir, statErr := StatDir(path)
	if statErr == nil {
		t.Error("Error: nil")
	}
	if isDir {
		t.Error("Recieve: true")
	}
}
