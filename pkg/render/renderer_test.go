package render

import (
	"fmt"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/errors"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/filter"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
	mocks "github.com/factorio-item-browser/export-icon-renderer/test/mock"
	"github.com/stretchr/testify/assert"
	"image"
	"testing"
)

func TestNewRenderer(t *testing.T) {
	result := NewRenderer()

	assert.NotNil(t, result)
}

func TestRenderer_Render(t *testing.T) {
	defer func(
		ref1 func(int) image.Image,
		ref2 func(image.Image) (string, error),
		ref3 func(string) (image.Image, error),
	) {
		create = ref1
		encode = ref2
		load = ref3
	}(create, encode, load)

	layer1 := transfer.Layer{
		FileName: "abc",
		Size:     64,
		Scale:    1,
	}
	layer2 := transfer.Layer{
		FileName: "def",
	}
	icon := transfer.Icon{
		Size:   64,
		Layers: []transfer.Layer{layer1, layer2},
	}
	expectedProperties := filter.Properties{
		RenderedSize:  64,
		RenderedScale: 1,
		OutputSize:    64,
	}
	encodedImage := "ghi"

	image1 := image.NewRGBA(image.Rect(0, 0, 1, 1))
	layer1a := image.NewRGBA(image.Rect(0, 0, 1, 2))
	layer1b := image.NewRGBA(image.Rect(0, 0, 1, 3))
	layer1c := image.NewRGBA(image.Rect(0, 0, 1, 4))

	image2 := image.NewRGBA(image.Rect(0, 0, 2, 1))
	layer2a := image.NewRGBA(image.Rect(0, 0, 2, 2))
	layer2b := image.NewRGBA(image.Rect(0, 0, 2, 3))
	layer2c := image.NewRGBA(image.Rect(0, 0, 2, 4))

	image3 := image.NewRGBA(image.Rect(0, 0, 3, 1))
	image4 := image.NewRGBA(image.Rect(0, 0, 4, 1))

	create = func(s int) image.Image {
		assert.Exactly(t, 64, s)
		return image1
	}
	load = func(f string) (image.Image, error) {
		load = func(f string) (image.Image, error) {
			assert.Exactly(t, "def", f)
			return layer2a, nil
		}
		assert.Exactly(t, "abc", f)
		return layer1a, nil
	}
	encode = func(img image.Image) (string, error) {
		assert.Exactly(t, image4, img)
		return encodedImage, nil
	}

	layerFilter1 := mocks.LayerFilter{}
	layerFilter1.On("Execute", layer1a, layer1, expectedProperties).Return(layer1b)
	layerFilter1.On("Execute", layer2a, layer2, expectedProperties).Return(layer2b)

	layerFilter2 := mocks.LayerFilter{}
	layerFilter2.On("Execute", layer1b, layer1, expectedProperties).Return(layer1c)
	layerFilter2.On("Execute", layer2b, layer2, expectedProperties).Return(layer2c)

	blendFilter := mocks.BlendFilter{}
	blendFilter.On("Execute", image1, layer1c, layer1, expectedProperties).Return(image2)
	blendFilter.On("Execute", image2, layer2c, layer2, expectedProperties).Return(image3)

	resizeFilter := mocks.LayerFilter{}
	resizeFilter.On("Execute", image3, transfer.Layer{}, expectedProperties).Return(image4)

	renderer := Renderer{
		layerFilters: []filter.LayerFilter{layerFilter1.Execute, layerFilter2.Execute},
		blendFilter:  blendFilter.Execute,
		resizeFilter: resizeFilter.Execute,
	}

	result, err := renderer.Render(icon)

	assert.Nil(t, err)
	assert.Exactly(t, encodedImage, result)
}

