package filter

import (
	"github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
	"image"
)

// LayerFilter manipulates the current layer image of the icon.
type LayerFilter func(source image.Image, layer transfer.Layer, icon transfer.Icon) image.Image

// BlendFilter blends the current layer image onto the image which has been generated so far.
type BlendFilter func(destination, source image.Image, layer transfer.Layer, icon transfer.Icon) image.Image
