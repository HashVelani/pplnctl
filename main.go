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

