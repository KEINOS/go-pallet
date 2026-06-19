package pallet_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/KEINOS/go-pallet/pallet"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestByOccurrence_ties_are_deterministic(t *testing.T) {
	t.Parallel()

	imgRGBA := image.NewRGBA(image.Rect(0, 0, 2, 2))
	imgRGBA.SetRGBA(0, 0, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	imgRGBA.SetRGBA(1, 0, color.RGBA{R: 0, G: 255, B: 0, A: 255})
	imgRGBA.SetRGBA(0, 1, color.RGBA{R: 0, G: 0, B: 255, A: 255})
	imgRGBA.SetRGBA(1, 1, color.RGBA{R: 0, G: 0, B: 0, A: 255})

	expectedKeys := []string{
		"000000000255",
		"000000255255",
		"000255000255",
		"255000000255",
	}

	for range 100 {
		pixInfoList := pallet.ByOccurrence(imgRGBA)
		actualKeys := make([]string, len(pixInfoList))

		for index, pixInfo := range pixInfoList {
			actualKeys[index] = pixInfo.GetKey()
		}

		require.Equal(t, expectedKeys, actualKeys)
	}
}

func TestPixInfoList_sort_interface(t *testing.T) {
	t.Parallel()

	pixInfoList := pallet.PixInfoList{
		{R: 0, G: 0, B: 0, A: 255, Count: 1},
		{R: 255, G: 255, B: 255, A: 255, Count: 2},
	}

	assert.Equal(t, 2, pixInfoList.Len())
	assert.True(t, pixInfoList.Less(0, 1))

	pixInfoList.Swap(0, 1)

	assert.Equal(t, 2, pixInfoList[0].Count)
	assert.Equal(t, 1, pixInfoList[1].Count)
}

func TestAsHistogram(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		path   string
		expect []int
	}{
		{
			path:   "../testdata/r1g1b1a1.png",    // 4 (2x2) pix image of each RGBA colors
			expect: []int{3, 1, 3, 1, 3, 1, 1, 3}, // R[0]=3, R[255]=1, G[0]=3, G[255]=1, B[0]=3, B[255]=1, A[0]=1, A[255]=3
		},
		{
			path:   "../testdata/c1m1y1k1.png",
			expect: []int{2, 2, 2, 2, 2, 2, 0, 4},
		},
	} {
		imgRGBA, err := pallet.Load(test.path)
		require.NoError(t, err)

		hist := pallet.AsHistogram(imgRGBA)

		// R
		assert.Equal(t, test.expect[0], hist.R[0], "Count of red with strength 0 (R[0]) did not match")
		assert.Equal(t, test.expect[1], hist.R[255], "Count of red with strength 255 (R[255]) did not match")
		// G
		assert.Equal(t, test.expect[2], hist.G[0], "Count of green with strength 0 (G[0]) did not match")
		assert.Equal(t, test.expect[3], hist.G[255], "Count of green with strength 255 (G[255]) did not match")
		// B
		assert.Equal(t, test.expect[4], hist.B[0], "Count of blue with strength 0 (B[0]) did not match")
		assert.Equal(t, test.expect[5], hist.B[255], "Count of blue with strength 255 (B[255]) did not match")
		// A
		assert.Equal(t, test.expect[6], hist.A[0], "Count of alpha with strength 0 (A[0]) did not match")
		assert.Equal(t, test.expect[7], hist.A[255], "Count of alpha with strength 255 (A[255]) did not match")
	}
}

func TestRGBAHelpers(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 255, pallet.Uint32ToInt(0xffff))
	assert.Equal(t, "012034056078", pallet.ColorToString(
		color.RGBA{R: 12, G: 34, B: 56, A: 78},
	))
}

func TestPixelAnalysis_subimage(t *testing.T) {
	t.Parallel()

	parent := image.NewRGBA(image.Rect(0, 0, 4, 4))
	subimage, ok := parent.SubImage(image.Rect(1, 1, 3, 3)).(*image.RGBA)
	require.True(t, ok)

	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	blue := color.RGBA{R: 0, G: 0, B: 255, A: 255}

	subimage.SetRGBA(1, 1, red)
	subimage.SetRGBA(2, 1, red)
	subimage.SetRGBA(1, 2, red)
	subimage.SetRGBA(2, 2, blue)

	histogram := pallet.AsHistogram(subimage)
	assert.Equal(t, 3, histogram.R[255])
	assert.Equal(t, 1, histogram.B[255])

	occurrences := pallet.ByOccurrence(subimage)
	require.Len(t, occurrences, 2)
	assert.Equal(t, pallet.PixInfo{R: 255, G: 0, B: 0, A: 255, Count: 3}, occurrences[0])
	assert.Equal(t, pallet.PixInfo{R: 0, G: 0, B: 255, A: 255, Count: 1}, occurrences[1])
}

func TestAsHistogram_InJSON_defalt(t *testing.T) {
	t.Parallel()

	// 2x2 pixel image with each RGBA color of 1-pixel
	imgRGBA, err := pallet.Load("../testdata/r1g1b1a1.png")
	require.NoError(t, err)

	hist := pallet.AsHistogram(imgRGBA)
	perLine := false // Get JSON as one line

	expectJSON := "{" +
		"\"r\":[" +
		"3,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,1]," +
		"\"g\":[" +
		"3,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,1]," +
		"\"b\":[" +
		"3,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,1]," +
		"\"a\":[" +
		"1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0," +
		"0,0,0,0,3]" +
		"}"

	actualJSON, err := hist.InJSON(perLine)
	require.NoError(t, err)

	assert.JSONEq(t, expectJSON, actualJSON)
}

