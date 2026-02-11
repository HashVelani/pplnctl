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

