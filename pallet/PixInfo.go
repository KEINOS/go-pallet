package pallet

import (
	"fmt"
)

// ----------------------------------------------------------------------------
//  Type: PixInfo
// ----------------------------------------------------------------------------

// PixInfo holds an RGBA color and its number of occurrences.
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

// GetKey returns the RGBA values as an RRRGGGBBBAAA key string.
func (p PixInfo) GetKey() string {
	return fmt.Sprintf("%03v%03v%03v%03v", p.R, p.G, p.B, p.A)
}

// MarshalJSON implements json.Marshaler and returns a single-line object.
func (p PixInfo) MarshalJSON() ([]byte, error) {
	d := fmt.Sprintf(
		`{"r":%v, "g":%v, "b":%v, "a":%v, "count":%v}`,
		p.R, p.G, p.B, p.A, p.Count,
	)

	return []byte(d), nil
}
