package ginst

import (
	"fmt"
	"reflect"
	"unsafe"
)

// TypeOf returns the type of the given data using reflect.TypeOf.
func TypeOf(data any) reflect.Type {
	return reflect.TypeOf(data)
}

// LenOf returns the length of a slice, array, or the number of keys in a map.
func LenOf(data any) int {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return v.Len()
	default:
		DebugPrintWithMessage("Unsupported type: must be a slice, array, or map", TypeOf(data))
		return -1
	}
}

// SizeOf returns an estimated size of the data in memory,
// including elements of arrays, slices, and maps.
func SizeOf(data any) string {
	totalSize := CalculateSize(reflect.ValueOf(data))
	return fmt.Sprintf("%d bytes", totalSize)
}

// calculateSize is a helper function that recursively calculates
// memory size for nested structures.
func CalculateSize(v reflect.Value) uintptr {
	if !v.IsValid() {
		return 0
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		var size uintptr
		for i := 0; i < v.Len(); i++ {
			size += CalculateSize(v.Index(i))
		}
		return size + unsafe.Sizeof(v.Pointer()) // Add overhead for slice structure
	case reflect.Map:
		var size uintptr
		for _, key := range v.MapKeys() {
			size += CalculateSize(key) + CalculateSize(v.MapIndex(key))
		}
		return size + unsafe.Sizeof(v.Pointer()) // Add overhead for map structure
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return 0
		}
		return unsafe.Sizeof(v.Pointer()) + CalculateSize(v.Elem())
	default:
		return unsafe.Sizeof(v.Interface())
	}
}
