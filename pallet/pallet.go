package pallet

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"sort"

	"github.com/pkg/errors"
)

// ----------------------------------------------------------------------------
//  Constants
// ----------------------------------------------------------------------------

const (
	// MaxUint8 is the divisor to convert RGBA 16-bit channel values to 8-bit.
	MaxUint8 = 256
	// MaxInt32 is the maximum value of signed 32-bit integer.
	MaxInt32 = 2147483647
)

// ----------------------------------------------------------------------------
//  Global Variables
// ----------------------------------------------------------------------------

//nolint:gochecknoglobals // Allow to ease mock during test.
var (
	// JSONMarshal is a copy of json.Marshal to ease mocking during tests.
	// Replace it temporarily when a test needs to force an error path.
	JSONMarshal = json.Marshal

	// JSONMarshalIndent is a copy of json.MarshalIndent to ease mocking during
	// tests. Replace it temporarily when a test needs to force an error path.
	JSONMarshalIndent = json.MarshalIndent
)

// ----------------------------------------------------------------------------
//  Functions
// ----------------------------------------------------------------------------

// Uint32ToInt converts a uint32 channel value to an 8-bit int.
func Uint32ToInt(u uint32) int {
	i := int(u)

	return i / MaxUint8
}

// AsHistogram returns the per-channel histogram of an image.
func AsHistogram(imgRGBA *image.RGBA) Histogram {
	// Get sizes
	bounds := imgRGBA.Bounds()

	// Create new histogram
	hist := NewHistogram()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			//nolint:varnamelen // Allow short variable name for readability
			r, g, b, a := imgRGBA.At(x, y).RGBA()

			// Count up the corresponding shade strength of each channel
			hist.R[Uint32ToInt(r)]++
			hist.G[Uint32ToInt(g)]++
			hist.B[Uint32ToInt(b)]++
			hist.A[Uint32ToInt(a)]++
		}
	}

	return *hist
}

// ByOccurrence returns the image colors sorted by occurrence.
func ByOccurrence(imgRGBA *image.RGBA) PixInfoList {
	// Get image bounds and estimate the maximum number of unique pixels.
	bounds := imgRGBA.Bounds()
	pixelCount := bounds.Dx() * bounds.Dy()

	// Count the number of times each color appears.
	pixmap := make(map[string]int, pixelCount)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			key := ColorToString(imgRGBA.At(x, y))

			pixmap[key]++
		}
	}

	// Convert the map into PixInfoList.
	pixList := make(PixInfoList, len(pixmap))

	i := 0

	for key, count := range pixmap {
		pixList[i] = PixKey(key).NewPixInfo(count)

		i++
	}

	// Sort PixInfos by occurrence
	sort.Stable(sort.Reverse(pixList))

	return pixList
}

// ColorToString returns a color value in RRRGGGBBBAAA format.
// It is mainly used as a map key.
func ColorToString(c color.Color) string {
	r, g, b, a := c.RGBA()

	return fmt.Sprintf("%03v%03v%03v%03v",
		Uint32ToInt(r), Uint32ToInt(g), Uint32ToInt(b), Uint32ToInt(a))
}

// Load reads an image file and returns it as image.RGBA.
func Load(pathFileImg string) (*image.RGBA, error) {
	img, err := Open(pathFileImg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load image")
	}

	bounds := img.Bounds()
	imgRGBA := image.NewRGBA(bounds)

	draw.Draw(imgRGBA, bounds, img, bounds.Min, draw.Src)

	return imgRGBA, nil
}
