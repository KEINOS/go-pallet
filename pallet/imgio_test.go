package pallet

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	errCloseFailed  = errors.New("close failed")
	errEncodeFailed = errors.New("encode failed")
)

type readCloserStub struct {
	reader   io.Reader
	closeErr error
}

func (r readCloserStub) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	if err != nil {
		return n, fmt.Errorf("read from stub: %w", err)
	}

	return n, nil
}

func (r readCloserStub) Close() error {
	return r.closeErr
}

type writeCloserStub struct {
	closeErr error
}

func (w writeCloserStub) Write(p []byte) (int, error) {
	return len(p), nil
}

func (w writeCloserStub) Close() error {
	return w.closeErr
}

func TestEncode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		value   image.Image
		encoder Encoder
		format  string
	}{
		{
			format:  "png",
			encoder: PNGEncoder(),
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
			encoder: JPEGEncoder(95),
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
			encoder: BMPEncoder(),
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

		_, err := Open(pathDirUnknown)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "no such file or directory")
	}

	// Path is a text file
	{
		pathFileTmp := filepath.Join(pathDirTmp, "tmp.txt")
		err := os.WriteFile(pathFileTmp, []byte("foo bar"), 0o600)
		require.NoError(t, err)

		_, err = Open(pathFileTmp)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to decode image file")
	}
}

func TestOpen_saved_image(t *testing.T) {
	t.Parallel()

	// Create image
	encPNG := PNGEncoder()
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

	err := Save(pathFileTmp, imgRAW, encPNG)
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
	actualImg, err := Open(pathFileTmp)
	require.NoError(t, err)

	assert.EqualValues(t, expectImg, actualImg)
}

func TestSave_fail_save(t *testing.T) {
	t.Parallel()

	// Create image
	encPNG := PNGEncoder()
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

	err := Save(pathDirTmp, imgRAW, encPNG)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create file to save",
		"it should contain the error reason")
}

func TestDecodeAndClose_close_error(t *testing.T) {
	t.Parallel()

	mockImg := image.NewRGBA(image.Rect(0, 0, 1, 1))

	actual, err := decodeAndClose(
		readCloserStub{reader: strings.NewReader("ignored"), closeErr: errCloseFailed},
		func(_ io.Reader) (image.Image, string, error) {
			return mockImg, "png", nil
		},
	)

	require.Error(t, err)
	assert.Nil(t, actual)
	assert.Contains(t, err.Error(), "failed to close image file")
}

func TestEncodeAndClose_encode_error(t *testing.T) {
	t.Parallel()

	mockImg := image.NewRGBA(image.Rect(0, 0, 1, 1))

	err := encodeAndClose(
		writeCloserStub{closeErr: nil},
		mockImg,
		func(_ io.Writer, _ image.Image) error {
			return errEncodeFailed
		},
	)

	require.Error(t, err)
	assert.ErrorIs(t, err, errEncodeFailed)
}

func TestEncodeAndClose_close_error(t *testing.T) {
	t.Parallel()

	mockImg := image.NewRGBA(image.Rect(0, 0, 1, 1))

	err := encodeAndClose(
		writeCloserStub{closeErr: errCloseFailed},
		mockImg,
		func(_ io.Writer, _ image.Image) error {
			return nil
		},
	)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to close image file")
}
