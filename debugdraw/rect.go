package debugdraw

import (
	"image"
	"image/color"
)

func RectRGBA(rgba *image.RGBA, bounds image.Rectangle, color color.RGBA) {
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		rgba.SetRGBA(x, bounds.Min.Y, color)
		rgba.SetRGBA(x, bounds.Max.Y-1, color)
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rgba.SetRGBA(bounds.Min.X, y, color)
		rgba.SetRGBA(bounds.Max.X-1, y, color)
	}
}

func RectPalette(m *image.Paletted, bounds image.Rectangle, color uint8) {
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		m.SetColorIndex(x, bounds.Min.Y, color)
		m.SetColorIndex(x, bounds.Max.Y-1, color)
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		m.SetColorIndex(bounds.Min.X, y, color)
		m.SetColorIndex(bounds.Max.X-1, y, color)
	}
}
