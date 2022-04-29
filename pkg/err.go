package dsviz

import (
	"fmt"
	"strings"
)

type pathError struct {
	msg  string
	path []string
}

func newPathError(msg string, path ...string) pathError {
	return pathError{
		msg:  msg,
		path: path,
	}
}

func (pe pathError) Error() string {
	return fmt.Sprintf("%s at %q", pe.msg, strings.Join(pe.path, ""))
}

func (pe pathError) toError() error {
	return pe
}

func (pe pathError) appendPaths(path ...string) pathError {
	return pathError{
		msg:  pe.msg,
		path: append(pe.path, path...),
	}
}
