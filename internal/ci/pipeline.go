// Package ci provides CI/CD pipeline functionality.
package ci

import (
	"fmt"
	"os"
)

// RunOptions contains options for running a pipeline.
type RunOptions struct {
	Name   string // Name of the pipeline to run
	Config string // Path to configuration file (optional)
	DryRun bool   // If true, perform a dry run without executing
}

// Status represents the status of a pipeline.
type Status string

const (
	// StatusUnknown indicates the pipeline status is unknown.
	StatusUnknown Status = "unknown"
	// StatusRunning indicates the pipeline is currently running.
	StatusRunning Status = "running"
	// StatusSuccess indicates the pipeline completed successfully.
	StatusSuccess Status = "success"
	// StatusFailed indicates the pipeline execution failed.
	StatusFailed Status = "failed"
	// StatusPending indicates the pipeline is waiting to be executed.
	StatusPending Status = "pending"
)

// Executor handles pipeline execution.
type Executor struct {
	configPath string
}

// NewExecutor creates a new pipeline executor.
func NewExecutor() *Executor {
	return &Executor{}
}

// Run executes a pipeline with the given options.
// It returns an error if the pipeline execution fails.
func (e *Executor) Run(opts RunOptions) error {
	if opts.DryRun {
		return e.dryRun(opts)
	}

	return e.execute(opts)
}

func (e *Executor) dryRun(opts RunOptions) error {
	fmt.Fprintf(os.Stdout, "[DRY RUN] Would execute pipeline: %s\n", opts.Name)
	if opts.Config != "" {
		fmt.Fprintf(os.Stdout, "[DRY RUN] Using config: %s\n", opts.Config)
	}
	return nil
}

func (e *Executor) execute(opts RunOptions) error {
	if opts.Config != "" {
		if _, err := os.Stat(opts.Config); os.IsNotExist(err) {
			return fmt.Errorf("config file not found: %s", opts.Config)
		}
	}

	fmt.Fprintf(os.Stdout, "Executing pipeline: %s\n", opts.Name)
	return nil
}

// Manager handles pipeline management operations.
type Manager struct {
	pipelines map[string]Status
}

// NewManager creates a new pipeline manager.
func NewManager() *Manager {
	return &Manager{
		pipelines: make(map[string]Status),
	}
}

// GetStatus returns the status of a pipeline.
// It returns an error if the pipeline is not found.
func (m *Manager) GetStatus(name string) (Status, error) {
	status, exists := m.pipelines[name]
	if !exists {
		return StatusUnknown, fmt.Errorf("pipeline not found: %s", name)
	}
	return status, nil
}

// List returns all available pipelines.
func (m *Manager) List() ([]string, error) {
	pipelines := make([]string, 0, len(m.pipelines))
	for name := range m.pipelines {
		pipelines = append(pipelines, name)
	}

	if len(pipelines) == 0 {
		pipelines = []string{"example-pipeline"}
	}

	return pipelines, nil
}

