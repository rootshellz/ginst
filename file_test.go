package ginst

import (
	"bytes"
	"os"
	"testing"
)

// TestWriteToFileBytes tests writing a byte slice to a file.
func TestWriteToFileBytes(t *testing.T) {
	path := "/tmp/testfile_bytes.txt"
	defer os.Remove(path)

	WriteToFile([]byte("test data"), path)

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Expected to read file, got error: %v", err)
	}

	if !bytes.Equal(content, []byte("test data")) {
		t.Errorf("Expected content 'test data', got '%s'", string(content))
	}
}

// TestWriteToFileString tests writing a string to a file.
func TestWriteToFileString(t *testing.T) {
	path := "/tmp/testfile_string.txt"
	defer os.Remove(path)

	WriteToFile("test data", path)

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Expected to read file, got error: %v", err)
	}

	if string(content) != "test data" {
		t.Errorf("Expected content 'test data', got '%s'", string(content))
	}
}

// TestWriteToFileJSON tests writing a JSON-marshalable struct to a file.
func TestWriteToFileJSON(t *testing.T) {
	path := "/tmp/testfile_json.txt"
	defer os.Remove(path)

	data := map[string]string{"key": "value"}
	WriteToFile(data, path)

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Expected to read file, got error: %v", err)
	}

	expected := "{\n  \"key\": \"value\"\n}"
	if string(content) != expected {
		t.Errorf("Expected JSON content %s, got '%s'", expected, string(content))
	}
}

// TestWriteToFileUnmarshalable tests handling of unmarshalable data.
func TestWriteToFileUnmarshalable(t *testing.T) {
	path := "/tmp/testfile_unmarshalable.txt"
	defer os.Remove(path)

	type unmarshalable struct {
		Ch chan int
	}
	data := unmarshalable{Ch: make(chan int)}
	WriteToFile(data, path)

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Expected to read file, got error: %v", err)
	}

	// Check if default string representation is used
	if string(content) == "" {
		t.Error("Expected non-empty content for unmarshalable data")
	}
}
