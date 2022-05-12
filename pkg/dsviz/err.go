package dsviz

import (
	"errors"
	"fmt"
	"strings"
)

// pathError represents an error that occured while traversing a schema.
type pathError struct {
	err  error    // The error that occured
	path []string // The path to the location where the error occured
}

// newPathError creates a new path error with the given message and path.
func newPathError(msg string, path ...string) pathError {
	return pathError{
		err:  errors.New(msg),
		path: path,
	}
}

// addPathToError adds the given path to the error if it is a pathError,
// otherwise it returns the error as a pathError.
func addPathToError(err error, path ...string) error {
	if err == nil {
		return nil
	}
	if pe, ok := err.(pathError); ok {
		return pe.appendPaths(path...)
	}
	return newPathError(err.Error(), path...)
}

// Error returns the path error's message.
func (pe pathError) Error() string {
	n := len(pe.path)
	ps := make([]string, n)
	for i, p := range pe.path {
		ps[n-i-1] = p
	}
	return fmt.Sprintf("%s at %q", pe.err, strings.Join(ps, ""))
}

// Unwrap returns the underlying error.
func (pe pathError) Unwrap() error {
	return pe.err
}

// toError converts the pathError to a regular error.
func (pe pathError) toError() error {
	return pe
}

// appendPaths appends the given paths to the pathError's path.
func (pe pathError) appendPaths(path ...string) pathError {
	return pathError{
		err:  pe.err,
		path: append(pe.path, path...),
	}
}
