package main

type Atlas struct {
	Meta   Meta             `json:"meta"`
	Frames map[string]Frame `json:"frames"`
}

type Rect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"w"`
	Height int `json:"h"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Size struct {
	Width  int `json:"w"`
	Height int `json:"h"`
}

type Meta struct {
	App     string `json:"app"`
	Version string `json:"version"`
	Image   string `json:"image"`
	Format  string `json:"format"`
	Size    Size   `json:"size"`
}

type Frame struct {
	Name             string `json:"name"`
	Rotated          bool   `json:"rotated"`
	Trimmed          bool   `json:"trimmed"`
	Frame            Rect   `json:"frame"`
	SpriteSourceSize Rect   `json:"spriteSourceSize"`
	SourceSize       Size   `json:"sourceSize"`
	Pivot            Point  `json:"pivot"`
}
