# go-ocr

[![Go Reference](https://pkg.go.dev/badge/github.com/tiagomelo/go-ocr.svg)](https://pkg.go.dev/github.com/tiagomelo/go-ocr)

![logo](doc/go-ocr.png)

A tiny [OCR](https://en.wikipedia.org/wiki/Optical_character_recognition) utility for Go, offering the ability of extracting text from an image.

It uses [Tesseract](https://en.wikipedia.org/wiki/Tesseract_(software)).

## available options

- [`TesseractPath`](ocr/options.go#L23): It allows users to specify a custom path for the Tesseract binary, overriding the default path.

## usage

Image: 

![test image](doc/image.png)

Specifying Tesseract's bin path:

```
package main

import (
	"fmt"
	"os"

	"github.com/tiagomelo/go-ocr/ocr"
)

func main() {
	const tp = "/opt/homebrew/bin/tesseract"
	t, err := ocr.New(ocr.TesseractPath(tp))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	const imagePath = "doc/image.png"
	extractedText, err := t.TextFromImageFile(imagePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("extracted text:")
	fmt.Println("")
	fmt.Println(extractedText)
}
```

Output:

```
extracted text:

Kubla Khan

In XANADU did Kubla Khan
A stately pleasure dome decree:
Where Alph, the sacred river, ran
Through caverns measureless to man
Down to a sunless sea.

So twice five miles of fertile ground
With walls and towers were girdled round:

And there were gardens bright with sinuous rills,
Where blossomed many an incense-bearing tree;
And here were forests ancient as the hills,
Enfolding sunny spots of greenery.

But Oh! that deep romantic chasm which slanted
Down the green hill athwart a cedar cover!

A savage place! as holy as enchanted

As cer beneath a waning moon was haunted

By woman wailing for her demon-lover!

And from this chasm, with ceaseless turmoil seething,
As if this earth in fast thick pants were breathing,

A might fountain momently was forced:

Amid whose swift half-intermitted burst

Huge fragments vaulted like rebounding hail,

Or chaffy grain beneath the thresher's flail:

And 'mid these dancing rocks at once and ever

It flung up momently the sacred river.
```

If you don't want to specify Tesseract's bin path,

```
	t, err := ocr.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
```

it will search for the default bin name `tesseract` in system's PATH.

## CLI example

[cmd/main.go](cmd/main.go) is a sample CLI that enables users to save the output to a file.

```
go run cmd/main.go -i doc/image.png -o image.png.text

extracted text from image file doc/image.png written to image.png.text
```

If you want to specify Tesseract's bin path,

```
go run cmd/main.go -t /opt/homebrew/bin/tesseract -i doc/image.png -o image.png.text
```

## unit tests

```
make test
```

## unit tests coverage

```
make coverage
```