# Golang Learning Project - Claude Code Context

## Project Overview

Repository for learning Go (Golang) using the book "Learning Go" by Jon Bodner.

**Module:** `github.com/ArbinBhasimaOfficial/Golang`  
**Go Version:** 1.26.3

## Project Structure

```
Golang/
├── go.mod           # Module initialized at root level
├── README.md        # Project documentation
├── CLAUDE.md        # This file - learning progress
├── Notes.txt        # Table of contents
├── Chapter1/         # COMPLETED
│   ├── Notes.txt
│   └── hello-world/
│       ├── main.go
│       └── Makefile
└── Chapter2/         # IN PROGRESS
    ├── Notes.txt
    └── Makefile
```

## Progress

### ✅ Chapter 1: Setting Up Your Go Environment (COMPLETED)

- Installing Go (Windows via Chocolatey)
- Go tooling (go version, build, fmt, mod, test, vet, etc.)
- Writing first Go program (hello-world)
- `fmt.Printf` with format verbs (`%s`)
- Go fmt and the semicolon insertion rule
- go vet for detecting bugs
- Makefiles for automation
- Go compatibility promise
- **Exercises completed:**
  - Go Playground
  - Makefile clean target
  - Formatting experiments

### 🔄 Chapter 2: Predeclared Types and Declarations (IN PROGRESS)

- Started: "The Predeclared Types" section
- Covered: Zero Value concept
- Remaining: Literals, more on predeclared types

## Table of Contents (From Notes.txt)

1. Setting Up Your Go Environment ✅
2. Predeclared Types and Declarations 🔄
3. Composite Types
4. Blocks, Shadows and Control Structures
5. Functions
6. Pointers
7. Types, Methods, and Interfaces
8. Generics
9. Errors
10. Modules, Package, and Imports
11. Go Tooling
12. Concurrency in Go
13. The Standard Library
14. The Context
15. Writing Tests
16. Here Be Dragons: Reflect, Unsafe, and Cgo

## Key Preferences

- Using **Cursor** (VS Code derivative) as IDE
- Using **Chocolatey** for Windows package management
- Using **Make** (installed via `choco install make`) for build automation
- Module initialized at root (`github.com/ArbinBhasimaOfficial/Golang`) so subdirectories don't need separate `go mod init`

## Notes

- User prefers clean, simple explanations
- Running Windows 11 Pro (based on environment)
- Completed exercises independently
- Uses `go clean -cache` for cleaning build artifacts