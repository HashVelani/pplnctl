package cli

import (
	"fmt"
	"os"

	"github.com/hashvelani/pplnctl/internal/ci"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status [pipeline-name]",
	Short: "Get pipeline status",
	Long:  "Get the current status of a pipeline",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pipelineName := args[0]

		manager := ci.NewManager()
		status, err := manager.GetStatus(pipelineName)
		if err != nil {
			return fmt.Errorf("failed to get pipeline status: %w", err)
		}
		
		fmt.Fprintf(os.Stdout, "Pipeline '%s' status: %s\n", pipelineName, status)
		return nil
	},
}

