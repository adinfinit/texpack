// GENERATED DO NOT EDIT
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
	"background.png": Background,
	"error.png":      Error,
}

var FontByPath = map[string]*Font{
	"Arcade.ttf": Arcade,
}

var Background = &Image{
	Atlas: "atlas.png",
	Name:  "background.png",
	Size:  Vector{64, 64},
	UV: Rectangle{
		Vector{0.009765625, 0.0078125},
		Vector{0.12695312, 0.1015625},
	},
}

var Error = &Image{
	Atlas: "atlas.png",
	Name:  "error.png",
	Size:  Vector{64, 64},
	UV: Rectangle{
		Vector{0.00390625, 0.1125},
		Vector{0.1328125, 0.215625},
	},
}

var Arcade = &Font{
	Atlas: "atlas.png",
	Name:  "Arcade.ttf",

	MaxAscent:  -1,
	MaxDescent: -0,

	Glyphs: map[rune]Glyph{
		' ': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.017829565},
				Vector{0.017829565, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.98828125, 0.2890625},
				Vector{0.99609375, 0.2953125},
			},
			Advance: 1.1232626,
		},
		'!': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.5883757, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.71484375, 0.8609375},
				Vector{0.75390625, 0.95},
			},
			Advance: 1.1232626,
		},
		'"': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.8914783, -0.39225045},
			},
			UV: Rectangle{
				Vector{0.8222656, 0.740625},
				Vector{0.9082031, 0.79375},
			},
			Advance: 1.1232626,
		},
		'#': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.82016003},
				Vector{0.8914783, -0.12480696},
			},
			UV: Rectangle{
				Vector{0.8046875, 0.6703125},
				Vector{0.890625, 0.734375},
			},
			Advance: 1.1232626,
		},
		'$': {
			Bounds: Rectangle{
				Vector{0.106977396, -0.96279657},
				Vector{0.85581917, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.48046875, 0.8609375},
				Vector{0.56640625, 0.95},
			},
			Advance: 1.1232626,
		},
		'%': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.87890625, 0.003125},
				Vector{0.9941406, 0.0921875},
			},
			Advance: 1.1232626,
		},
		'&': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.87890625, 0.0984375},
				Vector{0.9941406, 0.1875},
			},
			Advance: 1.1232626,
		},
		'\'': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.46356872, -0.39225045},
			},
			UV: Rectangle{
				Vector{0.76171875, 0.94375},
				Vector{0.80078125, 0.996875},
			},
			Advance: 1.1232626,
		},
		'(': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.74884176, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.8496094, 0.575},
				Vector{0.9199219, 0.6640625},
			},
			Advance: 1.1232626,
		},
		')': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.87364876, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.57421875, 0.8609375},
				Vector{0.64453125, 0.95},
			},
			Advance: 1.1232626,
		},
		'*': {
			Bounds: Rectangle{
				Vector{0.106977396, -0.8379896},
				Vector{0.85581917, -0.106977396},
			},
			UV: Rectangle{
				Vector{0.6816406, 0.765625},
				Vector{0.7675781, 0.8328125},
			},
			Advance: 1.1232626,
		},
		'+': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.82016003},
				Vector{0.8914783, -0.12480696},
			},
			UV: Rectangle{
				Vector{0.8984375, 0.6703125},
				Vector{0.984375, 0.734375},
			},
			Advance: 1.1232626,
		},
		',': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.53488696},
				Vector{0.46356872, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.9160156, 0.740625},
				Vector{0.9550781, 0.7921875},
			},
			Advance: 1.1232626,
		},
		'-': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.53488696},
				Vector{0.8914783, -0.39225045},
			},
			UV: Rectangle{
				Vector{0.6816406, 0.8390625},
				Vector{0.7675781, 0.8546875},
			},
			Advance: 1.1232626,
		},
		'.': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.26744348},
				Vector{0.46356872, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.40429688, 0.95625},
				Vector{0.44335938, 0.984375},
			},
			Advance: 1.1232626,
		},
		'/': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.140625, 0.19375},
				Vector{0.25585938, 0.2828125},
			},
			Advance: 1.1232626,
		},
		'0': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.26367188, 0.19375},
				Vector{0.37890625, 0.2828125},
			},
			Advance: 1.1232626,
		},
		'1': {
			Bounds: Rectangle{
				Vector{0.12480696, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6191406, 0.6703125},
				Vector{0.71875, 0.759375},
			},
			Advance: 1.1232626,
		},
		'2': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.38671875, 0.19375},
				Vector{0.5019531, 0.2828125},
			},
			Advance: 1.1232626,
		},
		'3': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5097656, 0.19375},
				Vector{0.625, 0.2828125},
			},
			Advance: 1.1232626,
		},
		'4': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6328125, 0.19375},
				Vector{0.7480469, 0.2828125},
			},
			Advance: 1.1232626,
		},
		'5': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7558594, 0.19375},
				Vector{0.87109375, 0.2828125},
			},
			Advance: 1.1232626,
		},
		'6': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.87890625, 0.19375},
				Vector{0.9941406, 0.2828125},
			},
			Advance: 1.1232626,
		},
		'7': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.12695312, 0.2890625},
				Vector{0.2421875, 0.378125},
			},
			Advance: 1.1232626,
		},
		'8': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.12695312, 0.384375},
				Vector{0.2421875, 0.4734375},
			},
			Advance: 1.1232626,
		},
		'9': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.25, 0.2890625},
				Vector{0.36523438, 0.378125},
			},
			Advance: 1.1232626,
		},
		':': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.82016003},
				Vector{0.46356872, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.76171875, 0.8609375},
				Vector{0.80078125, 0.9375},
			},
			Advance: 1.1232626,
		},
		';': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.46356872, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7753906, 0.765625},
				Vector{0.8144531, 0.8546875},
			},
			Advance: 1.1232626,
		},
		'<': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.82016003},
				Vector{0.74884176, -0.12480696},
			},
			UV: Rectangle{
				Vector{0.8222656, 0.8},
				Vector{0.8925781, 0.8640625},
			},
			Advance: 1.1232626,
		},
		'=': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.6775235},
				Vector{0.8914783, -0.26744348},
			},
			UV: Rectangle{
				Vector{0.80859375, 0.940625},
				Vector{0.89453125, 0.9796875},
			},
			Advance: 1.1232626,
		},
		'>': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.82016003},
				Vector{0.87364876, -0.12480696},
			},
			UV: Rectangle{
				Vector{0.80859375, 0.8703125},
				Vector{0.87890625, 0.934375},
			},
			Advance: 1.1232626,
		},
		'?': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.12695312, 0.4796875},
				Vector{0.2421875, 0.56875},
			},
			Advance: 1.1232626,
		},
		'@': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.25, 0.384375},
				Vector{0.36523438, 0.4734375},
			},
			Advance: 1.1232626,
		},
		'A': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.37304688, 0.2890625},
				Vector{0.48828125, 0.378125},
			},
			Advance: 1.1232626,
		},
		'B': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.12695312, 0.575},
				Vector{0.2421875, 0.6640625},
			},
			Advance: 1.1232626,
		},
		'C': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.25, 0.4796875},
				Vector{0.36523438, 0.56875},
			},
			Advance: 1.1232626,
		},
		'D': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.37304688, 0.384375},
				Vector{0.48828125, 0.4734375},
			},
			Advance: 1.1232626,
		},
		'E': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.49609375, 0.2890625},
				Vector{0.6113281, 0.378125},
			},
			Advance: 1.1232626,
		},
		'F': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.12695312, 0.6703125},
				Vector{0.2421875, 0.759375},
			},
			Advance: 1.1232626,
		},
		'G': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.25, 0.575},
				Vector{0.36523438, 0.6640625},
			},
			Advance: 1.1232626,
		},
		'H': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.37304688, 0.4796875},
				Vector{0.48828125, 0.56875},
			},
			Advance: 1.1232626,
		},
		'I': {
			Bounds: Rectangle{
				Vector{0.12480696, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7421875, 0.575},
				Vector{0.8417969, 0.6640625},
			},
			Advance: 1.1232626,
		},
		'J': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.49609375, 0.384375},
				Vector{0.6113281, 0.4734375},
			},
			Advance: 1.1232626,
		},
		'K': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6191406, 0.2890625},
				Vector{0.734375, 0.378125},
			},
			Advance: 1.1232626,
		},
		'L': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.12695312, 0.765625},
				Vector{0.2421875, 0.8546875},
			},
			Advance: 1.1232626,
		},
		'M': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.25, 0.6703125},
				Vector{0.36523438, 0.759375},
			},
			Advance: 1.1232626,
		},
		'N': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.37304688, 0.575},
				Vector{0.48828125, 0.6640625},
			},
			Advance: 1.1232626,
		},
		'O': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.49609375, 0.4796875},
				Vector{0.6113281, 0.56875},
			},
			Advance: 1.1232626,
		},
		'P': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6191406, 0.384375},
				Vector{0.734375, 0.4734375},
			},
			Advance: 1.1232626,
		},
		'Q': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7421875, 0.2890625},
				Vector{0.8574219, 0.378125},
			},
			Advance: 1.1232626,
		},
		'R': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.12695312, 0.8609375},
				Vector{0.2421875, 0.95},
			},
			Advance: 1.1232626,
		},
		'S': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.25, 0.765625},
				Vector{0.36523438, 0.8546875},
			},
			Advance: 1.1232626,
		},
		'T': {
			Bounds: Rectangle{
				Vector{0.12480696, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.8515625, 0.4796875},
				Vector{0.9511719, 0.56875},
			},
			Advance: 1.1232626,
		},
		'U': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.37304688, 0.6703125},
				Vector{0.48828125, 0.759375},
			},
			Advance: 1.1232626,
		},
		'V': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.49609375, 0.575},
				Vector{0.6113281, 0.6640625},
			},
			Advance: 1.1232626,
		},
		'W': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6191406, 0.4796875},
				Vector{0.734375, 0.56875},
			},
			Advance: 1.1232626,
		},
		'X': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7421875, 0.384375},
				Vector{0.8574219, 0.4734375},
			},
			Advance: 1.1232626,
		},
		'Y': {
			Bounds: Rectangle{
				Vector{0.106977396, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.8652344, 0.384375},
				Vector{0.9667969, 0.4734375},
			},
			Advance: 1.1232626,
		},
		'Z': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.8652344, 0.2890625},
				Vector{0.98046875, 0.378125},
			},
			Advance: 1.1232626,
		},
		'[': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.7310122, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.9277344, 0.575},
				Vector{0.9824219, 0.6640625},
			},
			Advance: 1.1232626,
		},
		'\\': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.25, 0.8609375},
				Vector{0.36523438, 0.95},
			},
			Advance: 1.1232626,
		},
		']': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.7310122, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.65234375, 0.8609375},
				Vector{0.70703125, 0.95},
			},
			Advance: 1.1232626,
		},
		'^': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.7310122, -0.6775235},
			},
			UV: Rectangle{
				Vector{0.34179688, 0.95625},
				Vector{0.39648438, 0.984375},
			},
			Advance: 1.1232626,
		},
		'_': {
			Bounds: Rectangle{
				Vector{0, -0.12480696},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.12695312, 0.95625},
				Vector{0.2421875, 0.971875},
			},
			Advance: 1.1232626,
		},
		'`': {
			Bounds: Rectangle{
				Vector{0.39225045, -0.96279657},
				Vector{0.7131826, -0.6775235},
			},
			UV: Rectangle{
				Vector{0.45117188, 0.95625},
				Vector{0.49023438, 0.984375},
			},
			Advance: 1.1232626,
		},
		'a': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.37304688, 0.765625},
				Vector{0.48828125, 0.8546875},
			},
			Advance: 1.1232626,
		},
		'b': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.49609375, 0.6703125},
				Vector{0.6113281, 0.759375},
			},
			Advance: 1.1232626,
		},
		'c': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6191406, 0.575},
				Vector{0.734375, 0.6640625},
			},
			Advance: 1.1232626,
		},
		'd': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.00390625, 0.221875},
				Vector{0.119140625, 0.3109375},
			},
			Advance: 1.1232626,
		},
		'e': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.00390625, 0.3171875},
				Vector{0.119140625, 0.40625},
			},
			Advance: 1.1232626,
		},
		'f': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.00390625, 0.4125},
				Vector{0.119140625, 0.5015625},
			},
			Advance: 1.1232626,
		},
		'g': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.00390625, 0.5078125},
				Vector{0.119140625, 0.596875},
			},
			Advance: 1.1232626,
		},
		'h': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.00390625, 0.603125},
				Vector{0.119140625, 0.6921875},
			},
			Advance: 1.1232626,
		},
		'i': {
			Bounds: Rectangle{
				Vector{0.12480696, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.37304688, 0.8609375},
				Vector{0.47265625, 0.95},
			},
			Advance: 1.1232626,
		},
		'j': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.00390625, 0.6984375},
				Vector{0.119140625, 0.7875},
			},
			Advance: 1.1232626,
		},
		'k': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.00390625, 0.79375},
				Vector{0.119140625, 0.8828125},
			},
			Advance: 1.1232626,
		},
		'l': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.00390625, 0.8890625},
				Vector{0.119140625, 0.978125},
			},
			Advance: 1.1232626,
		},
		'm': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.140625, 0.003125},
				Vector{0.25585938, 0.0921875},
			},
			Advance: 1.1232626,
		},
		'n': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.140625, 0.0984375},
				Vector{0.25585938, 0.1875},
			},
			Advance: 1.1232626,
		},
		'o': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.26367188, 0.003125},
				Vector{0.37890625, 0.0921875},
			},
			Advance: 1.1232626,
		},
		'p': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.26367188, 0.0984375},
				Vector{0.37890625, 0.1875},
			},
			Advance: 1.1232626,
		},
		'q': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.38671875, 0.003125},
				Vector{0.5019531, 0.0921875},
			},
			Advance: 1.1232626,
		},
		'r': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.38671875, 0.0984375},
				Vector{0.5019531, 0.1875},
			},
			Advance: 1.1232626,
		},
		's': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5097656, 0.003125},
				Vector{0.625, 0.0921875},
			},
			Advance: 1.1232626,
		},
		't': {
			Bounds: Rectangle{
				Vector{0.12480696, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.49609375, 0.765625},
				Vector{0.5957031, 0.8546875},
			},
			Advance: 1.1232626,
		},
		'u': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5097656, 0.0984375},
				Vector{0.625, 0.1875},
			},
			Advance: 1.1232626,
		},
		'v': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6328125, 0.003125},
				Vector{0.7480469, 0.0921875},
			},
			Advance: 1.1232626,
		},
		'w': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6328125, 0.0984375},
				Vector{0.7480469, 0.1875},
			},
			Advance: 1.1232626,
		},
		'x': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7558594, 0.003125},
				Vector{0.87109375, 0.0921875},
			},
			Advance: 1.1232626,
		},
		'y': {
			Bounds: Rectangle{
				Vector{0.106977396, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7421875, 0.4796875},
				Vector{0.84375, 0.56875},
			},
			Advance: 1.1232626,
		},
		'z': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7558594, 0.0984375},
				Vector{0.87109375, 0.1875},
			},
			Advance: 1.1232626,
		},
		'{': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.74884176, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6035156, 0.765625},
				Vector{0.6738281, 0.8546875},
			},
			Advance: 1.1232626,
		},
		'|': {
			Bounds: Rectangle{
				Vector{0.41008002, -0.96279657},
				Vector{0.5883757, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.9589844, 0.4796875},
				Vector{0.9824219, 0.56875},
			},
			Advance: 1.1232626,
		},
		'}': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.87364876, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7265625, 0.6703125},
				Vector{0.796875, 0.759375},
			},
			Advance: 1.1232626,
		},
		'~': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.87364876, -0.6775235},
			},
			UV: Rectangle{
				Vector{0.25, 0.95625},
				Vector{0.33398438, 0.984375},
			},
			Advance: 1.1232626,
		},
	},
	Kern: map[[2]rune]float32{},
}
