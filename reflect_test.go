package ginst

import (
	"reflect"
	"testing"
)

// TestTypeOf checks TypeOf function with various types.
func TestTypeOf(t *testing.T) {
	tests := []struct {
		input    any
		expected reflect.Type
	}{
		{input: 5, expected: reflect.TypeOf(5)},
		{input: "test", expected: reflect.TypeOf("test")},
		{input: map[string]int{"key": 1}, expected: reflect.TypeOf(map[string]int{})},
	}

	for _, tt := range tests {
		result := TypeOf(tt.input)
		if result != tt.expected {
			t.Errorf("Expected type %v, got %v", tt.expected, result)
		}
	}
}

// TestLenOf verifies LenOf works for slices, arrays, and maps.
func TestLenOf(t *testing.T) {
	tests := []struct {
		input    any
		expected int
	}{
		{input: []int{1, 2, 3}, expected: 3},
		{input: [2]int{1, 2}, expected: 2},
		{input: map[string]int{"a": 1}, expected: 1},
		{input: 10, expected: -1}, // Unsupported type
	}

	for _, tt := range tests {
		result := LenOf(tt.input)
		if result != tt.expected {
			t.Errorf("Expected length %d, got %d", tt.expected, result)
		}
	}
}

// TestSizeOf validates memory size calculation for different structures.
func TestSizeOf(t *testing.T) {
	slice := []int{1, 2, 3}
	size := SizeOf(slice)
	if size == "" || size == "0 bytes" {
		t.Error("Expected non-zero size for slice")
	}

	unsupportedType := 10
	sizeUnsupported := SizeOf(unsupportedType)
	if sizeUnsupported == "" {
		t.Error("Expected size to handle unsupported type gracefully")
	}
}
