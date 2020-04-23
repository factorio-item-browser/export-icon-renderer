package command

import (
	"encoding/json"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/filter"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/io"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
)

type RenderIcon struct {
	layerFilters []filter.LayerFilter
	blendFilter  filter.BlendFilter
	resizeFilter filter.LayerFilter
}

func NewRenderIcon() *RenderIcon {
	return &RenderIcon{
		layerFilters: []filter.LayerFilter{
			filter.RemoveMipMaps,
			filter.Scale,
			filter.Offset,
			filter.Expand,
			filter.Crop,
		},
		blendFilter:  filter.TintedBlendFilter,
		resizeFilter: filter.ResizeToOutput,
	}
}

func (c *RenderIcon) Run(serializedIcon string) (string, error) {
	var icon transfer.Icon
	err := json.Unmarshal([]byte(serializedIcon), &icon)
	if err != nil {
		return "", err
	}

	props := createProperties(icon)

	finalImage := io.Create(props.RenderedSize)
	for _, layer := range icon.Layers {
		layerImage, err := io.Load(layer.FileName)
		if err != nil {
			return "", err
		}

		for _, layerFilter := range c.layerFilters {
			layerImage = layerFilter(layerImage, layer, props)
		}

		finalImage = c.blendFilter(finalImage, layerImage, layer, props)
	}
	finalImage = c.resizeFilter(finalImage, transfer.Layer{}, props)

	return io.Encode(finalImage)
}

func createProperties(icon transfer.Icon) filter.Properties {
	renderedSize := icon.Size
	if len(icon.Layers) > 0 {
		renderedSize = int(float64(icon.Layers[0].Size) * icon.Layers[0].Scale)
	}

	renderedScale := 1.
	if renderedSize < icon.Size {
		// We will enlarge the icon in the end, so render it bigger to begin with.
		renderedScale = float64(icon.Size) / float64(renderedSize)
		renderedSize = icon.Size
	}

	return filter.Properties{
		RenderedSize:  renderedSize,
		RenderedScale: renderedScale,
		OutputSize:    icon.Size,
	}
}
