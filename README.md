# texpack

`texpack` is a texture packer with both image and font support in the same atlas.
It also has also signed-distance-field support.

## Installing

```
go get github.com/loov/texpack
```

## sub-packages

* `maxrect` - is a MaxRects implementation
* `pack` - implements utilities for interfacing with maxrect
* `sdf` - implements signed-distance-field generation

## Notes

* Does not support OpenType font kerning -- limitation of `github.com/golang/freetype`.
* SDF generation is not yet good