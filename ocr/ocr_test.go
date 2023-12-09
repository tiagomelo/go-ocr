// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ocr

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name          string
		options       []Option
		mockLookPath  func(file string) (string, error)
		expectedError error
	}{
		{
			name: "happy path, no options provided",
			mockLookPath: func(file string) (string, error) {
				return "path", nil
			},
		},
		{
			name: "error, default tesseract bin not found",
			mockLookPath: func(file string) (string, error) {
				return "", errors.New("look path error")
			},
			expectedError: errors.New("tesseract not found in the system's PATH"),
		},
		{
			name:    "happy path, tesseract bin provided",
			options: []Option{TesseractPath("path/to/tesseract")},
			mockLookPath: func(file string) (string, error) {
				return "path", nil
			},
		},
		{
			name:    "error, provided tesseract bin not found",
			options: []Option{TesseractPath("/path/to/tesseract")},
			mockLookPath: func(file string) (string, error) {
				return "", errors.New("look path error")
			},
			expectedError: errors.New("/path/to/tesseract not found in the system's PATH"),
		},
	}
	for _, tc := range testCases {
		lookPath = tc.mockLookPath
		t.Run(tc.name, func(t *testing.T) {
			ts, err := New(tc.options...)
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf("expected no error, got %v", err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf("expected error to be %v, got nil", tc.expectedError)
				}
				require.NotNil(t, ts)
			}
		})
	}
}

func TestTextFromImageFile(t *testing.T) {
	testCases := []struct {
		name           string
		bin            string
		fileName       string
		mockCommand    func(cmd *exec.Cmd) Command
		mockSysCommand *mockSysCommand
		expectedOutput string
		expectedError  error
	}{
		{
			name:     "happy path",
			bin:      "mockTesseract",
			fileName: "image.png",
			mockCommand: func(cmd *exec.Cmd) Command {
				return &command{
					sc: &mockSysCommand{
						StdoutBytes: []byte("extracted text"),
					},
				}
			},
			expectedOutput: "extracted text",
		},
		{
			name:     "error",
			bin:      "mockTesseract",
			fileName: "image.png",
			mockCommand: func(cmd *exec.Cmd) Command {
				return &command{
					sc: &mockSysCommand{
						StderrBytes: []byte("random error"),
						Err:         errors.New("command failed"),
					},
				}
			},
			expectedError: errors.New("command execution failed: random error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockNewCommand := func(cmd *exec.Cmd) Command {
				return tc.mockCommand(cmd)
			}
			ts := &ocr{
				bin:        tc.bin,
				newCommand: mockNewCommand,
			}
			output, err := ts.TextFromImageFile(tc.fileName)
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf("expected no error, got %v", err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf("expected error %v, got nil", tc.expectedError)
				}
				require.Equal(t, tc.expectedOutput, output)
			}
		})
	}
}

type mockSysCommand struct {
	StdoutBytes []byte
	StderrBytes []byte
	Err         error
}

func (m *mockSysCommand) Run() error {
	return m.Err
}

func (m *mockSysCommand) Stdout() string {
	return string(m.StdoutBytes)
}

func (m *mockSysCommand) Stderr() string {
	return string(m.StderrBytes)
}
