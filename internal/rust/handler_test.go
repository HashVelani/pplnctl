package rust

import (
	"testing"
)

func TestNewHandler(t *testing.T) {
	handler := NewHandler()
	if handler == nil {
		t.Fatal("NewHandler() returned nil")
	}
	if handler.GetToolchain() != "stable" {
		t.Errorf("NewHandler() toolchain = %v, want 'stable'", handler.GetToolchain())
	}
}

func TestHandler_SetToolchain(t *testing.T) {
	handler := NewHandler()
	handler.SetToolchain("nightly")
	if handler.GetToolchain() != "nightly" {
		t.Errorf("SetToolchain() toolchain = %v, want 'nightly'", handler.GetToolchain())
	}
}

func TestHandler_Build(t *testing.T) {
	handler := NewHandler()
	err := handler.Build("/test/path")
	if err != nil {
		t.Fatalf("Build() should not return error, got: %v", err)
	}
}

func TestHandler_Test(t *testing.T) {
	handler := NewHandler()
	err := handler.Test("/test/path")
	if err != nil {
		t.Fatalf("Test() should not return error, got: %v", err)
	}
}

func TestHandler_Clippy(t *testing.T) {
	handler := NewHandler()
	err := handler.Clippy("/test/path")
	if err != nil {
		t.Fatalf("Clippy() should not return error, got: %v", err)
	}
}

