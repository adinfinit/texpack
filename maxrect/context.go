package maxrect

import "image"

const maxInt = int(^uint(0) >> 1)

type Context struct {
	Size   image.Point
	Rotate bool
	Score  func(size image.Point) (image.Rectangle, int, int)

	Used []image.Rectangle
	Free []image.Rectangle

	DebugPlace func(r image.Rectangle)
}

func New(size image.Point) *Context {
	context := &Context{}
	context.Size = size
	context.Score = context.BottomLeftRule
	context.Free = append(context.Free, image.Rect(0, 0, size.X, size.Y))
	return context
}

func (context *Context) Add(size image.Point) (r image.Rectangle, ok bool) {
	node, _, _ := context.Score(size)
	if node == image.ZR {
		return node, false
	}
	context.PlaceRect(node)
	return node, true
}

func (context *Context) Adds(size ...image.Point) (r []image.Rectangle, ok bool) {
	r = make([]image.Rectangle, len(size))

	toAdd := make([]int, len(size))
	for i := range toAdd {
		toAdd[i] = i
	}

	for len(toAdd) > 0 {
		best, bestMajor, bestMinor, bestIndex := image.Rectangle{}, maxInt, maxInt, -1
		for pi, i := range toAdd {
			node, major, minor := context.Score(size[i])
			if node == image.ZR {
				// TODO: prune images that cannot fit immediately
				continue
			}
			if major < bestMajor || (major == bestMajor && minor < bestMinor) {
				best = node
				bestIndex = pi
				bestMajor = major
				bestMinor = minor
			}
		}

		if bestIndex == -1 {
			return r, false
		}

		context.PlaceRect(best)
		r[toAdd[bestIndex]] = best
		toAdd = append(toAdd[:bestIndex], toAdd[bestIndex+1:]...)
	}

	return r, true
}

func (context *Context) PlaceRect(rect image.Rectangle) {
	n := len(context.Free)
	for i := 0; i < n; i++ {
		if context.splitFree(context.Free[i], &rect) {
			context.Free = removeAt(context.Free, i)
			i--
			n--
		}
	}
	context.mergeFree()
	context.removeSmall(3)

	if context.DebugPlace != nil {
		context.DebugPlace(rect)
	}

	context.Used = append(context.Used, rect)
}

func removeAt(rects []image.Rectangle, i int) []image.Rectangle {
	return append(rects[:i], rects[i+1:]...)
}

// split free node if it overlaps the used rectangle
func (context *Context) splitFree(free image.Rectangle, used *image.Rectangle) bool {
	if !free.Overlaps(*used) {
		return false
	}

	if used.Min.X < free.Max.X && used.Max.X > free.Min.X {
		// New node at the top side of the used node.
		if used.Min.Y > free.Min.Y && used.Min.Y < free.Max.Y {
			split := free
			split.Max.Y = used.Min.Y
			context.Free = append(context.Free, split)
		}

		// New node at the bottom side of the used node.
		if used.Max.Y < free.Max.Y {
			split := free
			split.Min.Y = used.Max.Y
			context.Free = append(context.Free, split)
		}
	}

	if used.Min.Y < free.Max.Y && used.Max.Y > free.Min.Y {
		// New node at the left side of the used node.
		if used.Min.X > free.Min.X && used.Min.X < free.Max.X {
			split := free
			split.Max.X = used.Min.X
			context.Free = append(context.Free, split)
		}
		// New node at the right side of the used node.
		if used.Max.X < free.Max.X {
			split := free
			split.Min.X = used.Max.X
			context.Free = append(context.Free, split)
		}
	}
	return true
}

// mergeFree removes rects that are fully contained in another rect
func (context *Context) mergeFree() {
	for i := 0; i < len(context.Free); i++ {
		for k := i + 1; k < len(context.Free); k++ {
			if context.Free[i].In(context.Free[k]) {
				context.Free = removeAt(context.Free, i)
				i--
				break
			}
			if context.Free[k].In(context.Free[i]) {
				context.Free = removeAt(context.Free, k)
				k--
			}
		}
	}
}

// removeSmall removes rects that are unlikely to fit a rect
func (context *Context) removeSmall(limit int) {
	for i := 0; i < len(context.Free); i++ {
		sz := context.Free[i].Size()
		if sz.X <= limit || sz.Y <= limit {
			context.Free = removeAt(context.Free, i)
			i--
		}
	}
}
