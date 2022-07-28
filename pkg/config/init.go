package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/kimulaco/go-dir/pkg/datetime"
)

func Init() (Config, error) {
	var config Config
	configDir, configDirErr := GetConfigDir()
	if configDirErr != nil {
		return config, configDirErr
	}

	hasConfigDir, _ := StatDir(configDir)
	if !hasConfigDir {
		makeConfigDirErr := os.MkdirAll(configDir, os.ModePerm)
		if makeConfigDirErr != nil {
			return config, makeConfigDirErr
		}
	}

	configFilePath := filepath.Join(configDir, CONFIG_FILE_NAME)
	hasConfigFile, _ := StatFile(configFilePath)
	if hasConfigFile {
		_config, readErr := ReadConfig(configFilePath)
		config = _config
		if readErr != nil {
			return _config, readErr
		}
	} else {
		_config, createErr := createDefaultConfig(configFilePath)
		config = _config
		if createErr != nil {
			return _config, createErr
		}
	}

	return config, nil
}

func createDefaultConfig(configFilePath string) (Config, error) {
	config := Config{
		UpdatedAt: datetime.TimeToString(time.Now()),
		Dirs:      make([]ConfigDir, 0),
	}
	writeErr := WriteConfig(configFilePath, config)
	if writeErr != nil {
		fmt.Println(writeErr.Error())
		return config, writeErr
	}
	return config, nil
}
