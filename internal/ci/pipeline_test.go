package ci

import (
	"os"
	"testing"
)

func TestExecutor_Run_DryRun(t *testing.T) {
	executor := NewExecutor()
	opts := RunOptions{
		Name:   "test-pipeline",
		Config: "test-config.yaml",
		DryRun: true,
	}

	err := executor.Run(opts)
	if err != nil {
		t.Fatalf("Run() with dry-run should not return error, got: %v", err)
	}
}

func TestExecutor_Run_Execute(t *testing.T) {
	executor := NewExecutor()
	opts := RunOptions{
		Name:   "test-pipeline",
		DryRun: false,
	}

	err := executor.Run(opts)
	if err != nil {
		t.Fatalf("Run() should not return error, got: %v", err)
	}
}

func TestExecutor_Run_WithInvalidConfig(t *testing.T) {
	executor := NewExecutor()
	opts := RunOptions{
		Name:   "test-pipeline",
		Config: "/nonexistent/config.yaml",
		DryRun: false,
	}

	err := executor.Run(opts)
	if err == nil {
		t.Fatal("Run() with invalid config should return error")
	}
}

func TestExecutor_Run_WithValidConfig(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test-config-*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	executor := NewExecutor()
	opts := RunOptions{
		Name:   "test-pipeline",
		Config: tmpFile.Name(),
		DryRun: false,
	}

	err = executor.Run(opts)
	if err != nil {
		t.Fatalf("Run() with valid config should not return error, got: %v", err)
	}
}

func TestManager_GetStatus(t *testing.T) {
	manager := NewManager()

	tests := []struct {
		name          string
		pipelineName  string
		expectedError bool
	}{
		{
			name:          "existing pipeline",
			pipelineName:  "test-pipeline",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, err := manager.GetStatus(tt.pipelineName)
			if tt.expectedError && err == nil {
				t.Errorf("GetStatus() expected error but got none")
			}
			if !tt.expectedError && err != nil {
				t.Errorf("GetStatus() unexpected error: %v", err)
			}
			if !tt.expectedError && status == StatusUnknown {
				t.Errorf("GetStatus() returned unknown status")
			}
		})
	}
}

func TestManager_List(t *testing.T) {
	manager := NewManager()

	pipelines, err := manager.List()
	if err != nil {
		t.Fatalf("List() should not return error, got: %v", err)
	}

	if len(pipelines) == 0 {
		t.Error("List() should return at least one pipeline")
	}
}

func TestStatus_String(t *testing.T) {
	tests := []struct {
		status Status
		want   string
	}{
		{StatusUnknown, "unknown"},
		{StatusRunning, "running"},
		{StatusSuccess, "success"},
		{StatusFailed, "failed"},
		{StatusPending, "pending"},
	}

	for _, tt := range tests {
		t.Run(string(tt.status), func(t *testing.T) {
			if got := string(tt.status); got != tt.want {
				t.Errorf("Status.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

