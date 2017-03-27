package maxrect

import "image"

type Rule byte

const (
	Automatic Rule = iota
	ShortSide
	LongSide
	BottomLeft
	Area
	ContactPoint
)

func ParseRule(s string) Rule {
	switch s {
	case "short-side":
		return ShortSide
	case "long-side":
		return LongSide
	case "bottom-left":
		return BottomLeft
	case "area":
		return Area
	case "contact-point":
		return ContactPoint
	default:
		return Automatic
	}
}

func (context *Context) SetRule(rule Rule) {
	switch rule {
	case Automatic:
		context.Score = context.ContactPoint
	case ShortSide:
		context.Score = context.ShortSide
	case LongSide:
		context.Score = context.LongSide
	case BottomLeft:
		context.Score = context.BottomLeft
	case Area:
		context.Score = context.Area
	case ContactPoint:
		context.Score = context.ContactPoint
	default:
		panic("unknown rule")
	}
}

func (context *Context) ShortSide(size image.Point) (image.Rectangle, int, int) {
	return FindBest(context.Free, size, context.Rotate,
		func(free image.Rectangle, size image.Point) (int, int) {
			leftX, leftY := free.Dx()-size.X, free.Dy()-size.Y
			if leftX < leftY {
				return leftX, leftY
			} else {
				return leftY, leftX
			}
		})
}

func (context *Context) LongSide(size image.Point) (image.Rectangle, int, int) {
	return FindBest(context.Free, size, context.Rotate,
		func(free image.Rectangle, size image.Point) (int, int) {
			leftX, leftY := free.Dx()-size.X, free.Dy()-size.Y
			if leftX > leftY {
				return leftX, leftY
			} else {
				return leftY, leftX
			}
		})
}

func (context *Context) BottomLeft(size image.Point) (image.Rectangle, int, int) {
	return FindBest(context.Free, size, context.Rotate,
		func(free image.Rectangle, size image.Point) (int, int) {
			return free.Min.Y + size.Y, free.Min.X
		})
}

func (context *Context) Area(size image.Point) (image.Rectangle, int, int) {
	area := size.X * size.Y
	return FindBest(context.Free, size, context.Rotate,
		func(free image.Rectangle, size image.Point) (int, int) {
			areaFit := free.Dx()*free.Dy() - area
			leftX := free.Dx() - size.X
			leftY := free.Dy() - size.Y
			shortFit := min(leftX, leftY)
			return areaFit, shortFit
		})
}

func (context *Context) ContactPoint(size image.Point) (image.Rectangle, int, int) {
	return FindBest(context.Free, size, context.Rotate,
		func(free image.Rectangle, size image.Point) (int, int) {
			contact := maxInt
			target := image.Rectangle{free.Min, free.Min.Add(size)}
			// do we touch context sides
			if target.Min.X == 0 || target.Max.X == context.Size.X {
				contact -= size.X
			}
			if target.Min.Y == 0 || target.Max.Y == context.Size.Y {
				contact -= size.Y
			}
			for _, used := range context.Used {
				if used.Min.X == target.Max.X || used.Max.X == target.Min.X {
					contact -= CommonInterval(used.Min.Y, used.Max.Y, target.Min.Y, target.Max.Y)
				}
				if used.Min.Y == target.Max.Y || used.Max.Y == target.Min.Y {
					contact -= CommonInterval(used.Min.X, used.Max.X, target.Min.X, target.Max.X)
				}
			}
			return contact, contact
		})
}
