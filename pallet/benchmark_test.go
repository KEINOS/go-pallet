package pallet

import (
	"image"
	"image/color"
	"testing"
)

func benchmarkImage() *image.RGBA {
	const (
		width  = 1024
		height = 768
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for pixelY := range height {
		for pixelX := range width {
			img.SetRGBA(pixelX, pixelY, color.RGBA{
				R: uint8(pixelX * 17),
				G: uint8((pixelY * 31) % 256),
				B: uint8((pixelX + pixelY) * 13),
				A: uint8(255 - (pixelX*pixelY)%32),
			})
		}
	}

	return img
}

func BenchmarkAsHistogram(b *testing.B) {
	img := benchmarkImage()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_ = AsHistogram(img)
	}
}

func BenchmarkByOccurrence(b *testing.B) {
	img := benchmarkImage()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_ = ByOccurrence(img)
	}
}

func BenchmarkColorToString(b *testing.B) {
	pixel := color.RGBA{R: 12, G: 34, B: 56, A: 78}

	b.ReportAllocs()

	for b.Loop() {
		_ = ColorToString(pixel)
	}
}

func BenchmarkDiff(b *testing.B) {
	first := benchmarkImage()
	second := benchmarkImage()

	for index := 0; index < len(second.Pix); index += 4 {
		second.Pix[index] ^= 0xff
	}

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_, _ = Diff(first, second)
	}
}
