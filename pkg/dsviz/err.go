package dsviz

import (
	"errors"
	"fmt"
	"strings"
)

type pathError struct {
	err  error
	path []string
}

func addPathToError(err error, path ...string) error {
	if err == nil {
		return nil
	}
	if pe, ok := err.(pathError); ok {
		return pe.appendPaths(path...)
	}
	return newPathError(err.Error(), path...)
}

func newPathError(msg string, path ...string) pathError {
	return pathError{
		err:  errors.New(msg),
		path: path,
	}
}

func (pe pathError) Error() string {
	n := len(pe.path)
	ps := make([]string, n)
	for i, p := range pe.path {
		ps[n-i-1] = p
	}
	return fmt.Sprintf("%s at %q", pe.err, strings.Join(ps, ""))
}

func (pe pathError) Unwrap() error {
	return pe.err
}

func (pe pathError) toError() error {
	return pe
}

func (pe pathError) appendPaths(path ...string) pathError {
	return pathError{
		err:  pe.err,
		path: append(pe.path, path...),
	}
}
