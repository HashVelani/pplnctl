package main

import (
	"os"

	cli "github.com/hashvelani/pplnctl/internal/cmd"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}

