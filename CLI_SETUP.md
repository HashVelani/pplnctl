# CLI Setup - Step by Step

## Step 1: Basic CLI Structure

### Files needed:

1. **go.mod** - Module definition with cobra dependency
2. **main.go** - Entry point
3. **internal/cli/root.go** - Root command
4. **internal/cli/version.go** - Version command

## Step 1: Create go.mod

```go
module github.com/hashvelani/pplnctl

go 1.21

require github.com/spf13/cobra v1.8.0

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
```

## Step 2: Create main.go

```go
package main

import (
	"os"

	"github.com/hashvelani/pplnctl/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
```

## Step 3: Create internal/cli/root.go

```go
// Package cli provides the command-line interface for the pipeline control tool.
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pplnctl",
	Short: "Pipeline control CLI",
	Long:  "A CLI tool for managing and controlling pipelines",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Fprintf(os.Stderr, "Error displaying help: %v\n", err)
		}
	},
}

// Execute runs the root command and returns any error that occurred.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd)
	// Add more commands here as you build them
}
```

## Step 4: Create internal/cli/version.go

```go
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Print the version number of pplnctl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pplnctl v0.1.0")
	},
}
```

## Commands to run:

```bash
# Initialize go module (if starting fresh)
go mod init github.com/hashvelani/pplnctl

# Get dependencies
go mod tidy

# Build
go build -o bin/pplnctl .

# Test
./bin/pplnctl --help
./bin/pplnctl version
```

## Next Steps:

After CLI is working, you can add:
- `run` command
- `list` command  
- `status` command
- Pipeline functionality
- Language handlers

