# texpack

`texpack` is a texture packer with both image and font support in the same atlas.

## Notes

* Does not support OpenType font kerning -- limitation of `github.com/golang/freetype`.
* SDF support is weak
* Doesn't auto rotate

## Installing

```
go get github.com/adinfinit/texpack
```

## sub-packages

* `maxrect` - is a MaxRects implementation
* `pack` - implements utilities for interfacing with maxrect
* `sdf` - implements signed-distance-field generation
