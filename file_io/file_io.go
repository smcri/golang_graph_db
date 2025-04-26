package file_io

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadFile(filename string) (map[string]interface{}, error) {
	fileContent, err := os.ReadFile(filename + ".json")
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(fileContent, &data); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	return data, nil
}

func WriteFile(key string, value interface{}, filename string) error {
	// Try to read the file
	fileContent, err := os.ReadFile(filename + ".json")
	var jsonData map[string]interface{}

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// File does not exist, create a new map
			jsonData = make(map[string]interface{})
		} else {
			// Other unexpected errors
			return fmt.Errorf("could not read file: %w", err)
		}
	} else {
		// File exists, unmarshal the content
		if err := json.Unmarshal(fileContent, &jsonData); err != nil {
			return fmt.Errorf("could not unmarshal JSON: %w", err)
		}
	}

	// Update jsonData
	if key == "*" {
		jsonData, _ = value.(map[string]interface{})
	} else {
		jsonData[key] = value
	}

	// Marshal updated data
	jsonString, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal JSON string: %w", err)
	}

	// Write back to file
	if err := os.WriteFile(filename+".json", jsonString, 0644); err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	return nil
}

func DeleteFile(filename string) error {
	if err := os.Remove(filename + ".json"); err != nil {
		return fmt.Errorf("could not delete file: %w", err)
	}
	return nil
}
