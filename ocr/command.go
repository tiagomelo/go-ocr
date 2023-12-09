// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ocr

import (
	"bytes"
	"fmt"
	"os/exec"
)

// sysCommand is an interface that abstracts the functions of exec.Cmd
// for executing system commands and capturing their output, facilitating
// testing by allowing mocks to be substituted in place of actual commands.
type sysCommand interface {
	Run() error
	Stdout() string
	Stderr() string
}

// sysCommandWrapper provides a concrete implementation of the sysCommand
// interface. It wraps an exec.Cmd and captures standard output and standard
// error into internal buffers.
type sysCommandWrapper struct {
	cmd    *exec.Cmd    // cmd represents the system command to be executed.
	stdout bytes.Buffer // stdout captures the standard output of the command.
	stderr bytes.Buffer // stderr captures the standard error output of the command.
}

// Run executes the system command wrapped by sysCommandWrapper. It assigns
// the standard output and error of the command to internal buffers for later retrieval.
func (sc *sysCommandWrapper) Run() error {
	sc.cmd.Stdout = &sc.stdout
	sc.cmd.Stderr = &sc.stderr
	return sc.cmd.Run()
}

// Stdout returns the captured standard output of the command as a string.
func (sc *sysCommandWrapper) Stdout() string {
	return sc.stdout.String()
}

// Stderr returns the captured standard error output of the command as a string.
func (sc *sysCommandWrapper) Stderr() string {
	return sc.stderr.String()
}

// Command is an interface that abstracts functionalities for sending text input
// to a command and receiving text output. This abstraction aids in executing
// and processing the output of system commands.
type Command interface {
	TextOutput() (string, error)
}

// command provides an implementation of the Command interface using sysCommand
// to execute system commands and process their input and output.
type command struct {
	sc sysCommand
}

// newCommand creates and returns a new instance of a command that implements
// the Command interface. It takes an exec.Cmd and wraps it with sysCommandWrapper.
func newCommand(cmd *exec.Cmd) Command {
	return &command{
		sc: &sysCommandWrapper{
			cmd: cmd,
		},
	}
}

// TextOutput executes the command and returns its output as a string.
// If the command execution fails, it returns the standard error output.
func (c *command) TextOutput() (string, error) {
	err := c.sc.Run()
	output := c.sc.Stdout()
	if err != nil {
		stderr := c.sc.Stderr()
		return output, fmt.Errorf("command execution failed: %s", stderr)
	}
	return output, nil
}
