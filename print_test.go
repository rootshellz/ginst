package ginst

import (
	"bytes"
	"os"
	"testing"
)

// Capture output helper for print functions
func captureOutput(f func()) string {
	// Save the original os.Stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function that writes to stdout
	f()

	// Capture the output
	w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	os.Stdout = old // Restore original stdout

	return buf.String()
}

// Example test using captureOutput
func TestDebugPrint(t *testing.T) {
	output := captureOutput(func() {
		DebugPrint("test message")
	})

	expected := "--------------------------------------------------\n"
	expected += "test message\n"
	expected += "--------------------------------------------------\n"

	if output != expected {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expected, output)
	}
}
