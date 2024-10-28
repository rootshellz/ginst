# ginst

**ginst** is a collection of Go utilities designed to help with common
instrumentation and debugging tasks. It includes functions for JSON serialization,
file writing, reflection-based type inspection, and structured debugging prints,
making it ideal for developers needing quick, reusable utilities for Go applications.

## Table of Contents

- [ginst](#ginst)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
    - [File Utilities](#file-utilities)
    - [JSON Utilities](#json-utilities)
    - [Print Utilities](#print-utilities)
    - [Reflection Utilities](#reflection-utilities)
    - [Spew Utilities](#spew-utilities)
  - [Contributing](#contributing)
  - [License](#license)

---

## Installation

To use `ginst` in your Go project, install it via `go get`:

```bash
go get github.com/rootshellz/ginst
```

## Usage

### File Utilities

The `WriteToFile` function helps write various data types
(byte slices, strings, structs, or maps) to files.

```go
package main

import "github.com/rootshellz/ginst"

func main() {
    // Write a string to a file
    ginst.WriteToFile("Hello, World!", "/tmp/hello.txt")

    // Write a JSON object to a file
    data := map[string]string{"greeting": "Hello", "target": "World"}
    ginst.WriteToFile(data, "/tmp/greeting.json")
}
```

### JSON Utilities

The `JSON` function provides an easy way to serialize data into indented JSON.

```go
package main

import (
    "fmt"
    "github.com/rootshellz/ginst"
)

func main() {
    data := map[string]string{"key": "value"}
    jsonData, err := ginst.JSON(data)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(string(jsonData))
}
```

### Print Utilities

**DebugPrint** and **DebugPrintWithMessage** allow structured debugging output,
with separators for clarity. **PrintMapAsJSON** provides a quick way to print
maps in JSON format.

```go
package main

import "github.com/rootshellz/ginst"

func main() {
    ginst.DebugPrint("Debug message")

    ginst.DebugPrintWithMessage("Status Update", "All systems go")

    // Print a map as JSON
    data := map[string]int{"a": 1, "b": 2}
    ginst.PrintMapAsJSON(data)
}
```

### Reflection Utilities

Use `TypeOf`, `LenOf`, and `SizeOf` for type inspection and size calculations
in memory. These functions utilize reflection to analyze the underlying
structure of variables.

```go
package main

import (
    "fmt"
    "github.com/rootshellz/ginst"
)

func main() {
    fmt.Println(ginst.TypeOf(123))                 // Output: int
    fmt.Println(ginst.LenOf([]int{1, 2, 3}))       // Output: 3
    fmt.Println(ginst.SizeOf(map[string]int{"a": 1})) // Outputs memory size in bytes
}
```

### Spew Utilities

The `Spew` function provides a detailed console output of any data, while
`SpewToFile` saves the detailed dump to a specified file path. This is useful
for complex nested structures or deep debugging.

```go
package main

import "github.com/rootshellz/ginst"

func main() {
    // Print complex data structure to console
    ginst.Spew([][]int{{1, 2, 3}, {4, 5, 6}})

    // Write data structure to file
    data := map[string]int{"a": 1, "b": 2}
    ginst.SpewToFile(data, "/tmp/data_dump.txt")
}
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add a new feature'`).
4. Push the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

Ensure tests pass locally (`go test ./...`) before submitting a PR.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE)
file for details.
