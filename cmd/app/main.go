package main

import (
	"fmt"
	"io"
	"os"

	"github.com/nhkhai-vn/go-template/internal/greeting"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	name := "world"
	if len(args) > 0 {
		name = args[0]
	}
	text, err := greeting.Greet(name)
	if err != nil {
		return fmt.Errorf("greet: %w", err)
	}
	_, err = fmt.Fprintln(stdout, text)
	return err
}