func TestAsHistogram_InJSON_perline(t *testing.T) {
	t.Parallel()

	// 2x2 pixel image with each RGBA color of 1-pixel
	imgRGBA, err := pallet.Load("../testdata/r1g1b1a1.png")
	require.NoError(t, err)

	hist := pallet.AsHistogram(imgRGBA)
	perLine := true // Get JSON elements per line

	actualJSON, err := hist.InJSON(perLine)
	require.NoError(t, err)

	assert.Contains(t, actualJSON, "{\n  \"r\": [\n    3,\n")
	assert.Contains(t, actualJSON, "0,\n    1\n  ],\n  \"g\": [\n    3,\n")
	assert.Contains(t, actualJSON, "1\n  ],\n  \"b\": [\n    3,\n")
	assert.Contains(t, actualJSON, "],\n  \"a\": [\n    1,\n ")
	assert.Contains(t, actualJSON, " 3\n  ]\n}")
}

//nolint:paralleltest // Do not parallelize due to global state change
func TestAsHistogram_InJSON_default_fail(t *testing.T) {
	// Backup and defer restore
	oldJSONMarshal := pallet.JSONMarshal

	defer func() {
		pallet.JSONMarshal = oldJSONMarshal
	}()

	// Mock JSONMarshal
	pallet.JSONMarshal = func(_ any) ([]byte, error) {
		return []byte{}, errors.New("forced fail")
	}

	// 2x2 pixel image with each RGBA color of 1-pixel
	imgRGBA, err := pallet.Load("../testdata/r1g1b1a1.png")
	require.NoError(t, err)

	hist := pallet.AsHistogram(imgRGBA)
	perLine := false

	_, err = hist.InJSON(perLine)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "JSON conversion failed")
}

//nolint:paralleltest // Do not parallelize due to global state change
func TestAsHistogram_InJSON_perline_fail(t *testing.T) {
	// Backup and defer restore
	oldJSONMarshalIndent := pallet.JSONMarshalIndent

	defer func() {
		pallet.JSONMarshalIndent = oldJSONMarshalIndent
	}()

	// Mock JSONMarshalIndent
	pallet.JSONMarshalIndent = func(_ any, _ string, _ string) ([]byte, error) {
		return []byte{}, errors.New("forced fail")
	}

	// 2x2 pixel image with each RGBA color of 1-pixel
	imgRGBA, err := pallet.Load("../testdata/r1g1b1a1.png")
	require.NoError(t, err)

	hist := pallet.AsHistogram(imgRGBA)
	perLine := true

	_, err = hist.InJSON(perLine)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "JSON conversion failed")
}

func TestGetKey(t *testing.T) {
	t.Parallel()

	pixInfo := pallet.PixInfo{
		R:     12, // Red --> 012
		G:     34, // Green --> 034
		B:     56, // Blue --> 056
		A:     0,  // Alpha --> 255
		Count: 0,
	}

	key := pixInfo.GetKey()

	assert.Equal(t, "012034056000", key)
}

func TestLoad(t *testing.T) {
	t.Parallel()

	pathFile := t.TempDir()

	_, err := pallet.Load(pathFile)

	require.Error(t, err)
}

//nolint:paralleltest // Do not parallelize due to global state change
func TestPixInfoList_InJSON(t *testing.T) {
	pathFileImg := "../testdata/r1g2b4a2.png"

	// Load image
	imgRGBA, err := pallet.Load(pathFileImg)
	require.NoError(t, err)

	// Count by occurrence
	pixInfoList := pallet.ByOccurrence(imgRGBA)

	// Mock json.Marshal to fail
	oldJSONMarshal := pallet.JSONMarshal

	defer func() {
		pallet.JSONMarshal = oldJSONMarshal
	}()

	pallet.JSONMarshal = func(_ any) ([]byte, error) {
		return []byte{}, errors.New("forced fail")
	}

	// Fail on unmarshal in single line
	{
		outputPerLine := false

		_, err := pixInfoList.InJSON(outputPerLine)
		require.Error(t, err)
	}

	// Fail on unmarshal each element per line
	{
		outputPerLine := true

		_, err := pixInfoList.InJSON(outputPerLine)
		require.Error(t, err)
	}

	// Fail on JSON Lines output
	{
		_, err := pixInfoList.InJSONL()
		require.Error(t, err)
	}
}

func TestPixInfoList_InJSONL(t *testing.T) {
	t.Parallel()

	pixInfoList := pallet.PixInfoList{
		{R: 255, G: 0, B: 0, A: 255, Count: 2},
		{R: 0, G: 0, B: 255, A: 255, Count: 1},
	}

	actual, err := pixInfoList.InJSONL()

	require.NoError(t, err)
	assert.Equal(
		t,
		"{\"r\":255,\"g\":0,\"b\":0,\"a\":255,\"count\":2}\n"+
			"{\"r\":0,\"g\":0,\"b\":255,\"a\":255,\"count\":1}\n",
		actual,
	)

	actual, err = (pallet.PixInfoList{}).InJSONL()

	require.NoError(t, err)
	assert.Empty(t, actual)
}

func TestPixInfoList_InJSON_empty(t *testing.T) {
	t.Parallel()

	actual, err := (pallet.PixInfoList{}).InJSON(true)

	require.NoError(t, err)
	assert.Equal(t, "[\n]", actual)
}
