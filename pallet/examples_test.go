package pallet_test

import (
	"fmt"
	"log"

	"github.com/KEINOS/go-pallet/pallet"
)

func ExampleAsHistogram() {
	// 2x2 pixel image with each RGBA color of 1-pixel
	pathFile := "../testdata/r1g1b1a1.png"

	imgRGBA, err := pallet.Load(pathFile)
	if err != nil {
		log.Fatal(err)
	}

	hist := pallet.AsHistogram(imgRGBA)

	// Print the occurrences of each color channel's shade level.
	//   <channel>[<shade level>] = <occurrence>
	// If a red pixel with max-opacity (R,G,B,A=255,0,0,255) appeared twice in
	// an image then it will be:
	//   r[255]=2, g[0]=2, b[0]=2, a[255]=2
	fmt.Printf("r[0]=%v, r[255]=%v\n", hist.R[0], hist.R[255])
	fmt.Printf("g[0]=%v, g[255]=%v\n", hist.G[0], hist.G[255])
	fmt.Printf("b[0]=%v, b[255]=%v\n", hist.B[0], hist.B[255])
	fmt.Printf("a[0]=%v, a[255]=%v\n", hist.A[0], hist.A[255])

	// Output:
	// r[0]=3, r[255]=1
	// g[0]=3, g[255]=1
	// b[0]=3, b[255]=1
	// a[0]=1, a[255]=3
}

func ExampleByOccurrence() {
	pathFile := "../testdata/gopher.png"

	imgRGBA, err := pallet.Load(pathFile)
	if err != nil {
		log.Fatal(err)
	}

	pixInfoList := pallet.ByOccurrence(imgRGBA)

	// Print the first 2 most used colors
	fmt.Println(pixInfoList[0:2])

	// Output:
	// [{0 0 0 0 46618} {208 182 152 255 32505}]
}

func ExamplePixInfo_GetKey() {
	pixInfo := pallet.PixInfo{
		R: 12, // Red --> 012
		G: 34, // Green --> 034
		B: 56, // Blue --> 056
		A: 0,  // Alpha --> 255
	}

	key := pixInfo.GetKey()

	fmt.Println(key)

	// Note that each RGBA values are filled with zero
	// Output: 012034056000
}

func ExamplePixInfoList_InJSON_element_per_line() {
	pathFileImg := "../testdata/r1g2b4a2.png"

	// Load image
	imgRGBA, err := pallet.Load(pathFileImg)
	if err != nil {
		log.Fatal(err)
	}

	// Count by occurrence
	pixInfoList := pallet.ByOccurrence(imgRGBA)

	// Print in JSON (each element per line)
	outputPerLine := true

	result, err := pixInfoList.InJSON(outputPerLine)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	// Output:
	// [
	// {"r":0,"g":0,"b":0,"a":0,"count":12},
	// {"r":255,"g":255,"b":255,"a":255,"count":6},
	// {"r":0,"g":0,"b":255,"a":255,"count":4},
	// {"r":0,"g":255,"b":0,"a":255,"count":2},
	// {"r":255,"g":0,"b":0,"a":255,"count":1}
	// ]
}

func ExamplePixInfoList_InJSON_single_line() {
	pathFileImg := "../testdata/r1g2b4a2.png"

	// Load image
	imgRGBA, err := pallet.Load(pathFileImg)
	if err != nil {
		log.Fatal(err)
	}

	// Count by occurrence
	pixInfoList := pallet.ByOccurrence(imgRGBA)

	// Print in JSON as a single line
	outputPerLine := false

	result, err := pixInfoList.InJSON(outputPerLine)
	if err != nil {
		log.Fatal(err)
	}

	// Print-out in fixed width
	width := 70
	for i, r := range result {
		if i%width == 0 {
			fmt.Println()
		}

		fmt.Print(string(r))
	}

	// Output:
	// [{"r":0,"g":0,"b":0,"a":0,"count":12},{"r":255,"g":255,"b":255,"a":255
	// ,"count":6},{"r":0,"g":0,"b":255,"a":255,"count":4},{"r":0,"g":255,"b"
	// :0,"a":255,"count":2},{"r":255,"g":0,"b":0,"a":255,"count":1}]
}

func ExamplePixKey() {
	pix := pallet.PixKey("123456789255")

	fmt.Println("Red:", pix.GetRed())
	fmt.Println("Green:", pix.GetGreen())
	fmt.Println("Blue:", pix.GetBlue())
	fmt.Println("Alpha:", pix.GetAlpha())

	// Output:
	// Red: 123
	// Green: 456
	// Blue: 789
	// Alpha: 255
}

func ExamplePixKey_direct() {
	r := pallet.PixKey("123456789255").GetRed()
	g := pallet.PixKey("123456789255").GetGreen()
	b := pallet.PixKey("123456789255").GetBlue()
	a := pallet.PixKey("123456789255").GetAlpha()

	fmt.Println("Red:", r)
	fmt.Println("Green:", g)
	fmt.Println("Blue:", b)
	fmt.Println("Alpha:", a)

	// Output:
	// Red: 123
	// Green: 456
	// Blue: 789
	// Alpha: 255
}

func ExamplePixKey_NewPixInfo() {
	pixKey := pallet.PixKey("123456789255")

	// Create new PixInfo object from pixKey
	count := 0
	pixInfo := pixKey.NewPixInfo(count)

	fmt.Println(pixInfo)
	// Output: {123 456 789 255 0}
}
