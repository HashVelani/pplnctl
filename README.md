# pplnctl

A CLI tool for managing and controlling pipelines.

## Prerequisites

- Go 1.21 or later
- Make (optional, for using Makefile)

## Building

```bash
make build
```

Or directly with Go:

```bash
go build -o bin/pplnctl .
```

## Running

```bash
./bin/pplnctl --help
```

## Testing

```bash
make test
```

Or with coverage:

```bash
make test-coverage
```

## Development

This project uses a dev container with:
- Go (latest)
- Java 21 and 25
- Node.js LTS
- Terraform
- AWS CLI
- Git
- curl

Open the project in VS Code and use "Reopen in Container" to use the dev container.

## Available Commands

- `pplnctl version` - Print version information
- `pplnctl run <pipeline-name>` - Run a pipeline
- `pplnctl status <pipeline-name>` - Get pipeline status
- `pplnctl list` - List all available pipelines

## License

MIT

