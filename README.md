# Golang Learning Project

Notes, exercises, and runnable examples while working through [**Learning Go**](https://learning.oreilly.com/library/view/learning-go/9781098139297/) by Jon Bodner (O'Reilly).

**Module:** `github.com/ArbinBhasimaOfficial/Golang` · **Go:** 1.26.3

---

## Prerequisites

- **Go 1.26.3** or higher installed
- **Make** (optional, for using Makefiles)

---

## Project Structure

```text
Golang/
├── go.mod                     # Go module (init once at repo root)
├── README.md                  # This file
├── CLAUDE.md                  # Agent / learning context
├── Notes.txt                  # Book table of contents
├── Chapter1/                  # ✅ Setting Up Your Go Environment
│   ├── Notes.txt
│   └── hello-world/
│       ├── main.go
│       └── Makefile
├── Chapter2/                  # ✅ Predeclared Types and Declarations
│   ├── Notes.txt
│   ├── Makefile
│   ├── Exercise1/
│   ├── Exercise2/
│   └── Exercise3/
├── Chapter3/                  # ✅ Composite Types
│   ├── Notes.txt
│   ├── Excercises/            # Chapter exercises (1–3)
│   └── <topic-demos>/         # Runnable examples (slices, maps, structs, …)
└── Chapter4/                  # 🔄 Blocks, Shadows, and Control Structures
    ├── Notes.txt
    ├── shadowingVariable/
    ├── shadowingMultipleAssignment/
    ├── shadowPackageNames/
    ├── shadowTrue/
    ├── ifElseExample/
    ├── scopingVariableToIfStatement/
    ├── ifBadScope/
    ├── forStatementEg/
    ├── infiniteForLoop/
    ├── forRange/
    ├── forRangeIgnore/
    ├── iterateMap/
    └── iterateString/
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

### Chapter 3 - Composite Types

```bash
# Run any demo from the repo root (module is at root)
go run ./Chapter3/slicingSlices
go run ./Chapter3/len_cap
go run ./Chapter3/mapReadWrite
go run ./Chapter3/structDeclaration

# Chapter exercises
go run ./Chapter3/Excercises/Excercise1
go run ./Chapter3/Excercises/Excercise2
go run ./Chapter3/Excercises/Excercise3

# Demos with Makefiles (build only; run the binary or use go run above)
make -C Chapter3/slicingSlices
```

### Chapter 4 - Blocks, Shadows, and Control Structures

```bash
# Shadowing and blocks
go run ./Chapter4/shadowingVariable
go run ./Chapter4/shadowingMultipleAssignment
go run ./Chapter4/shadowPackageNames
go run ./Chapter4/shadowTrue

# if statements
go run ./Chapter4/ifElseExample
go run ./Chapter4/scopingVariableToIfStatement
go run ./Chapter4/ifBadScope

# for loops
go run ./Chapter4/forStatementEg
go run ./Chapter4/infiniteForLoop    # Ctrl-C to stop

# for range
go run ./Chapter4/forRange
go run ./Chapter4/forRangeIgnore
go run ./Chapter4/iterateMap
go run ./Chapter4/iterateString
```

---

# Topics Covered

| Chapter | Status | Topics |
|---------|--------|--------|
| 1 | ✅ Complete | Go install & tooling, `fmt.Printf`, format verbs, Makefiles, `go fmt`, `go vet` |
| 2 | ✅ Complete | Predeclared types, zero values, literals, `var` vs `:=`, constants, naming |
| 3 | ✅ Complete | Arrays, slices (`len`, `cap`, `append`, `make`, `copy`), maps, structs |
| 4 | 🔄 In progress | Blocks, variable shadowing, `if` init scope, `for` (C-style, condition-only, infinite), `break`/`continue`, `for range` (slices, maps, strings) |
| 5–16 | 📚 Upcoming | Functions, pointers, interfaces, generics, errors, concurrency, testing, … |

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


---

## Chapter 3: Composite Types

Detailed notes live in [`Chapter3/Notes.txt`](Chapter3/Notes.txt). Summary of what was covered:

| Section | Topics |
|---------|--------|
| 3.1–3.12 | Arrays vs slices, `len` / `cap`, `append`, `make`, emptying slices, slicing, `copy`, array↔slice conversion |
| 3.13 | Strings, runes, bytes, slicing strings |
| 3.14–3.20 | Maps, comma-ok idiom, delete, map-as-set |
| 3.21–3.23 | Structs, anonymous structs, comparison & conversion |

**Runnable demos** (non-exhaustive): `slicingSlices`, `len_cap`, `slice_append_storage`, `emptyingASlice`, `copy_slice`, `mapReadWrite`, `commaOkIdiom`, `structDeclaration`, `anonymousStructDeclaration`, and others under `Chapter3/`.

**Exercises** (in `Chapter3/Excercises/`):

1. Slice subslicing with multilingual greetings  
2. Fourth rune of a string containing emoji (`👩`, `👨`)  
3. `Employee` struct with three initialization styles  

---

## Chapter 4: Blocks, Shadows, and Control Structures

Detailed notes live in [`Chapter4/Notes.txt`](Chapter4/Notes.txt). Summary of what is covered so far:

| Section | Topics |
|---------|--------|
| 4.1–4.2 | Package and inner blocks, variable shadowing, `:=` in inner blocks, shadowing package names, universe block (`true`, `false`, etc.) |
| 4.3 | `if` / `else if` / `else`, no parens on conditions, init statement scope (`if n := …; n == 0`) |
| 4.4–4.8 | `for` only loop keyword: C-style, condition-only, infinite; `break` and `continue` |
| 4.9 | `for range` on slices, maps (unordered iteration), strings (runes vs bytes); `_` to ignore index or value |

**Runnable demos** (under `Chapter4/`): `shadowingVariable`, `shadowingMultipleAssignment`, `shadowPackageNames`, `shadowTrue`, `ifElseExample`, `scopingVariableToIfStatement`, `ifBadScope`, `forStatementEg`, `infiniteForLoop`, `forRange`, `forRangeIgnore`, `iterateMap`, `iterateString`.

**Still ahead in this chapter:** `switch` and remaining control-flow topics from the book.

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