package core

import (
	"encoding/json"
	"os"
)

type Config struct {
	Stacks []StackConfig `json:"stacks"`
}

type StackConfig struct {
	FoldersToRemove    []string `json:"folders_to_remove"`
	ProjectFilePattern string   `json:"project_file_pattern"`
}

func ReadConfig(filePath string) (*Config, error) {
	var config *Config

	bytes, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
