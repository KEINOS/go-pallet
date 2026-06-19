package pallet

import (
	"fmt"
	"image"
	"io"
	"strings"
	"testing"

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
