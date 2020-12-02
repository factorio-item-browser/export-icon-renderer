package filter

import "github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"

// Properties are additional values needed to render the icon.
type Properties struct {
	RenderedSize  int
	RenderedScale float64
	OutputSize    int
}

// NewPropertiesFromIcon creates the rendering properties for the specified icon.
func NewPropertiesFromIcon(icon transfer.Icon) Properties {
	renderedSize := icon.Size
	if len(icon.Layers) > 0 {
		renderedSize = int(float64(icon.Layers[0].Size) * icon.Layers[0].Scale)
	}

	renderedScale := 1.
	if renderedSize < icon.Size {
		// We would enlarge the icon in the end, so render it bigger to begin with.
		renderedScale = float64(icon.Size) / float64(renderedSize)
		renderedSize = icon.Size
	}

	return Properties{
		RenderedSize:  renderedSize,
		RenderedScale: renderedScale,
		OutputSize:    icon.Size,
	}
}
