package pallet

import (
	"image"
	"image/color"
	"math"

	"github.com/pkg/errors"
)

// Diff returns an image.RGBA object whose pixels are the absolute difference values between two images.
// The two input images must have the same bounds.
//
//nolint:varnamelen // Allow short variable names for readability.
func Diff(img1, img2 *image.RGBA) (*image.RGBA, error) {
	width := img1.Rect.Dx()
	height := img1.Rect.Dy()

	if width != img2.Rect.Dx() || height != img2.Rect.Dy() {
		return nil, errors.New("the two bounds should be the same size")
	}

	imgOut := image.NewRGBA(img1.Rect)

	getAbsDiff := func(a, b uint32) uint8 {
		return uint8(math.Abs(float64(a) - float64(b)))
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get color
			red1, green1, blue1, alpha1 := img1.At(x, y).RGBA()
			red2, green2, blue2, alpha2 := img2.At(x, y).RGBA()

			// Get absolute diff
			diffColor := color.RGBA{
				R: getAbsDiff(red1, red2),
				G: getAbsDiff(green1, green2),
				B: getAbsDiff(blue1, blue2),
				A: getAbsDiff(alpha1, alpha2),
			}

			// Set diff
			imgOut.SetRGBA(x, y, diffColor)
		}
	}

	return imgOut, nil
}
