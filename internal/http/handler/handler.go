package handler

import (
	"errors"
	"strings"
)

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// Handler is a concrete implementation of StringService
type Handler struct{}

func (Handler) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (Handler) Count(s string) int {
	return len(s)
}
