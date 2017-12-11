// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"golang.org/x/image/math/fixed"

	"github.com/adinfinit/texpack/maxrect"
	"github.com/adinfinit/texpack/pack"
	"github.com/adinfinit/texpack/walk"
)

var (
	atlaspath = "assets/atlas.png"
	datapath  = "assets/atlas.go"

	folder    = "resources"
	padding   = 3
	smoothpad = 1 // added to all images to avoid interpolation issues near quad borders
	fontSize  = 64
	maxSize   = image.Point{2048, 2048}

	// for some images we may want to skip smooth borders,
	// because they need to be stretched across the screen
	customSmoothPad = map[string]int{
		"background.png": -2,
	}
)

func main() {
	var images []*pack.Image
	var fonts []*pack.Font

	// collect all images to packing list
	walk.Exts(folder, pack.SupportedImages, func(name, path string) {
		fmt.Println("adding image: ", name)
		name = filepath.ToSlash(name)
		m, err := pack.LoadImage(name, path)
		if err != nil {
			fmt.Println(err)
			return
		}
		m.Padding = padding
		images = append(images, m)
	})

	// collect all fonts to packing list
	walk.Exts(folder, pack.SupportedFonts, func(name, path string) {
		fmt.Println("adding font: ", name)
		name = filepath.ToSlash(name)
		f, err := pack.LoadFont(name, path, fontSize)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Padding = padding
		// specify which characters to include in the font
		f.IncludeAscii()
		fonts = append(fonts, f)
	})

	// list of all boxes that need to be packed
	boxes := []pack.Box{}
	for _, m := range images {
		boxes = append(boxes, m.Box())
	}
	for _, f := range fonts {
		boxes = append(boxes, f.Boxes()...)
	}
	// sort based on size, to get a consistent output
	pack.SortBoxes(boxes)

	// pack boxes together
	size, ok := pack.PlaceBoxes(maxSize, maxrect.Automatic, boxes)
	if !ok {
		fmt.Println("images do not fit")
		os.Exit(1)
		return
	}

	// draw all the images to their respective boxes
	dst := image.NewRGBA(image.Rectangle{image.ZP, size})
	for _, m := range images {
		m.Draw(dst)
		// adding padding to when necessary
		if spad, ok := customSmoothPad[m.Name]; ok {
			m.Place = m.Place.Inset(-spad)
		} else {
			m.Place = m.Place.Inset(-smoothpad)
		}
	}

	// draw all glyphs to their respective boxes
	for _, f := range fonts {
		f.Draw(dst)
		// add padding to avoid aliasing near quad borders
		for _, g := range f.Glyphs {
			g.Bounds.Min = g.Bounds.Min.Sub(fixed.P(smoothpad, smoothpad))
			g.Bounds.Max = g.Bounds.Max.Add(fixed.P(smoothpad, smoothpad))
			g.Place = g.Place.Inset(-smoothpad)
		}
	}

	{ // save image atlas to a file
		f, err := os.Create(atlaspath)
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

	{ // generate go file for using the image atlas
		var buf bytes.Buffer
		err := T.Execute(&buf, map[string]interface{}{
			"Atlas": map[string]interface{}{
				"Name": "atlas.png",
				"Size": size,
			},
			"Images": images,
			"Fonts":  fonts,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		data, err := format.Source(buf.Bytes())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = ioutil.WriteFile(datapath, data, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func tofloat(v interface{}) float32 {
	switch v := v.(type) {
	case fixed.Int26_6:
		trunc := v.Floor()
		frac := int(v) - trunc
		return float32(trunc) + float32(frac)/(1<<6)
	case int:
		return float32(v)
	case float32:
		return float32(v)
	case float64:
		return float32(v)
	default:
		panic("cannot convert to float " + fmt.Sprintf("%T", v))
	}
}

type KernEntry struct {
	A, B    rune
	Advance fixed.Int26_6
}

var T = template.Must(template.New("").Funcs(template.FuncMap{
	"varname": func(name string) string {
		ext := path.Ext(name)
		if ext != "" {
			name = name[:len(name)-len(ext)]
		}
		name = strings.Replace(name, "-", "", -1)

		up := true
		name = strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				if up {
					up = false
					return unicode.ToUpper(r)
				}
				return r
			}
			up = true
			return '_'
		}, name)
		return name
	},
	"neg": func(a interface{}) float32 {
		return -tofloat(a)
	},
	"sub": func(a, b interface{}) float32 {
		return tofloat(a) - tofloat(b)
	},
	"div": func(a, b interface{}) float32 {
		return tofloat(a) / tofloat(b)
	},
	"max": func(a, b interface{}) float32 {
		af, bf := tofloat(a), tofloat(b)
		if af >= bf {
			return af
		}
		return bf
	},
	"sortkern": func(kern map[[2]rune]fixed.Int26_6) []KernEntry {
		entries := make([]KernEntry, 0, len(kern))
		for x, a := range kern {
			entries = append(entries, KernEntry{x[0], x[1], a})
		}
		sort.Slice(entries, func(i, k int) bool {
			if entries[i].A < entries[k].A {
				return true
			} else if entries[i].A > entries[k].A {
				return false
			}

			return entries[i].B < entries[k].B
		})
		return entries
	},
}).Parse(`// GENERATED DO NOT EDIT
package assets

type Vector struct {
	X, Y float32
}
type Rectangle struct {
	Min, Max Vector
}

type Image struct {
	Atlas string
	Name  string
	Size  Vector
	UV    Rectangle
}

type Font struct {
	Atlas string
	Name  string
	
	MaxAscent  float32
	MaxDescent float32

	Glyphs map[rune]Glyph
	Kern   map[[2]rune]float32
}

type Glyph struct {
	Bounds  Rectangle
	UV      Rectangle
	Advance float32
}

var ImageByPath = map[string]*Image{
	{{ range $image := .Images -}}
	{{ printf "%q" $image.Name }}: {{ varname $image.Name }},
	{{ end -}}
}

var FontByPath = map[string]*Font{
	{{ range $font := .Fonts -}}
	{{ printf "%q" $font.Name }}: {{ varname $font.Name }},
	{{ end -}}
}

{{ $atlas := .Atlas }}
{{ range $image := .Images }}
var {{varname $image.Name}} = &Image{
	Atlas: {{ printf "%q" $atlas.Name }},
	Name: {{ printf "%q" $image.Name }},
	Size: Vector{ {{$image.Size.X}}, {{$image.Size.Y}} },
	UV: Rectangle{
		Vector{ {{ div $image.Place.Min.X $atlas.Size.X }}, {{ div $image.Place.Min.Y $atlas.Size.Y }} },
		Vector{ {{ div $image.Place.Max.X $atlas.Size.X }}, {{ div $image.Place.Max.Y $atlas.Size.Y }} },
	},
}
{{ end }}

{{ range $font := .Fonts }}
var {{ varname $font.Name }} = &Font{
	Atlas: {{ printf "%q" $atlas.Name }},
	Name: {{ printf "%q" $font.Name }},
	{{ $scale := sub $font.MaxBounds.Max.Y $font.MaxBounds.Min.Y }}
	MaxAscent: {{ neg (div $font.MaxBounds.Max.Y $scale) }},
	MaxDescent: {{ neg (div $font.MaxBounds.Min.Y $scale) }},

	Glyphs: map[rune]Glyph{
		{{- range $glyph := $font.Glyphs }}
		{{ printf "%q" $glyph.Rune }}: {
			{{ if (and (eq $glyph.Bounds.Min.X 0) (eq $glyph.Bounds.Max.X 0))}} Empty: true, {{ else -}}
			Bounds: Rectangle{
				Vector{ {{ div $glyph.Bounds.Min.X $scale }}, {{ div $glyph.Bounds.Min.Y $scale }} },
				Vector{ {{ div $glyph.Bounds.Max.X $scale }}, {{ div $glyph.Bounds.Max.Y $scale }} },
			},
			UV: Rectangle{
				Vector{ {{ div $glyph.Place.Min.X $atlas.Size.X }}, {{ div $glyph.Place.Min.Y $atlas.Size.Y }} },
				Vector{ {{ div $glyph.Place.Max.X $atlas.Size.X }}, {{ div $glyph.Place.Max.Y $atlas.Size.Y }} },
			},{{ end }}
			Advance: {{ div $glyph.Advance $scale }},
		},{{ end }}
	},
	Kern: map[[2]rune]float32{
		{{ range $entry := sortkern $font.Kern -}}
		{ {{printf "%q" $entry.A}}, {{printf "%q" $entry.B}} }: {{ div $entry.Advance $scale }},
		{{ end -}}
	},
}
{{ end }}
`))
