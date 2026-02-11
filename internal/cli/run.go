package cli

import (
	"fmt"
	"os"

	"github.com/hashvelani/pplnctl/internal/ci"
	"github.com/spf13/cobra"
)

var (
	runConfigFile string
	runDryRun     bool
)

var runCmd = &cobra.Command{
	Use:   "run [pipeline-name]",
	Short: "Run a pipeline",
	Long:  "Execute a pipeline with the specified name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pipelineName := args[0]

		executor := ci.NewExecutor()

		opts := ci.RunOptions{
			Name:    pipelineName,
			Config:  runConfigFile,
			DryRun:  runDryRun,
		}
		
		if err := executor.Run(opts); err != nil {
			return fmt.Errorf("failed to run pipeline: %w", err)
		}
		
		fmt.Fprintf(os.Stdout, "Pipeline '%s' executed successfully\n", pipelineName)
		return nil
	},
}

func init() {
	runCmd.Flags().StringVarP(&runConfigFile, "config", "c", "", "Path to pipeline configuration file")
	runCmd.Flags().BoolVarP(&runDryRun, "dry-run", "d", false, "Perform a dry run without executing")
}

