package pallet

import (
	"encoding/json"
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

	rgbaChannelShift = 8
	redKeyShift      = 24
	greenKeyShift    = 16
	blueKeyShift     = 8
	rgbaChannelMask  = 0xff
	decimalDigits    = "0123456789"
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
	return int(u >> rgbaChannelShift)
}

// AsHistogram returns the per-channel histogram of an image.
func AsHistogram(imgRGBA *image.RGBA) Histogram {
	// Get sizes
	bounds := imgRGBA.Bounds()

	// Create new histogram
	hist := NewHistogram()
	rowBytes := bounds.Dx() * rgbaBytesPerPixel

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		offset := imgRGBA.PixOffset(bounds.Min.X, y)
		row := imgRGBA.Pix[offset : offset+rowBytes]

		for index := 0; index < len(row); index += rgbaBytesPerPixel {
			hist.R[row[index]]++
			hist.G[row[index+1]]++
			hist.B[row[index+2]]++
			hist.A[row[index+3]]++
		}
	}

	return *hist
}

// ByOccurrence returns the image colors sorted by descending occurrence.
// Colors with equal counts are sorted by ascending RGBA key.
func ByOccurrence(imgRGBA *image.RGBA) PixInfoList {
	// Get image bounds and estimate the maximum number of unique pixels.
	bounds := imgRGBA.Bounds()
	pixelCount := bounds.Dx() * bounds.Dy()

	// Count the number of times each color appears.
	pixmap := make(map[uint32]int, pixelCount)
	rowBytes := bounds.Dx() * rgbaBytesPerPixel

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		offset := imgRGBA.PixOffset(bounds.Min.X, y)
		row := imgRGBA.Pix[offset : offset+rowBytes]

		for index := 0; index < len(row); index += rgbaBytesPerPixel {
			key := uint32(row[index])<<redKeyShift |
				uint32(row[index+1])<<greenKeyShift |
				uint32(row[index+2])<<blueKeyShift |
				uint32(row[index+3])
			pixmap[key]++
		}
	}

	// Convert the map into PixInfoList.
	pixList := make(PixInfoList, len(pixmap))

	listIndex := 0

	for key, count := range pixmap {
		pixList[listIndex] = PixInfo{
			R:     int(key >> redKeyShift),
			G:     int(key>>greenKeyShift) & rgbaChannelMask,
			B:     int(key>>blueKeyShift) & rgbaChannelMask,
			A:     int(key) & rgbaChannelMask,
			Count: count,
		}

		listIndex++
	}

	// Sort by descending occurrence, then by ascending color key so ties are
	// deterministic instead of inheriting randomized map iteration order.
	sort.Slice(pixList, func(firstIndex, secondIndex int) bool {
		first := pixList[firstIndex]
		second := pixList[secondIndex]

		if first.Count != second.Count {
			return first.Count > second.Count
		}

		if first.R != second.R {
			return first.R < second.R
		}

		if first.G != second.G {
			return first.G < second.G
		}

		if first.B != second.B {
			return first.B < second.B
		}

		return first.A < second.A
	})

	return pixList
}

// ColorToString returns a color value in RRRGGGBBBAAA format.
// It is mainly used as a map key.
func ColorToString(c color.Color) string {
	red, green, blue, alpha := c.RGBA()
	key := [12]byte{}

	putThreeDigits(key[0:3], Uint32ToInt(red))
	putThreeDigits(key[3:6], Uint32ToInt(green))
	putThreeDigits(key[6:9], Uint32ToInt(blue))
	putThreeDigits(key[9:12], Uint32ToInt(alpha))

	return string(key[:])
}

func putThreeDigits(dst []byte, value int) {
	dst[0] = decimalDigits[value/100]
	dst[1] = decimalDigits[value/10%10]
	dst[2] = decimalDigits[value%10]
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
