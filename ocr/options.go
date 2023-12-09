// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ocr

// Option is a function type that applies a configuration to an ocr instance.
// It's used for optional configuration of the OCR process, allowing for flexible
// customization of the ocr struct's behavior.
type Option func(*ocr)

// TesseractPath returns an Option that sets the path to the Tesseract executable.
// This function allows users of the ocr package to specify a custom path for the
// Tesseract binary, overriding the default path.
//
// Parameters:
//
//	path - The file system path to the Tesseract executable.
//
// Returns:
//
//	An Option function that, when applied to an ocr instance, sets its binary path.
func TesseractPath(path string) Option {
	return func(i *ocr) {
		i.bin = path // Set the custom path for the Tesseract binary.
	}
}
