package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func ExecStringE(command string) error {
	return ExecArrayE(strings.Fields(command)...)
}

func ExecArrayE(args ...string) error {
	return ExecIOEArrayE(nil, nil, nil, args...)
}

func ExecIOEArrayE(ioIn io.Reader, ioOut io.Writer, ioErr io.Writer, args ...string) error {
	cmdPath, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}

	if ioIn == nil {
		ioIn = os.Stdin
	}

	if ioOut == nil {
		ioOut = os.Stdout
	}

	if ioErr == nil {
		ioErr = os.Stderr
	}

	fmt.Printf("%v\n", args)

	execCmd := &exec.Cmd{
		Path:   cmdPath,
		Args:   args,
		Env:    nil,
		Dir:    "",
		Stdin:  ioIn,
		Stdout: ioOut,
		Stderr: ioErr,
	}
	return execCmd.Run()
}
