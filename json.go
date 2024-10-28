package ginst

import (
	"encoding/json"
)

// ToJSON marshals the input data into an indented JSON byte slice.
// Returns an error if marshalling fails.
func JSON(data any) ([]byte, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
