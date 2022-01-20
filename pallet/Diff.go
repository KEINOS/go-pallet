package pallet

import (
	"errors"
	"image"
	"image/color"
	"math"
)

// Diff returns an image.RGBA object whose pixels are the absolute difference values between two images.
// The two input images must have the same bounds.
func Diff(img1, img2 *image.RGBA) (diff *image.RGBA, err error) {
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
			r1, g1, b1, a1 := img1.At(x, y).RGBA()
			r2, g2, b2, a2 := img2.At(x, y).RGBA()

			// Get absolute diff
			diffColor := color.RGBA{
				R: getAbsDiff(r1, r2),
				G: getAbsDiff(g1, g2),
				B: getAbsDiff(b1, b2),
				A: getAbsDiff(a1, a2),
			}

			// Set diff
			imgOut.SetRGBA(x, y, diffColor)
		}
	}

	return imgOut, nil
}
