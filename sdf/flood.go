package sdf

import (
	"image"
	"image/color"
)

// Flood returns new image where colors have been flooded for radius
func Flood(src *image.RGBA, radius int) *image.RGBA {
	prev := image.NewRGBA(src.Bounds())
	next := image.NewRGBA(src.Bounds())
	copy(prev.Pix, src.Pix)
	copy(next.Pix, src.Pix)

	for i := 0; i < radius; i++ {
		FloodStep(prev, next)
		prev, next = next, prev
	}

	return prev
}

func FloodStep(prev, next *image.RGBA) {
	size := prev.Bounds().Size()
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			pc := prev.RGBAAt(x, y)
			if pc.A == 0xFF {
				next.SetRGBA(x, y, pc)
				continue
			}

			tr, tg, tb, ta := 0, 0, 0, 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					sx, sy := x+dx, y+dy
					if sx < 0 || sx > size.X || sy < 0 || sy > size.Y {
						continue
					}
					s := prev.RGBAAt(sx, sy)
					tr += int(s.R)
					tg += int(s.G)
					tb += int(s.B)
					ta += int(s.A)
				}
			}

			if ta > 0xFF {
				next.SetRGBA(x, y, color.RGBA{
					R: uint8(tr * 0xFF / ta),
					G: uint8(tg * 0xFF / ta),
					B: uint8(tb * 0xFF / ta),
					A: uint8(0xFF),
				})
			} else {
				next.SetRGBA(x, y, color.RGBA{
					R: uint8(tr),
					G: uint8(tg),
					B: uint8(tb),
					A: uint8(ta),
				})
			}
		}
	}
}
