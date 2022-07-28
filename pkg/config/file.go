package config

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
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

func ReadConfig(filePath string) (Config, error) {
	var config Config
	bytes, readErr := ioutil.ReadFile(filePath)
	if readErr != nil {
		return config, readErr
	}

	unmarshalErr := yaml.Unmarshal(bytes, &config)
	if unmarshalErr != nil {
		return config, unmarshalErr
	}

	return config, nil
}

func WriteConfig(filePath string, data Config) error {
	buf, marshalErr := yaml.Marshal(data)
	if marshalErr != nil {
		return marshalErr
	}

	writeErr := ioutil.WriteFile(filePath, buf, os.ModePerm)
	if writeErr != nil {
		return writeErr
	}

	return nil
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
