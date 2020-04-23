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
		blendFilter: filter.TintedBlendFilter,
	}
}

func (c *RenderIcon) Run(serializedIcon string) (string, error) {
	var icon transfer.Icon
	err := json.Unmarshal([]byte(serializedIcon), &icon)
	if err != nil {
		return "", err
	}

	finalImage := io.Create(icon.Size)
	for _, layer := range icon.Layers {
		layerImage, err := io.Load(layer.FileName)
		if err != nil {
			return "", err
		}

		for _, layerFilter := range c.layerFilters {
			layerImage = layerFilter(layerImage, layer, icon)
		}

		finalImage = c.blendFilter(finalImage, layerImage, layer, icon)
	}

	return io.Encode(finalImage)
}
