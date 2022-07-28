package fs

import (
	"os"
)

func StatFile(path string) (bool, error) {
	info, err := os.Stat(path)
	os.IsNotExist(err)
	if err != nil {
		return false, err
	}
	if info.IsDir() {
		return false, err
	}
	return true, nil
}

func StatDir(path string) (bool, error) {
	info, err := os.Stat(path)
	os.IsNotExist(err)
	if err != nil {
		return false, err
	}
	if !info.IsDir() {
		return false, err
	}
	return true, nil
}
