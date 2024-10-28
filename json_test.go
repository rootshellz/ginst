package ginst

import (
	"testing"
)

// TestJSON tests marshaling valid data to JSON format.
func TestJSON(t *testing.T) {
	data := map[string]string{"key": "value"}
	jsonData, err := JSON(data)
	if err != nil {
		t.Fatalf("Expected JSON marshaling to succeed, got error: %v", err)
	}

	expected := "{\n  \"key\": \"value\"\n}"
	if string(jsonData) != expected {
		t.Errorf("Expected JSON output '%s', got '%s'", expected, string(jsonData))
	}
}

// TestJSONUnmarshalable tests handling of unmarshalable data in JSON().
func TestJSONUnmarshalable(t *testing.T) {
	type unmarshalable struct {
		Ch chan int
	}

	data := unmarshalable{Ch: make(chan int)}
	_, err := JSON(data)
	if err == nil {
		t.Error("Expected JSON marshaling to fail for unmarshalable data, got no error")
	}
}
