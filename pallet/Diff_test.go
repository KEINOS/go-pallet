package pallet_test

import (
	"image"
	"image/color"
	"math/rand"
	"testing"
	"time"

	"github.com/KEINOS/go-pallet/pallet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiff_size_unmatch(t *testing.T) {
	img1 := image.NewRGBA(image.Rect(0, 0, 3, 3))
	img2 := image.NewRGBA(image.Rect(0, 0, 5, 5))

	imgDiff, err := pallet.Diff(img1, img2)

	require.Error(t, err, "it should return an error if two images does not have the same size")
	assert.Nil(t, imgDiff, "on error the returned pointer should be a nil")
}

func TestDiff_change_each_rgb(t *testing.T) {
	img1 := image.NewRGBA(image.Rect(0, 0, 3, 3))
	img2 := image.NewRGBA(image.Rect(0, 0, 3, 3))

	{
		ColorIndex := 0 // Red channel as random
		randPix, expectDiff := GetRandColorRGBA(t, ColorIndex)

		// Set pix
		img1.SetRGBA(0, 0, color.RGBA{R: 0, G: 0, B: 0, A: 255})
		img2.SetRGBA(0, 0, randPix)

		// Get diff
		imgDiff, err := pallet.Diff(img1, img2)
		require.NoError(t, err)

		actualDiff := imgDiff.RGBAAt(0, 0).R
		assert.Equal(t, expectDiff, actualDiff)
	}
	{
		ColorIndex := 1 // Green channel as random
		randPix, expectDiff := GetRandColorRGBA(t, ColorIndex)

		// Set pix
		img1.SetRGBA(0, 0, color.RGBA{R: 0, G: 0, B: 0, A: 255})
		img2.SetRGBA(0, 0, randPix)

		// Get diff
		imgDiff, err := pallet.Diff(img1, img2)
		require.NoError(t, err)

		actualDiff := imgDiff.RGBAAt(0, 0).G
		assert.Equal(t, expectDiff, actualDiff)
	}
	{
		ColorIndex := 2 // Blue channel as random
		randPix, expectDiff := GetRandColorRGBA(t, ColorIndex)

		// Set pix
		img1.SetRGBA(0, 0, color.RGBA{R: 0, G: 0, B: 0, A: 255})
		img2.SetRGBA(0, 0, randPix)
		// Get diff
		imgDiff, err := pallet.Diff(img1, img2)
		require.NoError(t, err)

		actualDiff := imgDiff.RGBAAt(0, 0).B
		assert.Equal(t, expectDiff, actualDiff)
	}
}

// ----------------------------------------------------------------------------
//  Helper Functions
// ----------------------------------------------------------------------------

// GetRandColorRGBA returns a color.RGBA object and a random value.
// The color object has a random value for one of the RGBA channels defined by
// the argument. Other channels are set to zero.
func GetRandColorRGBA(t *testing.T, indexRGBA int) (pix color.RGBA, diff uint8) {
	t.Helper()

	rand.Seed(time.Now().UnixNano())
	//nolint:gosec // In this case, it is sufficient to use the pseudorandom for testing.
	diffInt := uint8(rand.Intn(256)) // rand 0-256

	switch indexRGBA {
	case 0: // R
		return color.RGBA{R: diffInt, G: 0, B: 0, A: 255}, diffInt
	case 1: // G
		return color.RGBA{R: 0, G: diffInt, B: 0, A: 255}, diffInt
	case 2: // B
		return color.RGBA{R: 0, G: 0, B: diffInt, A: 255}, diffInt
	case 3: // A
		return color.RGBA{R: 0, G: 0, B: 0, A: diffInt}, diffInt
	}

	return color.RGBA{R: 0, G: 0, B: 0, A: 255}, 0
}
