package ginst

import "fmt"

// DebugPrint wraps calls to print arbitrary data in a way useful for debugging.
func DebugPrint(data any) {
	fmt.Println("--------------------------------------------------")
	fmt.Println(data)
	fmt.Println("--------------------------------------------------")
}

// DebugPrintWithMessage wraps calls to print arbitrary data in a way useful for debugging
// with an added message appearing before the data.
func DebugPrintWithMessage(message string, data any) {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%s:\n", message)
	fmt.Println(data)
	fmt.Println("--------------------------------------------------")
}

// PrintMapAsJSON prints the input map as JSON by wrapping JSON()
func PrintMapAsJSON(data any) {
	json, err := JSON(data)
	if err != nil {
		DebugPrintWithMessage("Error converting data to JSON", err)
	}
	fmt.Printf("%s\n", json)
}
