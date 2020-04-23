package filter

import (
	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/transform"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
	"image"
)

// RemoveMipMaps is the filter for removing the mip maps of the source, only keeping the highest resolution of it.
func RemoveMipMaps(source image.Image, layer transfer.Layer, icon transfer.Icon) image.Image {
	return transform.Crop(source, image.Rect(0, 0, layer.Size, layer.Size))
}

// Scale applies the scaling value of the layer to the source.
func Scale(source image.Image, layer transfer.Layer, icon transfer.Icon) image.Image {
	if layer.Scale == 1. {
		return source
	}

	newSize := int(float64(layer.Size) * layer.Scale)
	return transform.Resize(source, newSize, newSize, transform.Lanczos)
}

// Offset applies the offset values of the layer to the source.
func Offset(source image.Image, layer transfer.Layer, icon transfer.Icon) image.Image {
	if layer.Offset.X == 0 && layer.Offset.Y == 0 {
		return source
	}

	size := int(float64(layer.Size) * layer.Scale)
	source = clone.Pad(source, size, size, clone.NoFill)
	source = transform.Translate(source, layer.Offset.X, layer.Offset.Y)
	return source
}

// Expand applies additional padding to the source so it definitively exceeds the desired size of the icon.
// Note that the image will most likely be too large afterwards.
func Expand(source image.Image, layer transfer.Layer, icon transfer.Icon) image.Image {
	if source.Bounds().Size().X >= icon.Size {
		return source
	}

	return clone.Pad(source, icon.Size, icon.Size, clone.NoFill)
}

// Crop will cut out the middle of the source to get the desired size of the icon.
func Crop(source image.Image, layer transfer.Layer, icon transfer.Icon) image.Image {
	if source.Bounds().Size().X <= icon.Size {
		return source
	}

	position := (source.Bounds().Size().X - icon.Size) / 2
	return transform.Crop(source, image.Rect(position, position, position+icon.Size, position+icon.Size))
}
