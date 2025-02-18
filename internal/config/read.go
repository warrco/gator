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

	filecontent, err := os.ReadFile(filepath)
	if err != nil {
		return cfg, fmt.Errorf("could not read the file: %w", err)
	}

	err = json.Unmarshal(filecontent, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("error unmarshalling json: %w", err)
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

	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	err = os.WriteFile(filepath, data, 0600)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}
