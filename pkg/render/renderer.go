package render

import (
	"github.com/factorio-item-browser/export-icon-renderer/pkg/errors"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/filter"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/io"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
)

var (
	create = io.Create
	encode = io.Encode
	load   = io.Load
)

type Renderer struct {
	layerFilters []filter.LayerFilter
	blendFilter  filter.BlendFilter
	resizeFilter filter.LayerFilter
}

func NewRenderer() *Renderer {
	return &Renderer{
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

func (r *Renderer) Render(icon transfer.Icon) (string, error) {
	props := filter.NewPropertiesFromIcon(icon)

	finalImage := create(props.RenderedSize)
	for _, layer := range icon.Layers {
		layerImage, err := load(layer.FileName)
		if err != nil {
			return "", errors.NewRenderIconError(icon.Id, err)
		}
		for _, layerFilter := range r.layerFilters {
			layerImage = layerFilter(layerImage, layer, props)
		}

		finalImage = r.blendFilter(finalImage, layerImage, layer, props)
	}
	finalImage = r.resizeFilter(finalImage, transfer.Layer{}, props)

	encodedImage, err := encode(finalImage)
	if err != nil {
		return "", errors.NewRenderIconError(icon.Id, err)
	}
	return encodedImage, nil
}
