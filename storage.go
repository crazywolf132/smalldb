package smalldb

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// readData reads the JSON data from the file into a map.
func readData[T any](filepath string) (map[string]T, error) {
	data := make(map[string]T)

	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil // Return empty data if file doesn't exist.
		}
		return nil, err
	}

	if len(fileData) == 0 {
		return data, nil // Return empty data if file is empty.
	}

	if err := json.Unmarshal(fileData, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// writeData writes the JSON data to the file.
func writeData[T any](filepath string, data map[string]T) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

// cloneMap creates a shallow copy of the map.
func cloneMap[T any](original map[string]T) map[string]T {
	copy := make(map[string]T, len(original))
	for k, v := range original {
		copy[k] = v
	}
	return copy
}
