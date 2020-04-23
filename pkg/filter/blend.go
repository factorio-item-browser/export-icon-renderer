package filter

import (
	"github.com/anthonynsimon/bild/blend"
	"github.com/anthonynsimon/bild/fcolor"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
	"image"
)

// TintedBlendFilter blends the source onto the destination while applying the tint of the layer.
func TintedBlendFilter(destination, source image.Image, layer transfer.Layer, props Properties) image.Image {
	tint := fcolor.RGBAF64{
		R: layer.Tint.Red,
		G: layer.Tint.Green,
		B: layer.Tint.Blue,
		A: layer.Tint.Alpha,
	}
	tint.Clamp()

	return blend.Blend(destination, source, func(dst, src fcolor.RGBAF64) fcolor.RGBAF64 {
		color := fcolor.RGBAF64{
			R: src.R*tint.R*src.A + dst.R*dst.A*(1-src.A*tint.A),
			G: src.G*tint.G*src.A + dst.G*dst.A*(1-src.A*tint.A),
			B: src.B*tint.B*src.A + dst.B*dst.A*(1-src.A*tint.A),
			A: src.A + dst.A*(1-src.A*tint.A),
		}
		color.Clamp()
		return color
	})
}
