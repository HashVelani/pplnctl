package cli

import (
	"fmt"
	"os"

	"github.com/hashvelani/pplnctl/internal/pipeline"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all pipelines",
	Long:  "List all available pipelines",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := pipeline.NewManager()
		pipelines, err := manager.List()
		if err != nil {
			return fmt.Errorf("failed to list pipelines: %w", err)
		}

		if len(pipelines) == 0 {
			fmt.Fprintln(os.Stdout, "No pipelines found")
			return nil
		}

		fmt.Fprintln(os.Stdout, "Available pipelines:")
		for _, p := range pipelines {
			fmt.Fprintf(os.Stdout, "  - %s\n", p)
		}

		return nil
	},
}
