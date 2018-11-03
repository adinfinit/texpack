package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/adinfinit/texpack/maxrect"
	"github.com/adinfinit/texpack/pack"
)

var (
	atlasimage = flag.String("out", "", "output image")
	atlasdata  = flag.String("xml", "", "output xml")

	padding  = flag.Int("pad", 3, "padding")
	fontsize = flag.Int("size", 32, "font size")

	ascii   = flag.Bool("ascii", true, "include ascii characters")
	exascii = flag.Bool("extended-ascii", false, "include extended ascii characters")

	place = flag.String("place", "automatic", "placing algorithm (short-side, long-side, bottom-left, area, contact-point)")
)

func main() {
	flag.Parse()

	fontname := flag.Arg(0)

	if fontname == "" {
		fmt.Fprintln(os.Stderr, "font not specified")
		flag.PrintDefaults()
		os.Exit(1)
	}

	maxSize := image.Point{2048, 2048}

	if !strings.EqualFold(filepath.Ext(fontname), ".ttf") {
		fmt.Fprintln(os.Stderr, "only ttf is supported")
		os.Exit(1)
	}

	if *atlasimage == "" {
		*atlasimage = fontname[:len(fontname)-4] + ".png"
	}
	if *atlasdata == "" {
		*atlasdata = fontname[:len(fontname)-4] + ".xml"
	}

	font, err := pack.LoadFont("", fontname, *fontsize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load font %q: %v\n", fontname, err)
		os.Exit(1)
	}
	font.Padding = *padding

	if *ascii {
		font.IncludeAscii()
	}
	if *exascii {
		font.IncludeExtendedAscii()
	}

	boxes := font.Boxes()
	pack.SortBoxes(boxes)

	size, ok := pack.PlaceBoxes(maxSize, maxrect.ParseRule(*place), boxes)
	if !ok {
		fmt.Fprintln(os.Stderr, "images do not fit")
		os.Exit(1)
	}

	dst := image.NewRGBA(image.Rectangle{image.ZP, size})
	font.Draw(dst)

	atlas, err := os.Create(*atlasimage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create atlas %v: %v", *atlasimage, err)
		os.Exit(1)
	}
	defer atlas.Close()

	if err := png.Encode(atlas, dst); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode bitmap: %v", err)
		os.Exit(1)
	}

	x := &Font{}
	x.Info.Face = fontname[:len(fontname)-4]
	x.Info.Size = strconv.Itoa(*fontsize)
	x.Info.Bold = "0"
	x.Info.Italic = "0"
	x.Info.Unicode = "1"
	x.Info.StretchH = "0"
	x.Info.Smooth = "1"
	x.Info.Aa = "1"
	pad := strconv.Itoa(*padding)
	x.Info.Padding = pad + "," + pad + "," + pad + "," + pad
	x.Info.Spacing = ""

	x.Common.LineHeight = x.Info.Size // TODO: fix
	x.Common.Base = "0"               // TODO: fix
	x.Common.ScaleW = strconv.Itoa(size.X)
	x.Common.ScaleH = strconv.Itoa(size.Y)
	x.Common.Pages = "1"
	x.Common.Packed = "0"

	x.Pages.Page = []Page{{
		ID:   "0",
		File: *atlasimage,
	}}

	x.Chars.Count = strconv.Itoa(len(font.Glyphs))
	for _, glyph := range font.Glyphs {
		x.Chars.Char = append(x.Chars.Char, Char{
			ID: strconv.Itoa(int(glyph.Rune)),

			X:      strconv.Itoa(glyph.Place.Min.X),
			Y:      strconv.Itoa(glyph.Place.Min.Y),
			Width:  strconv.Itoa(glyph.Place.Dx()),
			Height: strconv.Itoa(glyph.Place.Dy()),

			Xadvance: strconv.Itoa(glyph.Advance.Round()),
			Xoffset:  strconv.Itoa(glyph.Bounds.Min.X.Round()),
			Yoffset:  strconv.Itoa(glyph.Bounds.Min.Y.Round()),

			Page:   "0",
			Chnl:   "0",
			Letter: string(glyph.Rune),
		})
	}

	x.Kernings.Count = strconv.Itoa(len(font.Kern))
	for pair, kern := range font.Kern {
		x.Kernings.Kerning = append(x.Kernings.Kerning, Kerning{
			First:  strconv.Itoa(int(pair[0])),
			Second: strconv.Itoa(int(pair[1])),
			Amount: strconv.Itoa(kern.Round()),
		})
	}

	atlasxml, err := os.Create(*atlasdata)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create xml: %v", err)
		os.Exit(1)
	}
	defer atlasxml.Close()

	enc := xml.NewEncoder(atlasxml)
	enc.Indent("", "\t")
	err = enc.Encode(x)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode xml: %v", err)
		os.Exit(1)
	}

	/*
		<font>
			<info face="Desyrel" size="70" bold="0" italic="0" chasrset="" unicode="0" stretchH="100" smooth="1" aa="1" padding="0,0,0,0" spacing="1,1"/>
			<common lineHeight="87" base="61"
				scaleW="512" scaleH="512" pages="1" packed="0"/>
			<pages>
				<page id="0" file="desyrel.png"/>
				<page id="1" file="desyrel.png"/>
			</pages>
			<chars count="95">
				<char id="102" x="1" y="1" width="38" height="74" xoffset="2" yoffset="9" xadvance="28" page="0" chnl="0" letter="f"/>
				<char id="83" x="40" y="1" width="35" height="74" xoffset="5" yoffset="5" xadvance="31" page="0" chnl="0" letter="S"/>
				<char id="125" x="76" y="1" width="27" height="74" xoffset="-2" yoffset="4" xadvance="20" page="0" chnl="0" letter="}"/>
				<char id="123" x="104" y="1" width="27" height="74" xoffset="1" yoffset="4" xadvance="18" page="0" chnl="0" letter="{"/>
				<char id="93" x="132" y="1" width="23" height="72" xoffset="-3" yoffset="4" xadvance="15" page="0" chnl="0" letter="]"/>
			</chars>
			<kernings count="1816">
				<kerning first="102" second="102" amount="2"/>
				<kerning first="102" second="106" amount="-2"/>
				<kerning first="102" second="41" amount="4"/>
				<kerning first="102" second="100" amount="-9"/>
				<kerning first="102" second="103" amount="-8"/>
				<kerning first="102" second="108" amount="-3"/>
				<kerning first="102" second="63" amount="1"/>
			</kernings>
		</font>
	*/
}
