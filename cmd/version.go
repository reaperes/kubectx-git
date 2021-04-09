package cmd

import (
	"fmt"
	"io"
)

type VersionOp struct{}

const version = "0.1.0"

func (_ VersionOp) Run(stdout, _ io.Writer) error {
	_, err := fmt.Fprintf(stdout, "version")
	return err
}
