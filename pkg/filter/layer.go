package filter

import (
	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/transform"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
	"image"
	"math"
)

var resizeAlgorithm = transform.MitchellNetravali

// RemoveMipMaps is the filter for removing the mip maps of the source, only keeping the highest resolution of it.
func RemoveMipMaps(source image.Image, layer transfer.Layer, props Properties) image.Image {
	return transform.Crop(source, image.Rect(0, 0, layer.Size, layer.Size))
}

// Scale applies the scaling value of the layer to the source.
func Scale(source image.Image, layer transfer.Layer, props Properties) image.Image {
	scale := layer.Scale * props.RenderedScale
	if scale == 1. {
		return source
	}

	newSize := int(math.Round(float64(layer.Size) * scale))
	return transform.Resize(source, newSize, newSize, resizeAlgorithm)
}

// Offset applies the offset values of the layer to the source.
func Offset(source image.Image, layer transfer.Layer, props Properties) image.Image {
	if layer.Offset.X == 0 && layer.Offset.Y == 0 {
		return source
	}

	dx := int(math.Round(float64(layer.Offset.X) * props.RenderedScale))
	dy := int(math.Round(float64(layer.Offset.Y) * props.RenderedScale))

	source = clone.Pad(source, props.RenderedSize, props.RenderedSize, clone.NoFill)
	source = transform.Translate(source, dx, -dy)
	return source
}

// Expand applies additional padding to the source so it definitively exceeds the desired size of the icon.
// Note that the image will most likely be too large afterwards.
func Expand(source image.Image, layer transfer.Layer, props Properties) image.Image {
	if source.Bounds().Size().X >= props.RenderedSize {
		return source
	}

	return clone.Pad(source, props.RenderedSize, props.RenderedSize, clone.NoFill)
}

// Crop will cut out the middle of the source to get the desired size of the icon.
func Crop(source image.Image, layer transfer.Layer, props Properties) image.Image {
	if source.Bounds().Size().X <= props.RenderedSize {
		return source
	}

	position := (source.Bounds().Size().X - props.RenderedSize) / 2
	return transform.Crop(source, image.Rect(position, position, position+props.RenderedSize, position+props.RenderedSize))
}

// ResizeToOutput will resize the source to the output size specified in the properties.
func ResizeToOutput(source image.Image, layer transfer.Layer, props Properties) image.Image {
	if source.Bounds().Size().X == props.OutputSize {
		return source
	}

	return transform.Resize(source, props.OutputSize, props.OutputSize, resizeAlgorithm)
}
