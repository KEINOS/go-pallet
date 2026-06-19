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
	result := "[\n"
	lenList := len(p)

	var resultBuilder strings.Builder

	for index := range lenList {
		a := p[index]

		byteData, err := JSONMarshal(a)
		if err != nil {
			return "", err
		}

		resultBuilder.Write(byteData)

		if index != (lenList - 1) {
			resultBuilder.WriteString(",")
		}

		resultBuilder.WriteString("\n")
	}

	result += resultBuilder.String()

	result += "]"

	return result, nil
}
