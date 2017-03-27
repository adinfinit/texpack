package sdf

import (
	"image"
	"image/color"
	"math"
)

const (
	minDist = 0
	maxDist = 10 << 10
)

var (
	empty  = Point{maxDist, maxDist}
	inside = Point{minDist, minDist}
)

type Point struct{ Dx, Dy float64 }

func (p Point) Offset(v float64) Point { return Point{p.Dx + v, p.Dy + v} }

func (p Point) Distance2() float64 { return p.Dx*p.Dx + p.Dy*p.Dy }

type Field struct {
	Width  int
	Height int
	Pix    []Point
}

func NewField(w, h int) *Field {
	field := &Field{}
	field.Width = w
	field.Height = h
	field.Pix = make([]Point, w*h)
	return field
}

func (field *Field) Set(x, y int, p Point) {
	if 0 <= x && x < field.Width && 0 <= y && y < field.Height {
		field.Pix[field.Width*y+x] = p
	}
}

func (field *Field) Get(x, y int) Point {
	if 0 <= x && x < field.Width && 0 <= y && y < field.Height {
		return field.Pix[field.Width*y+x]
	} else {
		return empty
	}
}

func Compare(field *Field, p *Point, x, y int, xoff, yoff int) {
	other := field.Get(x+xoff, y+yoff)
	other.Dx += float64(xoff)
	other.Dy += float64(yoff)
	if other.Distance2() < p.Distance2() {
		*p = other
	}
}

func Process(field *Field) {
	// Pass 0
	for y := 0; y < field.Height; y++ {
		for x := 0; x < field.Width; x++ {
			p := field.Get(x, y)
			Compare(field, &p, x, y, -1, +0)
			Compare(field, &p, x, y, +0, -1)
			Compare(field, &p, x, y, -1, -1)
			Compare(field, &p, x, y, +1, -1)
			field.Set(x, y, p)
		}

		for x := field.Width - 1; x >= 0; x-- {
			p := field.Get(x, y)
			Compare(field, &p, x, y, 1, 0)
			field.Set(x, y, p)
		}
	}

	// Pass 1
	for y := field.Height - 1; y >= 0; y-- {
		for x := field.Width - 1; x >= 0; x-- {
			p := field.Get(x, y)
			Compare(field, &p, x, y, +1, +0)
			Compare(field, &p, x, y, +0, +1)
			Compare(field, &p, x, y, -1, +1)
			Compare(field, &p, x, y, +1, +1)
			field.Set(x, y, p)
		}

		for x := 0; x < field.Width; x++ {
			p := field.Get(x, y)
			Compare(field, &p, x, y, -1, 0)
			field.Set(x, y, p)
		}
	}
}

func ensureDist(x int) int {
	if x < minDist {
		return minDist
	} else if x > maxDist {
		return maxDist
	}
	return x
}

// ApplyRGBA_Alpha uses alpha channel to spread rgba radius
func ApplyRGBA_Alpha(m *image.RGBA, radius int) {
	size := m.Bounds().Size()

	low := NewField(size.X, size.Y)
	high := NewField(size.X, size.Y)
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			_, _, _, a := m.At(x, y).RGBA()

			af := float64(a) / 0xFFFF
			afi := 1 - af
			if af < 0.5 {
				low.Set(x, y, inside.Offset(af*0.5))
				high.Set(x, y, empty.Offset(-afi*0.5))
			} else {
				low.Set(x, y, empty.Offset(-afi*0.5))
				high.Set(x, y, inside.Offset(af*0.5))
			}

			if af < 0.5 {
				low.Set(x, y, inside)
				high.Set(x, y, empty)
			} else {
				low.Set(x, y, empty)
				high.Set(x, y, inside)
			}
		}
	}

	Process(low)
	Process(high)

	flooded := Flood(m, radius+1)

	scale := 128.0 / float64(radius)
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			lowd := math.Sqrt(low.Get(x, y).Distance2())
			highd := math.Sqrt(high.Get(x, y).Distance2())

			dist := lowd - highd

			alpha := int(dist*scale + 128.0)
			if alpha < 0 {
				alpha = 0
			} else if alpha > 255 {
				alpha = 255
			}

			original := m.RGBAAt(x, y)
			base := flooded.RGBAAt(x, y)
			final := Blend(original, base, uint8(alpha))
			m.SetRGBA(x, y, final)
		}
	}
}

func Blend(top, base color.RGBA, targetAlpha uint8) color.RGBA {
	fill := int(0xFF - top.A)

	top.R = clamp8((int(top.R) + int(base.R)*fill/0xFF) * int(targetAlpha) / 0xFF)
	top.G = clamp8((int(top.G) + int(base.G)*fill/0xFF) * int(targetAlpha) / 0xFF)
	top.B = clamp8((int(top.B) + int(base.B)*fill/0xFF) * int(targetAlpha) / 0xFF)
	top.A = targetAlpha

	return top
}

func clamp8(v int) uint8 {
	if v < 0 {
		return 0
	} else if v > 0xFF {
		return 0xFF
	}
	return uint8(v)
}

func f8(v uint8) float64 { return float64(v) / 0xFF }
func sat8(v float64) uint8 {
	x := int(v * 0xFF)
	if x < 0 {
		return 0
	} else if x > 0xFF {
		return 0xFF
	}
	return uint8(x)
}
