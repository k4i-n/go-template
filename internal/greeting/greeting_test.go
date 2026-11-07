package greeting

import (
	"errors"
	"testing"
)

func TestGreet_ByName(t *testing.T) {
	got, err := Greet("Khai")
	if err != nil {
		t.Fatalf("Greet: %v", err)
	}
	if got != "Hello, Khai!" {
		t.Errorf("got %q, want %q", got, "Hello, Khai!")
	}
}

func TestGreet_StripsWhitespace(t *testing.T) {
	got, err := Greet("  Khai ")
	if err != nil {
		t.Fatalf("Greet: %v", err)
	}
	if got != "Hello, Khai!" {
		t.Errorf("got %q, want %q", got, "Hello, Khai!")
	}
}

func TestGreet_RejectsEmptyName(t *testing.T) {
	for _, name := range []string{"", "   "} {
		if _, err := Greet(name); !errors.Is(err, ErrEmptyName) {
			t.Errorf("Greet(%q): got %v, want ErrEmptyName", name, err)
		}
	}
}
