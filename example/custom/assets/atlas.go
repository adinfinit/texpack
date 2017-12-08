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
		Vector{0.0065104165, 0.0055803573},
		Vector{0.084635414, 0.07254464},
	},
}

var Error = &Image{
	Atlas: "atlas.png",
	Name:  "error.png",
	Size:  Vector{64, 64},
	UV: Rectangle{
		Vector{0.0026041667, 0.08035714},
		Vector{0.088541664, 0.15401785},
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
				Vector{0.92578125, 0.8136161},
				Vector{0.93098956, 0.81808037},
			},
			Advance: 1.1232626,
		},
		'!': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.5883757, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.85286456, 0.78125},
				Vector{0.87890625, 0.8448661},
			},
			Advance: 1.1232626,
		},
		'"': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.8914783, -0.39225045},
			},
			UV: Rectangle{
				Vector{0.8932292, 0.6551339},
				Vector{0.9505208, 0.69308037},
			},
			Advance: 1.1232626,
		},
		'#': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.82016003},
				Vector{0.8914783, -0.12480696},
			},
			UV: Rectangle{
				Vector{0.9036458, 0.5546875},
				Vector{0.9609375, 0.6004464},
			},
			Advance: 1.1232626,
		},
		'$': {
			Bounds: Rectangle{
				Vector{0.106977396, -0.96279657},
				Vector{0.85581917, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7291667, 0.84933037},
				Vector{0.7864583, 0.9129464},
			},
			Advance: 1.1232626,
		},
		'%': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.17578125, 0.14174107},
				Vector{0.25260416, 0.20535715},
			},
			Advance: 1.1232626,
		},
		'&': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.2578125, 0.14174107},
				Vector{0.3346354, 0.20535715},
			},
			Advance: 1.1232626,
		},
		'\'': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.46356872, -0.39225045},
			},
			UV: Rectangle{
				Vector{0.9348958, 0.71316963},
				Vector{0.9609375, 0.7511161},
			},
			Advance: 1.1232626,
		},
		'(': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.74884176, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7916667, 0.84933037},
				Vector{0.8385417, 0.9129464},
			},
			Advance: 1.1232626,
		},
		')': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.87364876, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.78125, 0.91741073},
				Vector{0.828125, 0.98102677},
			},
			Advance: 1.1232626,
		},
		'*': {
			Bounds: Rectangle{
				Vector{0.106977396, -0.8379896},
				Vector{0.85581917, -0.106977396},
			},
			UV: Rectangle{
				Vector{0.8333333, 0.91741073},
				Vector{0.890625, 0.96540177},
			},
			Advance: 1.1232626,
		},
		'+': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.82016003},
				Vector{0.8914783, -0.12480696},
			},
			UV: Rectangle{
				Vector{0.9036458, 0.60491073},
				Vector{0.9609375, 0.65066963},
			},
			Advance: 1.1232626,
		},
		',': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.53488696},
				Vector{0.46356872, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.93619794, 0.75558037},
				Vector{0.96223956, 0.79241073},
			},
			Advance: 1.1232626,
		},
		'-': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.53488696},
				Vector{0.8914783, -0.39225045},
			},
			UV: Rectangle{
				Vector{0.8932292, 0.69754463},
				Vector{0.9505208, 0.70870537},
			},
			Advance: 1.1232626,
		},
		'.': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.26744348},
				Vector{0.46356872, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.93619794, 0.796875},
				Vector{0.96223956, 0.81696427},
			},
			Advance: 1.1232626,
		},
		'/': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.33984375, 0.14174107},
				Vector{0.41666666, 0.20535715},
			},
			Advance: 1.1232626,
		},
		'0': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.421875, 0.14174107},
				Vector{0.4986979, 0.20535715},
			},
			Advance: 1.1232626,
		},
		'1': {
			Bounds: Rectangle{
				Vector{0.12480696, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.65755206, 0.84933037},
				Vector{0.7239583, 0.9129464},
			},
			Advance: 1.1232626,
		},
		'2': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.50390625, 0.14174107},
				Vector{0.5807292, 0.20535715},
			},
			Advance: 1.1232626,
		},
		'3': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5859375, 0.14174107},
				Vector{0.66276044, 0.20535715},
			},
			Advance: 1.1232626,
		},
		'4': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.66796875, 0.14174107},
				Vector{0.7447917, 0.20535715},
			},
			Advance: 1.1232626,
		},
		'5': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.75, 0.14174107},
				Vector{0.82682294, 0.20535715},
			},
			Advance: 1.1232626,
		},
		'6': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.83203125, 0.14174107},
				Vector{0.9088542, 0.20535715},
			},
			Advance: 1.1232626,
		},
		'7': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.9140625, 0.14174107},
				Vector{0.99088544, 0.20535715},
			},
			Advance: 1.1232626,
		},
		'8': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.20982143},
				Vector{0.16145833, 0.2734375},
			},
			Advance: 1.1232626,
		},
		'9': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.2779018},
				Vector{0.16145833, 0.34151787},
			},
			Advance: 1.1232626,
		},
		':': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.82016003},
				Vector{0.46356872, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.9036458, 0.4955357},
				Vector{0.9296875, 0.55022323},
			},
			Advance: 1.1232626,
		},
		';': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.46356872, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.84375, 0.84933037},
				Vector{0.8697917, 0.9129464},
			},
			Advance: 1.1232626,
		},
		'<': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.82016003},
				Vector{0.74884176, -0.12480696},
			},
			UV: Rectangle{
				Vector{0.8828125, 0.71316963},
				Vector{0.9296875, 0.7589286},
			},
			Advance: 1.1232626,
		},
		'=': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.6775235},
				Vector{0.8914783, -0.26744348},
			},
			UV: Rectangle{
				Vector{0.8333333, 0.9698661},
				Vector{0.890625, 0.99776787},
			},
			Advance: 1.1232626,
		},
		'>': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.82016003},
				Vector{0.87364876, -0.12480696},
			},
			UV: Rectangle{
				Vector{0.88411456, 0.76339287},
				Vector{0.93098956, 0.80915177},
			},
			Advance: 1.1232626,
		},
		'?': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.16666667, 0.20982143},
				Vector{0.24348958, 0.2734375},
			},
			Advance: 1.1232626,
		},
		'@': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.34598213},
				Vector{0.16145833, 0.4095982},
			},
			Advance: 1.1232626,
		},
		'A': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.16666667, 0.2779018},
				Vector{0.24348958, 0.34151787},
			},
			Advance: 1.1232626,
		},
		'B': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.24869792, 0.20982143},
				Vector{0.32552084, 0.2734375},
			},
			Advance: 1.1232626,
		},
		'C': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.4140625},
				Vector{0.16145833, 0.47767857},
			},
			Advance: 1.1232626,
		},
		'D': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.16666667, 0.34598213},
				Vector{0.24348958, 0.4095982},
			},
			Advance: 1.1232626,
		},
		'E': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.24869792, 0.2779018},
				Vector{0.32552084, 0.34151787},
			},
			Advance: 1.1232626,
		},
		'F': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.33072916, 0.20982143},
				Vector{0.4075521, 0.2734375},
			},
			Advance: 1.1232626,
		},
		'G': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.48214287},
				Vector{0.16145833, 0.5457589},
			},
			Advance: 1.1232626,
		},
		'H': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.16666667, 0.4140625},
				Vector{0.24348958, 0.47767857},
			},
			Advance: 1.1232626,
		},
		'I': {
			Bounds: Rectangle{
				Vector{0.12480696, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7291667, 0.78125},
				Vector{0.79557294, 0.8448661},
			},
			Advance: 1.1232626,
		},
		'J': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.24869792, 0.34598213},
				Vector{0.32552084, 0.4095982},
			},
			Advance: 1.1232626,
		},
		'K': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.33072916, 0.2779018},
				Vector{0.4075521, 0.34151787},
			},
			Advance: 1.1232626,
		},
		'L': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.4127604, 0.20982143},
				Vector{0.48958334, 0.2734375},
			},
			Advance: 1.1232626,
		},
		'M': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.55022323},
				Vector{0.16145833, 0.61383927},
			},
			Advance: 1.1232626,
		},
		'N': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.16666667, 0.48214287},
				Vector{0.24348958, 0.5457589},
			},
			Advance: 1.1232626,
		},
		'O': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.24869792, 0.4140625},
				Vector{0.32552084, 0.47767857},
			},
			Advance: 1.1232626,
		},
		'P': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.33072916, 0.34598213},
				Vector{0.4075521, 0.4095982},
			},
			Advance: 1.1232626,
		},
		'Q': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.4127604, 0.2779018},
				Vector{0.48958334, 0.34151787},
			},
			Advance: 1.1232626,
		},
		'R': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.49479166, 0.20982143},
				Vector{0.57161456, 0.2734375},
			},
			Advance: 1.1232626,
		},
		'S': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.6183036},
				Vector{0.16145833, 0.68191963},
			},
			Advance: 1.1232626,
		},
		'T': {
			Bounds: Rectangle{
				Vector{0.12480696, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.65755206, 0.91741073},
				Vector{0.7239583, 0.98102677},
			},
			Advance: 1.1232626,
		},
		'U': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.16666667, 0.55022323},
				Vector{0.24348958, 0.61383927},
			},
			Advance: 1.1232626,
		},
		'V': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.24869792, 0.48214287},
				Vector{0.32552084, 0.5457589},
			},
			Advance: 1.1232626,
		},
		'W': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.33072916, 0.4140625},
				Vector{0.4075521, 0.47767857},
			},
			Advance: 1.1232626,
		},
		'X': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.4127604, 0.34598213},
				Vector{0.48958334, 0.4095982},
			},
			Advance: 1.1232626,
		},
		'Y': {
			Bounds: Rectangle{
				Vector{0.106977396, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.49479166, 0.34598213},
				Vector{0.5625, 0.4095982},
			},
			Advance: 1.1232626,
		},
		'Z': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.49479166, 0.2779018},
				Vector{0.57161456, 0.34151787},
			},
			Advance: 1.1232626,
		},
		'[': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.7310122, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.953125, 0.002232143},
				Vector{0.9895833, 0.06584822},
			},
			Advance: 1.1232626,
		},
		'\\': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.57682294, 0.20982143},
				Vector{0.6536458, 0.2734375},
			},
			Advance: 1.1232626,
		},
		']': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.7310122, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.9453125, 0.20982143},
				Vector{0.9817708, 0.2734375},
			},
			Advance: 1.1232626,
		},
		'^': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.7310122, -0.6775235},
			},
			UV: Rectangle{
				Vector{0.88411456, 0.8136161},
				Vector{0.92057294, 0.83370537},
			},
			Advance: 1.1232626,
		},
		'_': {
			Bounds: Rectangle{
				Vector{0, -0.12480696},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.9754464},
				Vector{0.079427086, 0.98660713},
			},
			Advance: 1.1232626,
		},
		'`': {
			Bounds: Rectangle{
				Vector{0.39225045, -0.96279657},
				Vector{0.7131826, -0.6775235},
			},
			UV: Rectangle{
				Vector{0.14583333, 0.97209823},
				Vector{0.171875, 0.9921875},
			},
			Advance: 1.1232626,
		},
		'a': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.6863839},
				Vector{0.16145833, 0.75},
			},
			Advance: 1.1232626,
		},
		'b': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.16666667, 0.6183036},
				Vector{0.24348958, 0.68191963},
			},
			Advance: 1.1232626,
		},
		'c': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.24869792, 0.55022323},
				Vector{0.32552084, 0.61383927},
			},
			Advance: 1.1232626,
		},
		'd': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.15848215},
				Vector{0.079427086, 0.22209822},
			},
			Advance: 1.1232626,
		},
		'e': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.2265625},
				Vector{0.079427086, 0.29017857},
			},
			Advance: 1.1232626,
		},
		'f': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.29464287},
				Vector{0.079427086, 0.35825893},
			},
			Advance: 1.1232626,
		},
		'g': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.3627232},
				Vector{0.079427086, 0.4263393},
			},
			Advance: 1.1232626,
		},
		'h': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.43080357},
				Vector{0.079427086, 0.49441963},
			},
			Advance: 1.1232626,
		},
		'i': {
			Bounds: Rectangle{
				Vector{0.12480696, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.65755206, 0.78125},
				Vector{0.7239583, 0.8448661},
			},
			Advance: 1.1232626,
		},
		'j': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.49888393},
				Vector{0.079427086, 0.5625},
			},
			Advance: 1.1232626,
		},
		'k': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.56696427},
				Vector{0.079427086, 0.63058037},
			},
			Advance: 1.1232626,
		},
		'l': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.63504463},
				Vector{0.079427086, 0.69866073},
			},
			Advance: 1.1232626,
		},
		'm': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.703125},
				Vector{0.079427086, 0.7667411},
			},
			Advance: 1.1232626,
		},
		'n': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.77120537},
				Vector{0.079427086, 0.8348214},
			},
			Advance: 1.1232626,
		},
		'o': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.83928573},
				Vector{0.079427086, 0.90290177},
			},
			Advance: 1.1232626,
		},
		'p': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.0026041667, 0.9073661},
				Vector{0.079427086, 0.97098213},
			},
			Advance: 1.1232626,
		},
		'q': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.09375, 0.07366072},
				Vector{0.17057292, 0.13727678},
			},
			Advance: 1.1232626,
		},
		'r': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.17578125, 0.07366072},
				Vector{0.25260416, 0.13727678},
			},
			Advance: 1.1232626,
		},
		's': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.2578125, 0.07366072},
				Vector{0.3346354, 0.13727678},
			},
			Advance: 1.1232626,
		},
		't': {
			Bounds: Rectangle{
				Vector{0.12480696, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5859375, 0.92410713},
				Vector{0.65234375, 0.98772323},
			},
			Advance: 1.1232626,
		},
		'u': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.33984375, 0.07366072},
				Vector{0.41666666, 0.13727678},
			},
			Advance: 1.1232626,
		},
		'v': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.421875, 0.07366072},
				Vector{0.4986979, 0.13727678},
			},
			Advance: 1.1232626,
		},
		'w': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.50390625, 0.07366072},
				Vector{0.5807292, 0.13727678},
			},
			Advance: 1.1232626,
		},
		'x': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5859375, 0.07366072},
				Vector{0.66276044, 0.13727678},
			},
			Advance: 1.1232626,
		},
		'y': {
			Bounds: Rectangle{
				Vector{0.106977396, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.4127604, 0.4140625},
				Vector{0.48046875, 0.47767857},
			},
			Advance: 1.1232626,
		},
		'z': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.66796875, 0.07366072},
				Vector{0.7447917, 0.13727678},
			},
			Advance: 1.1232626,
		},
		'{': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.74884176, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.80078125, 0.78125},
				Vector{0.84765625, 0.8448661},
			},
			Advance: 1.1232626,
		},
		'|': {
			Bounds: Rectangle{
				Vector{0.41008002, -0.96279657},
				Vector{0.5883757, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.875, 0.84933037},
				Vector{0.890625, 0.9129464},
			},
			Advance: 1.1232626,
		},
		'}': {
			Bounds: Rectangle{
				Vector{0.26744348, -0.96279657},
				Vector{0.87364876, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7291667, 0.91741073},
				Vector{0.7760417, 0.98102677},
			},
			Advance: 1.1232626,
		},
		'~': {
			Bounds: Rectangle{
				Vector{0.14263652, -0.96279657},
				Vector{0.87364876, -0.6775235},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.96875},
				Vector{0.140625, 0.98883927},
			},
			Advance: 1.1232626,
		},
		'\u00a0': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.017829565},
				Vector{0.017829565, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.8932292, 0.63839287},
				Vector{0.8984375, 0.64285713},
			},
			Advance: 1.1232626,
		},
		'¡': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.09375, 0.002232143},
				Vector{0.16015625, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'¢': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.16536458, 0.002232143},
				Vector{0.23177083, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'£': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.23697917, 0.002232143},
				Vector{0.3033854, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'¤': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.30859375, 0.002232143},
				Vector{0.375, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'¥': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.38020834, 0.002232143},
				Vector{0.4466146, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'¦': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.4518229, 0.002232143},
				Vector{0.5182292, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'§': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5234375, 0.002232143},
				Vector{0.58984375, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'¨': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.59505206, 0.002232143},
				Vector{0.6614583, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'©': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.75, 0.07366072},
				Vector{0.82682294, 0.13727678},
			},
			Advance: 1.1232626,
		},
		'ª': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6666667, 0.002232143},
				Vector{0.73307294, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'«': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.73828125, 0.002232143},
				Vector{0.8046875, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'¬': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.8098958, 0.002232143},
				Vector{0.87630206, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'®': {
			Bounds: Rectangle{
				Vector{-0.017829565, -0.96279657},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.83203125, 0.07366072},
				Vector{0.9088542, 0.13727678},
			},
			Advance: 1.1232626,
		},
		'¯': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.88151044, 0.002232143},
				Vector{0.9479167, 0.069196425},
			},
			Advance: 1.1232626,
		},
		'°': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.33072916, 0.48214287},
				Vector{0.3971354, 0.54910713},
			},
			Advance: 1.1232626,
		},
		'±': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.75446427},
				Vector{0.15104167, 0.8214286},
			},
			Advance: 1.1232626,
		},
		'²': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.82589287},
				Vector{0.15104167, 0.89285713},
			},
			Advance: 1.1232626,
		},
		'³': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.084635414, 0.8973214},
				Vector{0.15104167, 0.96428573},
			},
			Advance: 1.1232626,
		},
		'´': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.40234375, 0.48214287},
				Vector{0.46875, 0.54910713},
			},
			Advance: 1.1232626,
		},
		'µ': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.57682294, 0.2779018},
				Vector{0.6432292, 0.34486607},
			},
			Advance: 1.1232626,
		},
		'¶': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6588542, 0.20982143},
				Vector{0.72526044, 0.2767857},
			},
			Advance: 1.1232626,
		},
		'·': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.73046875, 0.20982143},
				Vector{0.796875, 0.2767857},
			},
			Advance: 1.1232626,
		},
		'¸': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.8020833, 0.20982143},
				Vector{0.86848956, 0.2767857},
			},
			Advance: 1.1232626,
		},
		'¹': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.87369794, 0.20982143},
				Vector{0.9401042, 0.2767857},
			},
			Advance: 1.1232626,
		},
		'º': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.16666667, 0.6863839},
				Vector{0.23307292, 0.75334823},
			},
			Advance: 1.1232626,
		},
		'»': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.24869792, 0.6183036},
				Vector{0.31510416, 0.68526787},
			},
			Advance: 1.1232626,
		},
		'¼': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.4856771, 0.4140625},
				Vector{0.5520833, 0.4810268},
			},
			Advance: 1.1232626,
		},
		'½': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.33072916, 0.5535714},
				Vector{0.3971354, 0.62053573},
			},
			Advance: 1.1232626,
		},
		'¾': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.40234375, 0.5535714},
				Vector{0.46875, 0.62053573},
			},
			Advance: 1.1232626,
		},
		'¿': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.15625, 0.7578125},
				Vector{0.22265625, 0.82477677},
			},
			Advance: 1.1232626,
		},
		'À': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.15625, 0.8292411},
				Vector{0.22265625, 0.89620537},
			},
			Advance: 1.1232626,
		},
		'Á': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.15625, 0.90066963},
				Vector{0.22265625, 0.9676339},
			},
			Advance: 1.1232626,
		},
		'Â': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.47395834, 0.48549107},
				Vector{0.54036456, 0.55245537},
			},
			Advance: 1.1232626,
		},
		'Ã': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.47395834, 0.55691963},
				Vector{0.54036456, 0.6238839},
			},
			Advance: 1.1232626,
		},
		'Ä': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6484375, 0.28125},
				Vector{0.71484375, 0.3482143},
			},
			Advance: 1.1232626,
		},
		'Å': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.72005206, 0.28125},
				Vector{0.7864583, 0.3482143},
			},
			Advance: 1.1232626,
		},
		'Æ': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7916667, 0.28125},
				Vector{0.85807294, 0.3482143},
			},
			Advance: 1.1232626,
		},
		'Ç': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.86328125, 0.28125},
				Vector{0.9296875, 0.3482143},
			},
			Advance: 1.1232626,
		},
		'È': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.23828125, 0.68973213},
				Vector{0.3046875, 0.7566964},
			},
			Advance: 1.1232626,
		},
		'É': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.22786458, 0.76116073},
				Vector{0.29427084, 0.828125},
			},
			Advance: 1.1232626,
		},
		'Ê': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.22786458, 0.83258927},
				Vector{0.29427084, 0.8995536},
			},
			Advance: 1.1232626,
		},
		'Ë': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.22786458, 0.90401787},
				Vector{0.29427084, 0.97098213},
			},
			Advance: 1.1232626,
		},
		'Ì': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5677083, 0.34933037},
				Vector{0.63411456, 0.41629463},
			},
			Advance: 1.1232626,
		},
		'Í': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.63932294, 0.35267857},
				Vector{0.7057292, 0.41964287},
			},
			Advance: 1.1232626,
		},
		'Î': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7109375, 0.35267857},
				Vector{0.77734375, 0.41964287},
			},
			Advance: 1.1232626,
		},
		'Ï': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.78255206, 0.35267857},
				Vector{0.8489583, 0.41964287},
			},
			Advance: 1.1232626,
		},
		'Ð': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.8541667, 0.35267857},
				Vector{0.92057294, 0.41964287},
			},
			Advance: 1.1232626,
		},
		'Ñ': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.3203125, 0.625},
				Vector{0.38671875, 0.69196427},
			},
			Advance: 1.1232626,
		},
		'Ò': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.3919271, 0.625},
				Vector{0.45833334, 0.69196427},
			},
			Advance: 1.1232626,
		},
		'Ó': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.46354166, 0.62834823},
				Vector{0.52994794, 0.6953125},
			},
			Advance: 1.1232626,
		},
		'Ô': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5572917, 0.42075893},
				Vector{0.62369794, 0.4877232},
			},
			Advance: 1.1232626,
		},
		'Õ': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.54557294, 0.4921875},
				Vector{0.6119792, 0.55915177},
			},
			Advance: 1.1232626,
		},
		'Ö': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.54557294, 0.5636161},
				Vector{0.6119792, 0.63058037},
			},
			Advance: 1.1232626,
		},
		'×': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.62890625, 0.42410713},
				Vector{0.6953125, 0.49107143},
			},
			Advance: 1.1232626,
		},
		'Ø': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7005208, 0.42410713},
				Vector{0.76692706, 0.49107143},
			},
			Advance: 1.1232626,
		},
		'Ù': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.77213544, 0.42410713},
				Vector{0.8385417, 0.49107143},
			},
			Advance: 1.1232626,
		},
		'Ú': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.84375, 0.42410713},
				Vector{0.91015625, 0.49107143},
			},
			Advance: 1.1232626,
		},
		'Û': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6171875, 0.4955357},
				Vector{0.68359375, 0.5625},
			},
			Advance: 1.1232626,
		},
		'Ü': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.68880206, 0.4955357},
				Vector{0.7552083, 0.5625},
			},
			Advance: 1.1232626,
		},
		'Ý': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7604167, 0.4955357},
				Vector{0.82682294, 0.5625},
			},
			Advance: 1.1232626,
		},
		'Þ': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.83203125, 0.4955357},
				Vector{0.8984375, 0.5625},
			},
			Advance: 1.1232626,
		},
		'ß': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.9140625, 0.07366072},
				Vector{0.99088544, 0.13727678},
			},
			Advance: 1.1232626,
		},
		'à': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6171875, 0.56696427},
				Vector{0.68359375, 0.6339286},
			},
			Advance: 1.1232626,
		},
		'á': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.68880206, 0.56696427},
				Vector{0.7552083, 0.6339286},
			},
			Advance: 1.1232626,
		},
		'â': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7604167, 0.56696427},
				Vector{0.82682294, 0.6339286},
			},
			Advance: 1.1232626,
		},
		'ã': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.83203125, 0.56696427},
				Vector{0.8984375, 0.6339286},
			},
			Advance: 1.1232626,
		},
		'ä': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.30989584, 0.6964286},
				Vector{0.3763021, 0.76339287},
			},
			Advance: 1.1232626,
		},
		'å': {
			Bounds: Rectangle{
				Vector{0, -0.96279657},
				Vector{1.0162853, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.09375, 0.14174107},
				Vector{0.17057292, 0.20535715},
			},
			Advance: 1.1232626,
		},
		'æ': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.3815104, 0.6964286},
				Vector{0.44791666, 0.76339287},
			},
			Advance: 1.1232626,
		},
		'ç': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.29947916, 0.76785713},
				Vector{0.3658854, 0.8348214},
			},
			Advance: 1.1232626,
		},
		'è': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.29947916, 0.83928573},
				Vector{0.3658854, 0.90625},
			},
			Advance: 1.1232626,
		},
		'é': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.37109375, 0.76785713},
				Vector{0.4375, 0.8348214},
			},
			Advance: 1.1232626,
		},
		'ê': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.37109375, 0.83928573},
				Vector{0.4375, 0.90625},
			},
			Advance: 1.1232626,
		},
		'ë': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.29947916, 0.91071427},
				Vector{0.3658854, 0.9776786},
			},
			Advance: 1.1232626,
		},
		'ì': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.37109375, 0.91071427},
				Vector{0.4375, 0.9776786},
			},
			Advance: 1.1232626,
		},
		'í': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.453125, 0.69977677},
				Vector{0.51953125, 0.7667411},
			},
			Advance: 1.1232626,
		},
		'î': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.44270834, 0.77120537},
				Vector{0.50911456, 0.83816963},
			},
			Advance: 1.1232626,
		},
		'ï': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.44270834, 0.8426339},
				Vector{0.50911456, 0.90959823},
			},
			Advance: 1.1232626,
		},
		'ð': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.44270834, 0.9140625},
				Vector{0.50911456, 0.98102677},
			},
			Advance: 1.1232626,
		},
		'ñ': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.53515625, 0.63504463},
				Vector{0.6015625, 0.7020089},
			},
			Advance: 1.1232626,
		},
		'ò': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.6067708, 0.63839287},
				Vector{0.67317706, 0.70535713},
			},
			Advance: 1.1232626,
		},
		'ó': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.67838544, 0.63839287},
				Vector{0.7447917, 0.70535713},
			},
			Advance: 1.1232626,
		},
		'ô': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.75, 0.63839287},
				Vector{0.81640625, 0.70535713},
			},
			Advance: 1.1232626,
		},
		'õ': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.82161456, 0.63839287},
				Vector{0.8880208, 0.70535713},
			},
			Advance: 1.1232626,
		},
		'ö': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.52473956, 0.70647323},
				Vector{0.5911458, 0.7734375},
			},
			Advance: 1.1232626,
		},
		'÷': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.51432294, 0.77790177},
				Vector{0.5807292, 0.8448661},
			},
			Advance: 1.1232626,
		},
		'ø': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.51432294, 0.84933037},
				Vector{0.5807292, 0.91629463},
			},
			Advance: 1.1232626,
		},
		'ù': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.51432294, 0.9207589},
				Vector{0.5807292, 0.98772323},
			},
			Advance: 1.1232626,
		},
		'ú': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5963542, 0.7098214},
				Vector{0.66276044, 0.77678573},
			},
			Advance: 1.1232626,
		},
		'û': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.66796875, 0.7098214},
				Vector{0.734375, 0.77678573},
			},
			Advance: 1.1232626,
		},
		'ü': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.7395833, 0.7098214},
				Vector{0.80598956, 0.77678573},
			},
			Advance: 1.1232626,
		},
		'ý': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.81119794, 0.7098214},
				Vector{0.8776042, 0.77678573},
			},
			Advance: 1.1232626,
		},
		'þ': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5859375, 0.78125},
				Vector{0.65234375, 0.84821427},
			},
			Advance: 1.1232626,
		},
		'ÿ': {
			Bounds: Rectangle{
				Vector{0.12480696, -1.0162853},
				Vector{0.9984557, 0.017829565},
			},
			UV: Rectangle{
				Vector{0.5859375, 0.8526786},
				Vector{0.65234375, 0.91964287},
			},
			Advance: 1.1232626,
		},
	},
	Kern: map[[2]rune]float32{},
}
