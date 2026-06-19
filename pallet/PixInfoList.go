package pallet

import "strings"

// ----------------------------------------------------------------------------
//  Type: PixInfoList
// ----------------------------------------------------------------------------

// PixInfoList is a sortable slice of PixInfo.
type PixInfoList []PixInfo

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// Len implements sort.Interface.
func (p PixInfoList) Len() int { return len(p) }

// Less implements sort.Interface.
func (p PixInfoList) Less(i, j int) bool { return p[i].Count < p[j].Count }

// Swap implements sort.Interface.
func (p PixInfoList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// InJSON returns a JSON formatted string of the color map.
// If perLine is true, each element is written on its own line.
func (p PixInfoList) InJSON(perLine bool) (string, error) {
	if perLine {
		return p.inJSONPerLine()
	}

	return p.inJSONSingleLine()
}

// InJSONL returns the color map as JSON Lines, ending each record with a newline.
func (p PixInfoList) InJSONL() (string, error) {
	lines, err := p.inJSONLines()
	if err != nil {
		return "", err
	}

	if len(lines) == 0 {
		return "", nil
	}

	return strings.Join(lines, "\n") + "\n", nil
}

func (p PixInfoList) inJSONSingleLine() (string, error) {
	b, err := JSONMarshal(p)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// inJSONPerLine is a JSON marshaler but in a special format.
//
//	[
//	{"r":255,"g":0,"b":0,"a":255,"count":1},
//	{"r":0,"g":0,"b":255,"a":255,"count":1},
//	{"r":0,"g":255,"b":0,"a":255,"count":1},
//	{"r":0,"g":0,"b":0,"a":0,"count":1}
//	]
func (p PixInfoList) inJSONPerLine() (string, error) {
	lines, err := p.inJSONLines()
	if err != nil {
		return "", err
	}

	if len(lines) == 0 {
		return "[\n]", nil
	}

	return "[\n" + strings.Join(lines, ",\n") + "\n]", nil
}

func (p PixInfoList) inJSONLines() ([]string, error) {
	lines := make([]string, 0, len(p))

	for _, pixInfo := range p {
		byteData, err := JSONMarshal(pixInfo)
		if err != nil {
			return nil, err
		}

		lines = append(lines, string(byteData))
	}

	return lines, nil
}
