package util

import (
	"errors"
	"runtime/debug"
)

// ErrCast is returned when a cast fails AND the program should not panic.
func ErrCast() error {
	debug.PrintStack()
	return errCast
}

var errCast = errors.New("cast error")
