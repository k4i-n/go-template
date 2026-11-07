package greeting

import (
	"errors"
	"strings"
)

var ErrEmptyName = errors.New("name is empty")

func Greet(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", ErrEmptyName
	}
	return "Hello, " + name + "!", nil
}
