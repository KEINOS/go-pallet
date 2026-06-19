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

// NewHistogram returns an initialized Histogram.
func NewHistogram() *Histogram {
	shadesMax := 256 // each channel has 256 shade levels (0-255)
	channelCount := 4
	shades := make([]int, shadesMax*channelCount)

	// Each channel has 256 entries. The index is the shade level, and the value
	// is the number of pixels with that shade.
	return &Histogram{
		R: shades[:shadesMax:shadesMax],
		G: shades[shadesMax : shadesMax*2 : shadesMax*2],
		B: shades[shadesMax*2 : shadesMax*3 : shadesMax*3],
		A: shades[shadesMax*3 : shadesMax*4 : shadesMax*4],
	}
}

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// InJSON returns the histogram as a JSON string.
//
//	{
//	  "r": [...],
//	  "g": [...],
//	  "b": [...],
//	  "a": [...],
//	}
//
// Each channel contains 256 entries. The index is the shade level, and the
// value is the number of pixels with that shade.
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
