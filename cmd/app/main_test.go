package main

import (
	"errors"
	"strings"
	"testing"

	"github.com/sgl-k/go-template/internal/greeting"
)

func TestRun_PrintsGreeting(t *testing.T) {
	var out strings.Builder
	if err := run([]string{"Khai"}, &out); err != nil {
		t.Fatalf("run: %v", err)
	}
	if out.String() != "Hello, Khai!\n" {
		t.Errorf("got %q, want %q", out.String(), "Hello, Khai!\n")
	}
}

func TestRun_DefaultsToWorld(t *testing.T) {
	var out strings.Builder
	if err := run(nil, &out); err != nil {
		t.Fatalf("run: %v", err)
	}
	if out.String() != "Hello, world!\n" {
		t.Errorf("got %q, want %q", out.String(), "Hello, world!\n")
	}
}

func TestRun_RejectsEmptyName(t *testing.T) {
	var out strings.Builder
	if err := run([]string{" "}, &out); !errors.Is(err, greeting.ErrEmptyName) {
		t.Errorf("got %v, want ErrEmptyName", err)
	}
	if out.Len() != 0 {
		t.Errorf("wrote %q, want nothing", out.String())
	}
}
