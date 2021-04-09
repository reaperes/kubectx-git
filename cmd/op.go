package cmd

import (
	"io"
)

type Op interface {
	Run(stdout, stderr io.Writer) error
}

type UnsupportedOp struct{ Err error }

func (op UnsupportedOp) Run(_, _ io.Writer) error {
	return op.Err
}
