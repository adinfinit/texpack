package pack

import (
	"image"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/image/draw"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

type Image struct {
	Name string
	Data image.Image
	Size image.Point

	// extra info for GIF-s
	Animated bool
	Index    int // animation-only
	Delay    int // animation-only, ms

	// parameters
	Padding int
	Rotated bool
	SDF     bool

	// final location inside texture
	Place image.Rectangle
}

func (m *Image) Box() Box {
	return Box{
		Padding: m.Padding,
		Size:    m.Size,
		Place:   &m.Place,
	}
}

func (m *Image) Draw(dst draw.Image) {
	draw.Draw(dst, m.Place, m.Data, image.ZP, draw.Src)
}

var SupportedImages = []string{
	".gif",
	".jpg", ".jpeg",
	".png",
	".bmp",
	".tif",
	".tiff",
	".webp",
}

func LoadImage(name, filename string) (*Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	m := &Image{}
	m.Name = name
	m.Data = data
	m.Size = data.Bounds().Size()

	return m, nil
}
