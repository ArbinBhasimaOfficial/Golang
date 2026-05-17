# Golang Learning Project

A comprehensive Go (Golang) learning repository covering fundamental to advanced concepts.

---

## Prerequisites

- **Go 1.26.3** or higher installed
- **Make** (optional, for using Makefiles)

---

## Project Structure

```text
Golang/
├── go.mod                     # Go module definition
├── README.md                  # This file
├── CLAUDE.md                  # Claude Code context
├── Notes.txt                  # Table of contents
├── Chapter1/                  # ✅ COMPLETED - Hello World & Basics
│   ├── Notes.txt
│   └── hello-world/
│       ├── main.go            # First Go program
│       └── Makefile           # Build automation
└── Chapter2/                  # ✅ COMPLETED - Predeclared Types & Declarations
    ├── Notes.txt
    ├── Makefile
    ├── Exercise1/             # Explicit Type Conversion
    ├── Exercise2/             # Untyped Constants
    └── Exercise3/             # Integer Overflow
```

---

# Running the Projects

## Using Make (if installed)

```bash
make -C Chapter1/hello-world       # Build hello-world
make -C Chapter1/hello-world run   # Run hello-world
make -C Chapter1/hello-world clean # Clean build files
```

---

## Using Go Commands Directly

### Chapter 1 - Hello World

```bash
cd Chapter1/hello-world
go run main.go      # Run the program
go build            # Build executable
go clean -cache     # Clean cached files
go fmt              # Format code
go vet              # Run code analysis
```

### Chapter 2 - Predeclared Types

```bash
cd Chapter2
go run ./Exercise1   # Run Exercise 1
go test ./...        # Run all chapter tests
```

---

# Topics Covered

| Chapter | Status | Topics |
|---------|--------|--------|
| Chapter 1 | ✅ Complete | `fmt.Printf`, format verbs, Go tooling, Makefiles, `go fmt`, `go vet` |
| Chapter 2 | ✅ Complete | Predeclared types, Zero values, Literals, Variables (`var` vs `:=`), Constants, Naming idioms |
| Chapter 3+ | 📚 Upcoming | Composite types (Arrays, Slices, Maps), Functions, Pointers, Interfaces, Concurrency |

---

# Detailed Chapter Notes

## Chapter 2: Predeclared Types and Declarations

> **The overriding principle when writing idiomatic Go is:**
> **Write your programs in a way that makes your intentions clear.**

---

## 1. Predeclared Types & Concepts

### The Zero Value

Go automatically assigns a default value to variables declared without an explicit value.

Examples:

- `0` for numbers
- `false` for booleans
- `""` for strings

This eliminates uninitialized variable bugs common in C/C++.

---

### Literals

Go supports:

- Integer literals
- Floating-point literals
- Rune literals
- String literals
- Imaginary literals

Literals are fundamentally **untyped**.

#### Readability Tip

Use underscores to group large numbers:

```go
1_000_000
0b1111_0000
```

Invalid examples:

```go
_123
123_
1__000
```

---

### Strings

#### Interpreted Strings

Use double quotes for standard strings:

```go
"Hello\nWorld"
```

Escape sequences like `\n` and `\"` are interpreted.

#### Raw Strings

Use backticks for multiline or literal strings:

```go
`C:\Users\Name\Desktop`
```

---

## 2. Type Categories

### Booleans

Represents `true` or `false`.

Zero value:

```go
false
```

---

### Integers

Go provides signed and unsigned integer types:

```go
int8 int16 int32 int64
uint8 uint16 uint32 uint64
```

Aliases:

```go
byte == uint8
rune == int32
```

`int` and `uint` adapt to system architecture (32-bit or 64-bit).

### Rule of Thumb

Use **`int` by default** unless working with:

- Binary files
- Network protocols
- Generic functions requiring explicit sizing

Go **does not** allow mixing `int` and `int64` without explicit conversion.

---

### Floating-Point Numbers

Available types:

```go
float32
float64
```

Prefer:

```go
float64
```

Why?

- Better precision
- Avoid repeated casting

Avoid direct comparison:

```go
a == b
```

Instead compare against an epsilon threshold:

```go
abs(a-b) < ε
```

---

### Complex Numbers

Go supports first-class complex types:

```go
complex64
complex128
```

Mostly useful for scientific computing.

---

### Strings & Runes

A `rune` is a Unicode code point:

```go
'a'
```

Strings use double quotes:

```go
"a"
```

---

## 3. Assignments & Best Practices

### Explicit Type Conversion

Go never performs automatic numeric conversion.

Example:

```go
var i int = 20
var f float64

f = float64(i)
```

---

### `var` vs `:=`

Use `var` when:

- Initializing to zero value
- Specifying explicit type

Example:

```go
var x int
var y byte = 20
```

Use `:=` inside functions:

```go
count := 10
```

⚠️ `:=` **cannot** be used at package scope.

---

### Avoid Package-Level Mutable Variables

Global mutable variables:

- Obscure data flow
- Make testing harder
- Reduce clarity

Prefer dependency injection and local state.

---

### Constants (`const`)

Constants are compile-time values.

```go
const pi = 3.14159
```

They may be:

- Typed
- Untyped

Unused constants **do not** trigger compilation errors.

---

## 4. Naming Conventions

### Camel Case

Idiomatic Go uses:

- `camelCase`
- `PascalCase`

Avoid:

```go
snake_case
```

---

### Scope vs Name Length

Smaller scope → shorter names.

Examples:

```go
i, j   // loop indexes
f      // float
buf    // buffer
```

---

### Package Visibility

Capitalization controls visibility:

```go
ExportedFunction()   // Public
internalFunction()   // Private
```

---

# Chapter 2 Exercises

## Exercise 1: Integer to Float Assignment

Demonstrates explicit type conversion.

```go
package main

import "fmt"

func main() {
	var i int = 20
	var f float64

	f = float64(i)

	fmt.Println(i, f)
}
```

---

## Exercise 2: Untyped Constant Flexibility

Untyped constants can initialize multiple types.

```go
package main

import "fmt"

func main() {
	const value = 300

	var i int = int(value)
	var f float64 = float64(value)

	fmt.Println(i, f)
}
```

---

## Exercise 3: Max Numeric Capacity & Overflow Behavior

Demonstrates integer overflow.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Println(b, smallI, bigI)
	// Output: 0 -2147483648 0
}
```


done upto Chapter 3
make



---

# Go Development Tools

| Command | Description |
|---------|-------------|
| `go version` | Show Go version |
| `go run` | Compile and run program |
| `go build` | Compile packages |
| `go install` | Compile and install packages |
| `go fmt` | Format code |
| `go vet` | Analyze code for errors |
| `go test` | Run tests |
| `go mod init` | Initialize new module |
| `go mod tidy` | Clean up dependencies |
| `go clean` | Remove cached build files |

---

# Additional Resources

- **Official Go Documentation**  
  https://go.dev/doc/

- **Go by Example**  
  https://gobyexample.com/

- **A Tour of Go**  
  https://go.dev/tour/

---