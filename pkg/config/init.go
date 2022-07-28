package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/kimulaco/go-dir/pkg/datetime"
	"github.com/kimulaco/go-dir/pkg/dir"
	"github.com/kimulaco/go-dir/pkg/fs"
)

func Init() (Config, error) {
	var c Config
	configDirPath, configDirPathErr := GetConfigDirPath()
	if configDirPathErr != nil {
		return c, configDirPathErr
	}

	hasConfigDir, _ := fs.StatDir(configDirPath)
	if !hasConfigDir {
		makeConfigDirErr := os.MkdirAll(configDirPath, os.ModePerm)
		if makeConfigDirErr != nil {
			return c, makeConfigDirErr
		}
	}

	configFilePath := filepath.Join(configDirPath, DEFAULT_CONFIG_FILE_NAME)
	hasConfigFile, _ := fs.StatFile(configFilePath)
	if hasConfigFile {
		_config, readErr := ReadConfig(configFilePath)
		c = _config
		if readErr != nil {
			return c, readErr
		}
	} else {
		_config, createErr := createDefaultConfig(configFilePath)
		c = _config
		if createErr != nil {
			return c, createErr
		}
	}

	return c, nil
}

func createDefaultConfig(configFilePath string) (Config, error) {
	config := Config{
		UpdatedAt: datetime.TimeToString(time.Now()),
		Dirs:      make([]dir.Dir, 0),
	}
	writeErr := WriteConfig(configFilePath, config)
	if writeErr != nil {
		fmt.Println(writeErr.Error())
		return config, writeErr
	}
	return config, nil
}
