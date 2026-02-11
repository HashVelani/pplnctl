package golang

import (
	"testing"
)

func TestNewHandler(t *testing.T) {
	handler := NewHandler()
	if handler == nil {
		t.Fatal("NewHandler() returned nil")
	}
	if handler.GetVersion() != "1.21" {
		t.Errorf("NewHandler() version = %v, want '1.21'", handler.GetVersion())
	}
}

func TestHandler_SetVersion(t *testing.T) {
	handler := NewHandler()
	handler.SetVersion("1.22")
	if handler.GetVersion() != "1.22" {
		t.Errorf("SetVersion() version = %v, want '1.22'", handler.GetVersion())
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

func TestHandler_Lint(t *testing.T) {
	handler := NewHandler()
	err := handler.Lint("/test/path")
	if err != nil {
		t.Fatalf("Lint() should not return error, got: %v", err)
	}
}

