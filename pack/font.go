package pack

import (
	"fmt"
	"image"
	"io/ioutil"
	"strconv"
	"unicode"

	"github.com/golang/freetype/truetype"

	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var SupportedFonts = []string{
	".ttf",
	// ".otf",
}

type Font struct {
	Name      string
	Font      *truetype.Font
	Face      font.Face
	Padding   int
	MaxBounds fixed.Rectangle26_6
	Glyphs    map[rune]*Glyph
	Kern      map[[2]rune]fixed.Int26_6
}

func LoadFont(name, filename string, fontSize int) (*Font, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	fnt, err := truetype.Parse(data)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(fnt, &truetype.Options{
		Size:    float64(fontSize),
		Hinting: font.HintingFull,
	})

	font := &Font{}
	font.Name = name
	font.Font = fnt
	font.Face = face
	font.MaxBounds = fnt.Bounds(fixed.I(fontSize))
	font.Glyphs = make(map[rune]*Glyph, 256)
	font.Kern = make(map[[2]rune]fixed.Int26_6, 256)

	return font, nil
}

func (font *Font) Boxes() []Box {
	boxes := []Box{}
	for r, g := range font.Glyphs {
		boxes = append(boxes, Box{
			Name:    font.Name + "_" + strconv.Itoa(int(r)),
			Padding: font.Padding,
			Size:    g.Size,
			Place:   &g.Place,
		})
	}
	return boxes
}

func (font *Font) IncludeExtendedAscii() {
	for r := rune(0); r < 256; r++ {
		font.Include(r)
	}
}

func (font *Font) IncludeAscii() {
	for r := rune(0); r < 128; r++ {
		font.Include(r)
	}
}

func (font *Font) IncludeString(s string) {
	for _, r := range s {
		font.Include(r)
	}
}

func (font *Font) Include(r rune) bool {
	if _, ok := font.Glyphs[r]; ok {
		return false
	}

	// skip non-graphic and combining marks
	if !unicode.IsGraphic(r) || unicode.IsMark(r) {
		return false
	}

	var ok bool

	g := &Glyph{}
	g.Rune = r

	font.Font.Index(r)
	g.Bounds, g.Advance, ok = font.Face.GlyphBounds(r)
	if !ok {
		fmt.Println("no glyph ", g.Rune)
		return false
	}

	g.Size.X = (g.Bounds.Max.X - g.Bounds.Min.X).Ceil()
	g.Size.Y = (g.Bounds.Max.Y - g.Bounds.Min.Y).Ceil()

	g.Size.X += 2
	g.Size.Y += 2

	font.Glyphs[r] = g
	font.includeKern(r)

	return true
}

func (font *Font) includeKern(x rune) {
	for b := range font.Glyphs {
		if k := font.Face.Kern(x, b); k != 0 {
			font.Kern[[2]rune{x, b}] = k
		}
		if k := font.Face.Kern(b, x); k != 0 {
			font.Kern[[2]rune{b, x}] = k
		}
	}
}

//
//   advance          ----  maxAscent
//  |-------|          |
//           ■■■      ----  bounds.Max.Y
//          ■   ■      |
//          ■          |
//   ■■■■    ■■■■      |
//  ■    ■       ■     |
//  ■    ■       ■     |
//  •■■■■■  •■■■■   • ----  point.Y
//       ■             |
//  ■    ■             |
//   ■■■■             ----  bounds.Min.Y
//                     |
//                    ----  maxDescent
//
type Glyph struct {
	Rune    rune
	Size    image.Point         // destination size in pixels
	Bounds  fixed.Rectangle26_6 // such that point + bounds, gives image bounds
	Advance fixed.Int26_6
	Place   image.Rectangle // without padding
}

func (font *Font) Draw(m draw.Image) {
	for _, g := range font.Glyphs {
		dot := fixed.P(g.Place.Min.X, g.Place.Min.Y).Sub(g.Bounds.Min)
		dr, mask, maskp, _, ok := font.Face.Glyph(dot, g.Rune)
		if !ok {
			continue
		}
		draw.Draw(m, dr, mask, maskp, draw.Src)
	}
}
