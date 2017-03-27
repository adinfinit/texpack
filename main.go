package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	"github.com/loov/texpack/maxrect"
	"github.com/loov/texpack/pack"

	"image/color"
	"image/gif"
	"image/png"
)

const (
	defaultPadding = 6
)

var (
	atlasimage = flag.String("atlas", "atlas.png", "output image")
	atlasdata  = flag.String("data", "atlas.json", "output description")

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

	maxSize := image.Point{2048, 2048}

	Walk(folder, pack.SupportedImages, func(name, path string) {
		fmt.Println("adding image: ", name)
		m, err := pack.LoadImage(name, path)
		if err != nil {
			fmt.Println(err)
			return
		}
		m.Padding = defaultPadding
		images = append(images, m)
	})

	Walk(folder, pack.SupportedFonts, func(name, path string) {
		fmt.Println("adding font: ", name)
		f, err := pack.LoadFont(name, path, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Padding = defaultPadding
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

	size, ok := pack.PlaceBoxes(maxSize, boxes)
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

	if *debug {
		for _, box := range boxes {
			drawRect(dst, *box.Place, color.RGBA{0, 0xFF, 0, 0x80})
			drawRect(dst, box.Place.Inset(-box.Padding), color.RGBA{0xFF, 0, 0, 0xFF})
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
	context.DebugPlace = func(r image.Rectangle) {
		pal := image.NewPaletted(image.Rectangle{image.ZP, size}, palette)
		for _, r := range context.Used {
			drawRectPal(pal, r, 1)
		}
		for _, r := range context.Free {
			drawRectPal(pal, r, 2)
		}
		drawRectPal(pal, r, 3)

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
