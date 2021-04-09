package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/reaperes/kubectx-git/cmd"
)

func main() {
	op := parseArgs(os.Args[1:])
	if err := op.Run(color.Output, color.Error); err != nil {
		_, _ = fmt.Fprintf(color.Error, err.Error())

		defer os.Exit(1)
	}
}

func parseArgs(argv []string) cmd.Op {
	if len(argv) == 0 {
		fmt.Errorf("invalid command")
	}

	if len(argv) == 1 {
		v := argv[0]
		if v == "version" {
			return cmd.VersionOp{}
		}
		return cmd.UnsupportedOp{Err: fmt.Errorf("invalid command: %s", argv[0])}
	}
	return cmd.UnsupportedOp{Err: fmt.Errorf("too many arguments")}
}
