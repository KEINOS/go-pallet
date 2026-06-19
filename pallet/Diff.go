package pallet

import (
	"image"

	"github.com/pkg/errors"
)

const rgbaBytesPerPixel = 4

// Diff returns an image.RGBA object whose pixels are the absolute difference values between two images.
// The two input images must have the same bounds.
func Diff(img1, img2 *image.RGBA) (*image.RGBA, error) {
	if img1.Rect != img2.Rect {
		return nil, errors.New("the two image bounds should be the same")
	}

	imgOut := image.NewRGBA(img1.Rect)

	for y := img1.Rect.Min.Y; y < img1.Rect.Max.Y; y++ {
		firstOffset := img1.PixOffset(img1.Rect.Min.X, y)
		secondOffset := img2.PixOffset(img2.Rect.Min.X, y)
		outputOffset := imgOut.PixOffset(imgOut.Rect.Min.X, y)
		rowBytes := img1.Rect.Dx() * rgbaBytesPerPixel

		for index := range rowBytes {
			first := img1.Pix[firstOffset+index]
			second := img2.Pix[secondOffset+index]

			if first >= second {
				imgOut.Pix[outputOffset+index] = first - second
			} else {
				imgOut.Pix[outputOffset+index] = second - first
			}
		}
	}

	return imgOut, nil
}
