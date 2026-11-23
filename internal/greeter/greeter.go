// Package greeter builds greeting messages.
package greeter

import (
	"errors"
	"strings"
)

// ErrNameTooLong is returned when a name exceeds MaxNameLen.
var ErrNameTooLong = errors.New("name too long")

// MaxNameLen is the longest name Greet accepts, in bytes.
const MaxNameLen = 64

// Greet returns a greeting for name, or for "world" when name is blank.
func Greet(name string) (string, error) {
	name = strings.TrimSpace(name)
	if len(name) > MaxNameLen {
		return "", ErrNameTooLong
	}
	if name == "" {
		name = "world"
	}
	return "Hello, " + name + "!", nil
}
