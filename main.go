package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/adinfinit/texpack/debugdraw"
	"github.com/adinfinit/texpack/maxrect"
	"github.com/adinfinit/texpack/pack"
	"github.com/adinfinit/texpack/sdf"
	"github.com/adinfinit/texpack/walk"

	"image/color"
	"image/gif"
	"image/png"
)

var (
	atlasimage = flag.String("atlas", "atlas.png", "output image")
	atlasdata  = flag.String("data", "atlas.json", "output description")

	padding = flag.Int("pad", 3, "padding")
	place   = flag.String("place", "automatic", "placing algorithm (short-side, long-side, bottom-left, area, contact-point)")

	sdfRadius = flag.Int("sdf-radius", 0, "use alpha channel for distance field")

	debug      = flag.Bool("debug-size", false, "debug box sizes")
	debugPlace = flag.Bool("debug-place", false, "debug placing gif")
)

func main() {
	flag.Parse()

	var images []*pack.Image
	var fonts []*pack.Font

	folder := flag.Arg(0)
	if folder == "" {
		fmt.Println("folder not specified")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *padding < *sdfRadius {
		*padding = *sdfRadius
	}

	maxSize := image.Point{2048, 2048}

	walk.Exts(folder, pack.SupportedImages, func(name, path string) {
		fmt.Println("adding image: ", name)
		name = filepath.ToSlash(name)
		m, err := pack.LoadImage(name, path)
		if err != nil {
			fmt.Println(err)
			return
		}
		m.Padding = *padding
		images = append(images, m)
	})

	walk.Exts(folder, pack.SupportedFonts, func(name, path string) {
		fmt.Println("adding font: ", name)
		name = filepath.ToSlash(name)
		f, err := pack.LoadFont(name, path, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Padding = *padding
		f.IncludeExtendedAscii()
		fonts = append(fonts, f)
	})

	boxes := []pack.Box{}
	for _, m := range images {
		boxes = append(boxes, m.Box())
	}
	for _, f := range fonts {
		boxes = append(boxes, f.Boxes()...)
	}
	pack.SortBoxes(boxes)

	size, ok := pack.PlaceBoxes(maxSize, maxrect.ParseRule(*place), boxes)
	if !ok {
		fmt.Println("images do not fit")
		return
	}

	if *debugPlace {
		WritePlacingGif(size, boxes, *atlasimage+".gif")
	}

	dst := image.NewRGBA(image.Rectangle{image.ZP, size})
	for _, m := range images {
		m.Draw(dst)
	}
	for _, f := range fonts {
		f.Draw(dst)
	}

	if *sdfRadius > 0 {
		sdf.ApplyRGBA_Alpha(dst, *sdfRadius)
	}

	if *debug {
		for _, box := range boxes {
			debugdraw.RectRGBA(dst, *box.Place, color.RGBA{0, 0xFF, 0, 0x80})
			debugdraw.RectRGBA(dst, box.Place.Inset(-box.Padding), color.RGBA{0xFF, 0, 0, 0xFF})
		}
	}

	f, err := os.Create(*atlasimage)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	if err := png.Encode(f, dst); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	atlas := &Atlas{}
	atlas.Frames = make(map[string]Frame)
	atlas.Meta.App = "texpack"
	atlas.Meta.Version = "1.0"
	atlas.Meta.Image = *atlasimage
	atlas.Meta.Format = "RGBA8888"
	atlas.Meta.Size.Width = size.X
	atlas.Meta.Size.Height = size.Y
	for _, m := range images {
		frame := Frame{
			Name:    m.Name,
			Rotated: m.Rotated,
			// Trimmed: m.Trimmed,
			Frame: Rect{
				X:      m.Place.Min.X,
				Y:      m.Place.Min.Y,
				Width:  m.Place.Max.X - m.Place.Min.X,
				Height: m.Place.Max.Y - m.Place.Min.Y,
			},
			SpriteSourceSize: Rect{
				X:      0,
				Y:      0,
				Width:  m.Place.Max.X - m.Place.Min.X,
				Height: m.Place.Max.Y - m.Place.Min.Y,
			},
			SourceSize: Size{
				Width:  m.Place.Max.X - m.Place.Min.X,
				Height: m.Place.Max.Y - m.Place.Min.Y,
			},
			Pivot: Point{0.5, 0.5},
		}
		atlas.Frames[frame.Name] = frame
	}

	fd, err := os.Create(*atlasdata)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fd.Close()

	enc := json.NewEncoder(fd)
	enc.SetIndent("", "\t")
	if err := enc.Encode(atlas); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func WritePlacingGif(size image.Point, boxes []pack.Box, outfile string) {
	palette := color.Palette{
		color.RGBA{0, 0, 0, 0},            // 0 - transparent
		color.RGBA{0xFF, 0, 0, 0xFF},      // 1 - box
		color.RGBA{0x0, 0xFF, 0, 0xFF},    // 2 - free
		color.RGBA{0x0, 0x00, 0xFF, 0xFF}, // 3 - placed
	}

	out := &gif.GIF{}
	context := maxrect.New(size)
	context.SetRule(maxrect.ParseRule(*place))
	context.DebugPlace = func(r image.Rectangle) {
		pal := image.NewPaletted(image.Rectangle{image.ZP, size}, palette)
		for _, r := range context.Used {
			debugdraw.RectPalette(pal, r, 1)
		}
		for _, r := range context.Free {
			debugdraw.RectPalette(pal, r, 2)
		}
		debugdraw.RectPalette(pal, r, 3)

		out.Image = append(out.Image, pal)
		out.Delay = append(out.Delay, 0)
		out.Disposal = append(out.Disposal, 3)
	}

	context.Adds(pack.BoxesToSizes(boxes)...)

	f, err := os.Create(outfile)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	gif.EncodeAll(f, out)
}
