package pallet

import (
	"github.com/pkg/errors"
)

// ----------------------------------------------------------------------------
//  Type: Histogram
// ----------------------------------------------------------------------------

// Histogram holds the total occurrence of each RGBA channel.
type Histogram struct {
	R []int `json:"r"`
	G []int `json:"g"`
	B []int `json:"b"`
	A []int `json:"a"`
}

// ----------------------------------------------------------------------------
//  Constructor
// ----------------------------------------------------------------------------

// NewHistogram returns an initialized object pointer of Histogram.
func NewHistogram() *Histogram {
	shadesMax := 256 // each channel has 256 shade levels (0-255)

	// Initializes each channel with an index in the range of 0 to 255. Each
	// index represents the shade level of the color channel, and its value will
	// be the occurrence.
	return &Histogram{
		R: make([]int, shadesMax),
		G: make([]int, shadesMax),
		B: make([]int, shadesMax),
		A: make([]int, shadesMax),
	}
}

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// InJSON returns the histogram of the image in JSON string.
//
//   {
//     "r": [...],
//     "g": [...],
//     "g": [...],
//     "a": [...],
//   }
//
// Each channel contains a matrix consisting of 256 elements. The index of the
// matrix represents the shadow level, and the value represents the number of
// occurrence of that level.
func (h *Histogram) InJSON(perLine bool) (string, error) {
	if perLine {
		byteJSON, err := JSONMarshalIndent(h, "", "  ")
		if err != nil {
			return "", errors.Wrap(err, "JSON conversion failed")
		}

		return string(byteJSON), nil
	}

	byteJSON, err := JSONMarshal(h)
	if err != nil {
		return "", errors.Wrap(err, "JSON conversion failed")
	}

	return string(byteJSON), nil
}
