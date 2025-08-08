# Templ LSP Struct Test

A Go project demonstrating Templ templating with hot reload development setup using Air and Mage for task automation.

## Project Structure

```
templ-lsp-struct-test/
├── air.toml              # Air configuration for hot reload
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
├── mage.go               # Mage build automation tasks
├── main.go               # Main application entry point
├── templates/
│   ├── example.templ     # Templ template file
│   └── example_templ.go  # Generated Go code from template
└── .gitignore           # Git ignore patterns
```

## Prerequisites

- Go 1.24.5 or later
- [Templ](https://templ.guide/) - Go templating language
- [Mage](https://magefile.org/) - Build automation tool
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps (optional, for development)

## Installation

1. Fork the repository:

> <https://github.com/dontWatchMeCode/templ-issue-template>

2. Install dependencies:

```bash
go mod tidy
```

3. Install dev tools:

```bash
# Install Templ
go install github.com/a-h/templ/cmd/templ@latest

# Install Mage (optional, for better commands)
go install github.com/magefile/mage@latest

# Install Air (optional, for development)
go install github.com/cosmtrek/air@latest
```

## Mage Commands

This project uses Mage for build automation. Here are the available commands:

### `mage` or `mage find`

Interactive command selector using fzf (if available). If fzf is not installed, it falls back to listing all available commands.

### `mage list`

List all available Mage targets.

### `mage build`

- Removes existing generated Templ files
- Generates new Templ files with verbose output
- Builds the Go application to `.bin/main`

### `mage dev`

Starts development mode with:

- Templ file watching and auto-generation
- Air hot reload for Go files
- Graceful shutdown on interrupt

### `mage tidyAndUpdate`

Comprehensive update command that:

- Updates Templ to the latest version
- Updates all Go dependencies
- Runs `go mod tidy`
- Formats Go code with `gofmt`
- Formats Templ files with `templ fmt`
- Builds the project

### `mage removeTemplFiles`

Removes all generated `*_templ.go` files from the project.

## Development

### Quick Start

```bash
# Start development mode with hot reload
mage dev
```

### Manual Development Workflow

```bash
# Generate Templ files
templ generate

# Build the application
go build -o .bin/main main.go

# Run the application
./.bin/main
```

## Configuration

### Air Configuration

The `air.toml` file is configured to:

- Watch `.go`, `.tpl`, `.tmpl`, `.templ`, and `.html` files
- Exclude generated `*_templ.go` files from watching
- Use a 250ms delay for file changes
- Build to `.air/main` directory

### Templ Configuration

Templ files are located in the `templates/` directory and are automatically generated to Go files with the `_templ.go` suffix.
