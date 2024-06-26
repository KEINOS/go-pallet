package pallet

// ----------------------------------------------------------------------------
//  Type: PixInfoList
// ----------------------------------------------------------------------------

// PixInfoList is a slice of PixInfo which is sortable.
type PixInfoList []PixInfo

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// Len is an implementation of Len() for sort function. Which returns the current
// object's slice length.
func (p PixInfoList) Len() int { return len(p) }

// Less is an implementation of Less() for sort function. Which returns true if
// the current value of Count in "i" is less than "j".
func (p PixInfoList) Less(i, j int) bool { return p[i].Count < p[j].Count }

// Swap is an implementation of Swap() for sort function. It will swap the elements
// between "i" and "j".
func (p PixInfoList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// InJSON returns a JSON formatted string of the color map.
// If perLine is true then it will output each element per line.
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

	for index := range lenList {
		a := p[index]

		byteData, err := JSONMarshal(a)
		if err != nil {
			return "", err
		}

		result += string(byteData)
		if index != (lenList - 1) {
			result += ","
		}

		result += "\n"
	}

	result += "]"

	return result, nil
}
