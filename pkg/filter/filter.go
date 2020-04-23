package filter

import (
	"github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
	"image"
)

// BlendFilter blends the current layer image onto the image which has been generated so far.
type BlendFilter func(destination, source image.Image, layer transfer.Layer, props Properties) image.Image

// LayerFilter manipulates the current layer image of the icon.
type LayerFilter func(source image.Image, layer transfer.Layer, props Properties) image.Image

// Properties are additional values needed to render the icon.
type Properties struct {
	RenderedSize  int
	RenderedScale float64
	OutputSize    int
}
