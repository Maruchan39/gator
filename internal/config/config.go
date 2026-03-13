package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFileName), nil

}

func Read() (Config, error) {
	var config Config

	configFilePath, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configFilePath)

	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	return config, err

}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	configFilePath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	return os.WriteFile(configFilePath, data, 0644)
}
