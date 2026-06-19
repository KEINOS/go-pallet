package pallet_test

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/KEINOS/go-pallet/pallet"
	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleAsHistogram() {
	// 2x2 pixel image with each RGBA color of 1-pixel
	const pathFile = "../testdata/r1g1b1a1.png"

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
	const pathFile = "../testdata/gopher.png"

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

func ExampleDiff() {
	// Get image1 (3x3pix)
	const pathFileImg1 = "../testdata/rgbacmykw.png"

	imgRGBA1, err := pallet.Load(pathFileImg1)
	if err != nil {
		log.Fatal(err)
	}

	// Get image2 (3x3pix)
	const pathFileImg2 = "../testdata/rgbacmykw.png"

	imgRGBA2, err := pallet.Load(pathFileImg2)
	if err != nil {
		log.Fatal(err)
	}

	// Get the absolute diff between two images
	imgDiff, err := pallet.Diff(imgRGBA1, imgRGBA2)
	if err != nil {
		log.Fatal(err)
	}

	// It should be all zero since it's the same image
	fmt.Printf("%v", imgDiff.Pix)

	// Output:
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func ExamplePixInfo_GetKey() {
	pixInfo := pallet.PixInfo{
		R:     12, // Red   --> 012
		G:     34, // Green --> 034
		B:     56, // Blue  --> 056
		A:     0,  // Alpha --> 000
		Count: 0,  // Not used
	}

	key := pixInfo.GetKey()

	// Print the RGBA values in RRRGGGBBBAAA format. Note that each RGBA values are filled with zero
	fmt.Println(key)

	// Output: 012034056000
}

func ExamplePixInfoList_InJSON_element_per_line() {
	const pathFileImg = "../testdata/r1g2b4a2.png"

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
	const pathFileImg = "../testdata/r1g2b4a2.png"

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

/* Public Function Tests */

func TestEncode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		value   image.Image
		encoder pallet.Encoder
		format  string
	}{
		{
			format:  "png",
			encoder: pallet.PNGEncoder(),
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
		},
		{
			format:  "jpg,jpeg",
			encoder: pallet.JPEGEncoder(95),
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
		},
		{
			format:  "bmp",
			encoder: pallet.BMPEncoder(),
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
		},
	}

	for _, test := range tests {
		buf := bytes.Buffer{}
		err := test.encoder(&buf, test.value)
		require.NoError(t, err)

		_, outFormat, err := image.Decode(&buf)
		require.NoError(t, err)

		assert.Contains(t, test.format, outFormat)
	}
}

func TestOpen_fail(t *testing.T) {
	t.Parallel()

	// util.GetTempDir is similar to t.TempDir() but for compatibility with Go 1.14
	pathDirTmp, cleanup := util.GetTempDir()
	defer cleanup()

	// Un-existing path
	{
		pathDirUnknown := filepath.Join(pathDirTmp, "unknown_dir")

		_, err := pallet.Open(pathDirUnknown)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "no such file or directory")
	}

	// Path is a text file
	{
		pathFileTmp := filepath.Join(pathDirTmp, "tmp.txt")
		err := os.WriteFile(pathFileTmp, []byte("foo bar"), 0o600)
		require.NoError(t, err)

		_, err = pallet.Open(pathFileTmp)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to decode image file")
	}
}

func TestOpen_saved_image(t *testing.T) {
	t.Parallel()

	// Create image
	encPNG := pallet.PNGEncoder()
	imgRAW := &image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
			0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
			0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
		},
	}

	// Save
	pathDirTmp, cleanup := util.GetTempDir()
	defer cleanup()

	pathFileTmp := filepath.Join(pathDirTmp, "temp.png")

	err := pallet.Save(pathFileTmp, imgRAW, encPNG)
	require.NoError(t, err)

	// Assert
	expect := []byte{
		0x89, 0x50, 0x4e, 0x47, 0xd, 0xa, 0x1a, 0xa, 0x0, 0x0, 0x0, 0xd, 0x49, 0x48,
		0x44, 0x52, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x3, 0x8, 0x2, 0x0, 0x0, 0x0,
		0xd9, 0x4a, 0x22, 0xe8, 0x0, 0x0, 0x0, 0x19, 0x49, 0x44, 0x41, 0x54, 0x78,
		0x9c, 0x62, 0xf9, 0xcf, 0x0, 0x5, 0x4c, 0x10, 0xaa, 0x91, 0x81, 0x81, 0x11,
		0x2e, 0x6, 0x8, 0x0, 0x0, 0xff, 0xff, 0x2d, 0x2f, 0x2, 0x87, 0xd4, 0xef, 0xa8,
		0xdf, 0x0, 0x0, 0x0, 0x0, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
	}

	safePathFileTmp := filepath.Clean(pathFileTmp)

	actual, err := os.ReadFile(safePathFileTmp)
	require.NoError(t, err)

	assert.Equal(t, expect, actual)

	// Open - reopen file and compare to origin
	expectImg := imgRAW
	actualImg, err := pallet.Open(pathFileTmp)
	require.NoError(t, err)

	assert.EqualValues(t, expectImg, actualImg)
}

func TestSave_fail_save(t *testing.T) {
	t.Parallel()

	// Create image
	encPNG := pallet.PNGEncoder()
	imgRAW := &image.RGBA{
		Rect:   image.Rect(0, 0, 3, 3),
		Stride: 3 * 4,
		Pix: []uint8{
			0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
			0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
			0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
		},
	}

	// Save
	pathDirTmp, cleanup := util.GetTempDir()
	defer cleanup()

	err := pallet.Save(pathDirTmp, imgRAW, encPNG)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create file to save",
		"it should contain the error reason")
}
