// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/tiagomelo/go-ocr/ocr"
)

var opts struct {
	TesseractPath  string `short:"t" long:"tesseractPath" description:"Image file path"`
	ImageFilePath  string `short:"i" long:"imageFilePath" description:"Image file path" required:"true"`
	OutputFilePath string `short:"o" long:"outputFilePath" description:"Output file path (optional)"`
}

func main() {
	if _, err := flags.ParseArgs(&opts, os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	t, err := ocr.New(ocr.TesseractPath(opts.TesseractPath))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	extractedText, err := t.TextFromImageFile(opts.ImageFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if opts.OutputFilePath != "" {
		err := os.WriteFile(opts.OutputFilePath, []byte(extractedText), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("extracted text from image file %s written to %s\n", opts.ImageFilePath, opts.OutputFilePath)
		os.Exit(0)
	}
	fmt.Println("extracted text:")
	fmt.Println("")
	fmt.Println(extractedText)
}