func TestRenderer_Render_WithLoadError(t *testing.T) {
	defer func(
		ref1 func(int) image.Image,
		ref3 func(string) (image.Image, error),
	) {
		create = ref1
		load = ref3
	}(create, load)

	layer1 := transfer.Layer{
		FileName: "abc",
		Size:     64,
		Scale:    1,
	}
	layer2 := transfer.Layer{
		FileName: "def",
	}
	icon := transfer.Icon{
		Size:   64,
		Layers: []transfer.Layer{layer1, layer2},
	}

	image1 := image.NewRGBA(image.Rect(0, 0, 1, 1))

	create = func(s int) image.Image {
		assert.Exactly(t, 64, s)
		return image1
	}
	load = func(f string) (image.Image, error) {
		assert.Exactly(t, "abc", f)
		return nil, fmt.Errorf("test error")
	}

	layerFilter1 := mocks.LayerFilter{}
	layerFilter2 := mocks.LayerFilter{}
	blendFilter := mocks.BlendFilter{}
	resizeFilter := mocks.LayerFilter{}

	renderer := Renderer{
		layerFilters: []filter.LayerFilter{layerFilter1.Execute, layerFilter2.Execute},
		blendFilter:  blendFilter.Execute,
		resizeFilter: resizeFilter.Execute,
	}

	_, err := renderer.Render(icon)

	assert.IsType(t, &errors.RenderIconError{}, err)
}

func TestRenderer_Render_WithEncodeError(t *testing.T) {
	defer func(
		ref1 func(int) image.Image,
		ref2 func(image.Image) (string, error),
		ref3 func(string) (image.Image, error),
	) {
		create = ref1
		encode = ref2
		load = ref3
	}(create, encode, load)

	layer1 := transfer.Layer{
		FileName: "abc",
		Size:     64,
		Scale:    1,
	}
	layer2 := transfer.Layer{
		FileName: "def",
	}
	icon := transfer.Icon{
		Size:   64,
		Layers: []transfer.Layer{layer1, layer2},
	}
	expectedProperties := filter.Properties{
		RenderedSize:  64,
		RenderedScale: 1,
		OutputSize:    64,
	}

	image1 := image.NewRGBA(image.Rect(0, 0, 1, 1))
	layer1a := image.NewRGBA(image.Rect(0, 0, 1, 2))
	layer1b := image.NewRGBA(image.Rect(0, 0, 1, 3))
	layer1c := image.NewRGBA(image.Rect(0, 0, 1, 4))

	image2 := image.NewRGBA(image.Rect(0, 0, 2, 1))
	layer2a := image.NewRGBA(image.Rect(0, 0, 2, 2))
	layer2b := image.NewRGBA(image.Rect(0, 0, 2, 3))
	layer2c := image.NewRGBA(image.Rect(0, 0, 2, 4))

	image3 := image.NewRGBA(image.Rect(0, 0, 3, 1))
	image4 := image.NewRGBA(image.Rect(0, 0, 4, 1))

	create = func(s int) image.Image {
		assert.Exactly(t, 64, s)
		return image1
	}
	load = func(f string) (image.Image, error) {
		load = func(f string) (image.Image, error) {
			assert.Exactly(t, "def", f)
			return layer2a, nil
		}
		assert.Exactly(t, "abc", f)
		return layer1a, nil
	}
	encode = func(img image.Image) (string, error) {
		assert.Exactly(t, image4, img)
		return "", fmt.Errorf("test error")
	}

	layerFilter1 := mocks.LayerFilter{}
	layerFilter1.On("Execute", layer1a, layer1, expectedProperties).Return(layer1b)
	layerFilter1.On("Execute", layer2a, layer2, expectedProperties).Return(layer2b)

	layerFilter2 := mocks.LayerFilter{}
	layerFilter2.On("Execute", layer1b, layer1, expectedProperties).Return(layer1c)
	layerFilter2.On("Execute", layer2b, layer2, expectedProperties).Return(layer2c)

	blendFilter := mocks.BlendFilter{}
	blendFilter.On("Execute", image1, layer1c, layer1, expectedProperties).Return(image2)
	blendFilter.On("Execute", image2, layer2c, layer2, expectedProperties).Return(image3)

	resizeFilter := mocks.LayerFilter{}
	resizeFilter.On("Execute", image3, transfer.Layer{}, expectedProperties).Return(image4)

	renderer := Renderer{
		layerFilters: []filter.LayerFilter{layerFilter1.Execute, layerFilter2.Execute},
		blendFilter:  blendFilter.Execute,
		resizeFilter: resizeFilter.Execute,
	}

	_, err := renderer.Render(icon)

	assert.IsType(t, &errors.RenderIconError{}, err)
}
