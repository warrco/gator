package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	var cfg Config
	filepath, err := getConfigFilePath()
	if err != nil {
		return cfg, fmt.Errorf("could not find a home directory for the user: %w", err)
	}

	file, err := os.Open(filepath)
	if err != nil {
		return cfg, fmt.Errorf("could not read the file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("error decoding json: %w", err)
	}
	return cfg, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(*c)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find a home directory for the user: %w", err)
	}

	filepath := filepath.Join(homedir, configFileName)
	return filepath, nil
}

func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("could not retrieve file path: %w", err)
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}
