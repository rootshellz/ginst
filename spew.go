package ginst

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// Spew prints a detailed dump of data to the console using
// github.com/davecgh/go-spew/spew Dump.
func Spew(data any) {
	spew.Dump(data)
}

// SpewToFile writes a detailed dump of data to the specified file path, using
// github.com/davecgh/go-spew/spew Fdump.
func SpewToFile(data any, path string) {
	file, err := os.Create(path)
	if err != nil {
		DebugPrintWithMessage("Error creating file", fmt.Sprintf("%s: %s", path, err))
		return
	}
	defer file.Close()

	spew.Fdump(file, data)
	DebugPrintWithMessage("Spewed data to", path)
}
