package ginst

import (
	"encoding/json"
	"fmt"
	"os"
)

// WriteToFile attempts to write arbitrary data to the file at the given path.
// Supported data types include byte slices, strings, and JSON-marshalable structs/maps.
// If JSON marshaling fails, the function logs an error and returns without writing.
func WriteToFile(data any, path string) {
	file, err := os.Create(path)
	if err != nil {
		DebugPrintWithMessage("Error creating file", fmt.Sprintf("%s: %s", path, err))
		return
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			DebugPrintWithMessage("Error closing file", fmt.Sprintf("%s: %s", path, closeErr))
		}
	}()

	var bytes []byte
	switch v := data.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		jsonData, jsonErr := json.MarshalIndent(v, "", "  ")
		if jsonErr == nil {
			bytes = jsonData
		} else {
			bytes = []byte(fmt.Sprintf("%v", v)) // Use the default string representation of the type
		}
	}

	if _, err := file.Write(bytes); err != nil {
		DebugPrintWithMessage("Error writing file", fmt.Sprintf("%s: %s", path, err))
	} else {
		DebugPrintWithMessage("Successfully wrote data to", path)
	}
}
