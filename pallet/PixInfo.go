package pallet

import (
	"fmt"
)

// ----------------------------------------------------------------------------
//  Type: PixInfo
// ----------------------------------------------------------------------------

// PixInfo holds the color (RGBA) and it's number of occurrences.
type PixInfo struct {
	R     int `json:"r"`     // R is the red channel value
	G     int `json:"g"`     // G is the green channel value
	B     int `json:"b"`     // B is the blue channel value
	A     int `json:"a"`     // A is the alpha channel value
	Count int `json:"count"` // Count is the number of occurrences
}

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// GetKey returns the RGBA values in RRRGGGBBBAAA format string for ID key.
func (p PixInfo) GetKey() string {
	return fmt.Sprintf("%03v%03v%03v%03v", p.R, p.G, p.B, p.A)
}

// MarshalJSON is an implementation of Marshaler which returns the elements in
// a single line.
func (p PixInfo) MarshalJSON() ([]byte, error) {
	d := fmt.Sprintf(
		`{"r":%v, "g":%v, "b":%v, "a":%v, "count":%v}`,
		p.R, p.G, p.B, p.A, p.Count,
	)

	return []byte(d), nil
}
