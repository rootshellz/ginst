package ginst

import (
	"os"
	"testing"
)

// TestSpewToFile verifies SpewToFile writes detailed data dump to the file.
func TestSpewToFile(t *testing.T) {
	path := "spew_test_output.txt"
	defer os.Remove(path) // Clean up after test

	SpewToFile(map[string]int{"key": 1}, path)

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Expected to read file, got error: %v", err)
	}

	if len(content) == 0 {
		t.Error("Expected non-empty output from SpewToFile")
	}
}
