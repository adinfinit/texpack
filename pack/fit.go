// pack implements packing common things
package pack

import (
	"image"
	"sort"

	"github.com/loov/texpack/maxrect"
)

type Box struct {
	Padding int
	Size    image.Point
	Place   *image.Rectangle
}

func SortBoxes(boxes []Box) {
	sort.Slice(boxes, func(i, k int) bool {
		a, b := boxes[i], boxes[k]
		if a.Size.X > b.Size.X {
			return true
		} else if a.Size.X < b.Size.X {
			return false
		}
		if a.Size.Y > b.Size.Y {
			return true
		} else if a.Size.Y < b.Size.Y {
			return false
		}
		return false
	})
}

func BoxesToSizes(boxes []Box) []image.Point {
	sizes := []image.Point{}
	for _, b := range boxes {
		sizes = append(sizes, b.Size.Add(image.Point{b.Padding * 2, b.Padding * 2}))
	}
	return sizes
}

func PlaceBoxes(maxContextSize image.Point, rule maxrect.Rule, boxes []Box) (contextSize image.Point, ok bool) {
	sizes := BoxesToSizes(boxes)

	var rects []image.Rectangle
	contextSize, rects, ok = minimizeFit(maxContextSize, rule, sizes)
	for i := range rects {
		box := &boxes[i]
		*box.Place = rects[i].Inset(box.Padding)
	}

	return
}

func minimizeFit(maxContextSize image.Point, rule maxrect.Rule, sizes []image.Point) (contextSize image.Point, rects []image.Rectangle, ok bool) {

	try := func(size image.Point) ([]image.Rectangle, bool) {
		context := maxrect.New(size)
		context.SetRule(rule)
		return context.Adds(sizes...)
	}

	contextSize = maxContextSize
	rects, ok = try(contextSize)
	if !ok {
		return
	}

	shrunk, shrinkX, shrinkY := true, true, true
	for shrunk {
		shrunk = false
		if shrinkX {
			trySize := image.Point{contextSize.X - 128, contextSize.Y}
			tryRects, tryOk := try(trySize)
			if tryOk {
				contextSize = trySize
				rects = tryRects
				shrunk = true
			} else {
				shrinkX = false
			}
		}

		if shrinkY {
			trySize := image.Point{contextSize.X, contextSize.Y - 128}
			tryRects, tryOk := try(trySize)
			if tryOk {
				contextSize = trySize
				rects = tryRects
				shrunk = true
			} else {
				shrinkX = false
			}
		}
	}

	return
}
