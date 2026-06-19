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

	// Each channel has 256 entries. The index is the shade level, and the value
	// is the number of pixels with that shade.
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
