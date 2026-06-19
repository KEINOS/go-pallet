package pallet

/*
The code here were originally taken from github.com/anthonynsimon/bild/imgio
with the below license. The code has been modified to fit the needs of this
project.

---

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
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/image/bmp"
)

// Encoder encodes the provided image and writes it.
type Encoder func(io.Writer, image.Image) error

// Open loads and decodes an image from a file.
func Open(filename string) (image.Image, error) {
	cleanPath := filepath.Clean(filename)

	file, err := os.Open(cleanPath)
	if err != nil {
		return nil, errors.Wrap(err, "no such file or directory")
	}

	return decodeAndClose(file, image.Decode)
}

// JPEGEncoder returns an encoder that writes JPEG output.
func JPEGEncoder(quality int) Encoder {
	return func(w io.Writer, img image.Image) error {
		return errors.Wrap(jpeg.Encode(w, img, &jpeg.Options{Quality: quality}),
			"failed to encode image to JPEG")
	}
}

// PNGEncoder returns an encoder that writes PNG output.
func PNGEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return errors.Wrap(png.Encode(w, img),
			"failed to encode image to PNG")
	}
}

// BMPEncoder returns an encoder that writes BMP output.
func BMPEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return errors.Wrap(bmp.Encode(w, img),
			"failed to encode image to BMP")
	}
}

// Save creates a file and writes the image using the provided encoder.
func Save(filename string, img image.Image, encoder Encoder) error {
	cleanPath := filepath.Clean(filename)

	file, err := os.Create(cleanPath)
	if err != nil {
		return errors.Wrap(err, "failed to create file to save")
	}

	return encodeAndClose(file, img, encoder)
}

// imageDecoder decodes an image stream and reports its detected format.

type imageDecoder func(io.Reader) (image.Image, string, error)

func decodeAndClose(
	file io.ReadCloser,
	decoder imageDecoder,
) (image.Image, error) {
	img, _, err := decoder(file)
	closeErr := file.Close()

	if err != nil {
		return nil, errors.Wrap(err, "failed to decode image file")
	}

	if closeErr != nil {
		return nil, errors.Wrap(closeErr, "failed to close image file")
	}

	return img, nil
}

func encodeAndClose(
	file io.WriteCloser,
	img image.Image,
	encoder Encoder,
) error {
	encodeErr := encoder(file, img)
	closeErr := file.Close()

	if encodeErr != nil {
		return encodeErr
	}

	if closeErr != nil {
		return errors.Wrap(closeErr, "failed to close image file")
	}

	return nil
}
