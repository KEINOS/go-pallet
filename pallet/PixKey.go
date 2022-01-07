package pallet

import "strconv"

// ----------------------------------------------------------------------------
//  Type: PixKey
// ----------------------------------------------------------------------------

// PixKey is a string type for RRRGGGBBBAAA formatted string.
//
// RRRGGGBBBAAA formatted string is the key of the map during count of the occurrence
// of colors in an image.
//
// See: ColorToString() at ./pallet.go as well.
type PixKey string

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// GetAlpha returns the alpha value from the RRRGGGBBBAAA format key string.
//
//   a := GetAlpha("255255255100") // --> 100
func (k PixKey) GetAlpha() int {
	i, _ := strconv.Atoi(string(k)[9:12])

	return i
}

// GetBlue returns the blue value from the RRRGGGBBBAAA format key string.
//
//   a := GetBlue("255255100255") // --> 100
func (k PixKey) GetBlue() int {
	i, _ := strconv.Atoi(string(k)[6:9])

	return i
}

// GetGreen returns the green value from the RRRGGGBBBAAA format key string.
//
//   a := GetGreen("255100255255") // --> 100
func (k PixKey) GetGreen() int {
	i, _ := strconv.Atoi(string(k)[3:6])

	return i
}

// GetRed returns the red value from the RRRGGGBBBAAA format key string.
//
//   a := GetRed("100255255255") // --> 100
func (k PixKey) GetRed() int {
	i, _ := strconv.Atoi(string(k)[0:3])

	return i
}

// NewPixInfo creates PixInfo object from PixKey.
func (k PixKey) NewPixInfo(count int) PixInfo {
	return PixInfo{
		R:     k.GetRed(),
		G:     k.GetGreen(),
		B:     k.GetBlue(),
		A:     k.GetAlpha(),
		Count: count,
	}
}
