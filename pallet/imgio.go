package pallet

/*
The coeds here were originally taken from github.com/anthonynsimon/bild/imgio.
Which holds the same MIT license as the Go-Pallet has.

MIT License

Copyright (c) 2021 Anthony Najjar Simon

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/image/bmp"
)

// Encoder encodes the provided image and writes it.
type Encoder func(io.Writer, image.Image) error

// Open loads and decodes an image from a file and returns it.
//
// Usage example:
//		// Decodes an image from a file with the given filename
//		// returns an error if something went wrong
//		img, err := Open("exampleName")
//
func Open(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "no such file or directory")
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode image file")
	}

	return img, nil
}

// JPEGEncoder returns an encoder to JPEG given the argument 'quality'.
func JPEGEncoder(quality int) Encoder {
	return func(w io.Writer, img image.Image) error {
		return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
	}
}

// PNGEncoder returns an encoder to PNG.
func PNGEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return png.Encode(w, img)
	}
}

// BMPEncoder returns an encoder to BMP.
func BMPEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return bmp.Encode(w, img)
	}
}

// Save creates a file and writes to it an image using the provided encoder.
//
// Usage example:
//		// Save an image to a file in PNG format,
//		// returns an error if something went wrong
//		err := Save("exampleName", img, imgio.JPEGEncoder(100))
//
func Save(filename string, img image.Image, encoder Encoder) error {
	// filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	return encoder(f, img)
}
