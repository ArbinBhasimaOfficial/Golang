# Golang Learning Project

A comprehensive Go (Golang) learning repository covering fundamental to advanced concepts.

## Prerequisites

- **Go 1.26.3** or higher installed
- **Make** (optional, for using Makefiles)

## Project Structure

```
Golang/
├── go.mod                    # Go module definition
├── README.md                 # This file
├── Notes.txt                 # General notes
├── Chapter1/                 # Chapter 1: Hello World & Basics
│   ├── Notes.txt
│   └── hello-world/
│       ├── main.go           # First Go program
│       └── Makefile          # Build automation
└── Chapter2/                 # Chapter 2
    ├── Notes.txt
    └── Makefile
```

## Running the Projects

### Using Make (if installed)
```bash
make -C Chapter1/hello-world   # Build hello-world
make -C Chapter1/hello-world run  # Run hello-world
make -C Chapter1/hello-world clean # Clean build files
```

### Using Go Commands Directly

**Chapter 1 - Hello World:**
```bash
cd Chapter1/hello-world
go run main.go      # Run the program
go build           # Build executable
go clean -cache    # Clean cached files
go fmt             # Format code
go vet             # Run code analysis
```

**Chapter 2:**
```bash
cd Chapter2
go run .           # Run Chapter 2 code
go build           # Build
go test            # Run tests
```

## Topics Covered

| Chapter | Topics |
|---------|--------|
| Chapter 1 | `fmt.Printf`, basic syntax, Makefiles |
| Chapter 2 | (See Notes.txt) |

## Go Development Tools

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

## Additional Resources

- [Official Go Docs](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://go.dev/tour/)