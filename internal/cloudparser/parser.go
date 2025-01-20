package cloudparser

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadCloudEnvironment(filePath string) (*CloudEnvironment, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var cloudEnv CloudEnvironment
	if err := json.NewDecoder(file).Decode(&cloudEnv); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &cloudEnv, nil
}
