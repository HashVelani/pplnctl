// Package java provides Java-specific pipeline handlers.
package java

import (
	"fmt"
)

// Handler handles Java-specific pipeline operations.
type Handler struct {
	version string
}

// NewHandler creates a new Java pipeline handler.
func NewHandler() *Handler {
	return &Handler{
		version: "21",
	}
}

// Build executes a Java build pipeline.
func (h *Handler) Build(projectPath string) error {
	fmt.Printf("Building Java project at: %s\n", projectPath)
	return nil
}

// Test executes Java tests.
func (h *Handler) Test(projectPath string) error {
	fmt.Printf("Running Java tests at: %s\n", projectPath)
	return nil
}

// SetVersion sets the Java version to use.
func (h *Handler) SetVersion(version string) {
	h.version = version
}

// GetVersion returns the current Java version.
func (h *Handler) GetVersion() string {
	return h.version
}

