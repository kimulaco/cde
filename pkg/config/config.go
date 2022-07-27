package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/kimulaco/go-dir/pkg/datetime"
)

const CONFIG_DIR_NAME string = ".go-dir"
const CONFIG_FILE_NAME string = "config.yml"

type ConfigDir struct {
	Name string
	Path string
}

type Config struct {
	UpdatedAt string      `yaml:"updatedAt"`
	Dirs      []ConfigDir `yaml:"dirs"`
}

func GetConfigDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, CONFIG_DIR_NAME), nil
}

func GetConfigPath() (string, error) {
	configDir, configDirErr := GetConfigDir()
	if configDirErr != nil {
		return "", configDirErr
	}
	return filepath.Join(configDir, CONFIG_FILE_NAME), nil
}

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

func Print(c Config) string {
	var value string

	value += "\n" + "Dirs:"
	for _, dir := range c.Dirs {
		value += "\n" + dir.Name + " " + dir.Path
	}

	value += "\n\n" + "Updated At: " + c.UpdatedAt

	fmt.Println(value)
	return value
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
