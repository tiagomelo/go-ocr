// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ocr

import (
	"os/exec"

	"github.com/pkg/errors"
)

var (
	// lookPath is a variable holding the exec.LookPath function,
	// used to check for the presence of a command in the system's PATH.
	lookPath = exec.LookPath

	// _newCommand is a reference to the newCommand function, allowing it to be
	// replaced or mocked in tests.
	_newCommand = newCommand
)

// Ocr interface abstracts the OCR operations. It provides a function to extract
// text from image files.
type Ocr interface {
	// TextFromImageFile executes the OCR process on the specified image file and returns the extracted text.
	// It returns any errors encountered during the OCR operation.
	TextFromImageFile(fileName string) (string, error)
}

// ocr struct implements the Ocr interface. It encapsulates the necessary
// configuration and dependencies for performing OCR operations.
type ocr struct {
	bin        string
	newCommand func(*exec.Cmd) Command
}

// New is a constructor function for creating a new ocr instance with optional configurations.
// It returns an Ocr instance and any error encountered during its creation.
func New(options ...Option) (Ocr, error) {
	const tesseractBinName = "tesseract"
	ts := &ocr{
		newCommand: _newCommand,
	}
	for _, option := range options {
		option(ts)
	}
	if ts.bin == "" {
		ts.bin = tesseractBinName
	}
	if toolNotAvailable(ts.bin) {
		return nil, errors.Errorf("%s not found in the system's PATH", ts.bin)
	}
	return ts, nil
}

func (t *ocr) TextFromImageFile(fileName string) (string, error) {
	cmdArgs := []string{fileName, "stdout"}
	cmd := exec.Command(t.bin, cmdArgs...)
	sysCmd := t.newCommand(cmd)
	out, err := sysCmd.TextOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// toolNotAvailable checks if the specified tool is available in the system's PATH.
// It returns true if the tool is not found, otherwise false.
func toolNotAvailable(toolName string) bool {
	if _, err := lookPath(toolName); err == nil {
		return false
	}
	return true
}
