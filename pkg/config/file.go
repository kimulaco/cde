package config

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

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

func GetConfigDirPath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, DEFAULT_CONFIG_DIR_NAME), nil
}

func GetConfigFilePath() (string, error) {
	configDir, configDirErr := GetConfigDirPath()
	if configDirErr != nil {
		return "", configDirErr
	}
	return filepath.Join(configDir, DEFAULT_CONFIG_FILE_NAME), nil
}
