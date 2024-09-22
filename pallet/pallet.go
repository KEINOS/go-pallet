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
	MaxUint8 = 256
	MaxInt32 = 2147483647
)

// ----------------------------------------------------------------------------
//  Global Variables
// ----------------------------------------------------------------------------

// JSONMarshal is a copy of json.Marshal() to ease mock during test.
// Temporary replace the function to mock its behavior.
var JSONMarshal = json.Marshal

// JSONMarshalIndent is a copy of json.MarshalIndent() to ease mock during test.
// Temporary replace the function to mock its behavior.
var JSONMarshalIndent = json.MarshalIndent

// ----------------------------------------------------------------------------
//  Functions
// ----------------------------------------------------------------------------

// Uint32ToInt converts uint32 to int in the range of 0 to MaxUint8 (0-255).
func Uint32ToInt(u uint32) int {
	i := int(u)

	return i / MaxUint8
}

// AsHistogram returns a Histogram object from an image.
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

// ByOccurrence returns PixInfoList which is a slice of PixInfo sorted by occurrence of color.
func ByOccurrence(imgRGBA *image.RGBA) PixInfoList {
	// Get sizes
	bounds := imgRGBA.Bounds()
	lenImg := bounds.Dx() * bounds.Dx()

	// Count up occurrence
	pixmap := make(map[string]int, lenImg)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			key := ColorToString(imgRGBA.At(x, y))

			pixmap[key]++
		}
	}

	// Re-map to PixInfoList
	pixList := make(PixInfoList, len(pixmap))

	i := 0

	for k, v := range pixmap {
		pixList[i] = PixKey(k).NewPixInfo(v)

		i++
	}

	// Sort PixInfos by occurrence
	sort.Stable(sort.Reverse(pixList))

	return pixList
}

// ColorToString returns color.RGBA object's RGBA value as a RRRGGGBBBAAA formatted string.
// Mostly used for the key of a map.
func ColorToString(c color.Color) string {
	r, g, b, a := c.RGBA()

	return fmt.Sprintf("%03v%03v%03v%03v",
		Uint32ToInt(r), Uint32ToInt(g), Uint32ToInt(b), Uint32ToInt(a))
}

// Load returns the image.RGBA object pointer read image from pathFileImg.
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
