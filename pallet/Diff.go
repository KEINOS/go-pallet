package pallet

import (
	"image"
	"image/color"

	"github.com/pkg/errors"
)

// Diff returns an image.RGBA object whose pixels are the absolute difference values between two images.
// The two input images must have the same bounds.
//
//nolint:varnamelen // Allow short variable names for readability.
func Diff(img1, img2 *image.RGBA) (*image.RGBA, error) {
	if img1.Rect != img2.Rect {
		return nil, errors.New("the two image bounds should be the same")
	}

	imgOut := image.NewRGBA(img1.Rect)

	getAbsDiff := func(first, second uint8) uint8 {
		if first >= second {
			return first - second
		}

		return second - first
	}

	for y := img1.Rect.Min.Y; y < img1.Rect.Max.Y; y++ {
		for x := img1.Rect.Min.X; x < img1.Rect.Max.X; x++ {
			// Get color
			color1 := img1.RGBAAt(x, y)
			color2 := img2.RGBAAt(x, y)

			// Get absolute diff
			diffColor := color.RGBA{
				R: getAbsDiff(color1.R, color2.R),
				G: getAbsDiff(color1.G, color2.G),
				B: getAbsDiff(color1.B, color2.B),
				A: getAbsDiff(color1.A, color2.A),
			}

			// Set diff
			imgOut.SetRGBA(x, y, diffColor)
		}
	}

	return imgOut, nil
}
