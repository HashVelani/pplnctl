package cli

import (
	"testing"
)

func TestExecute(t *testing.T) {
	rootCmd.SetArgs([]string{"--help"})
	
	err := Execute()
	if err != nil {
		t.Fatalf("Execute() should not return error for help command, got: %v", err)
	}
}

func TestRootCommand(t *testing.T) {
	if rootCmd.Use != "pplnctl" {
		t.Errorf("rootCmd.Use = %v, want 'pplnctl'", rootCmd.Use)
	}
}

