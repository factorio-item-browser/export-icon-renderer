package filter

import (
	"github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
	"github.com/stretchr/testify/assert"
	"image"
	"testing"
)

func TestTintedBlendFilter(t *testing.T) {
	layer := transfer.Layer{
		Tint: transfer.Color{
			Red:   .2,
			Green: .4,
			Blue:  .6,
			Alpha: .8,
		},
	}
	source := image.RGBA{
		Pix: []uint8{
			0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
			0xFF, 0x00, 0x00, 0x80, 0x80, 0x00, 0xFF, 0x80,
		},
		Stride: 8,
		Rect:   image.Rect(0, 0, 2, 2),
	}
	destination := image.RGBA{
		Pix: []uint8{
			0x00, 0x00, 0xFF, 0xFF, 0x00, 0x80, 0x00, 0xFF,
			0x00, 0x00, 0xFF, 0x80, 0x00, 0x80, 0x00, 0x80,
		},
		Stride: 8,
		Rect:   image.Rect(0, 0, 2, 2),
	}
	expectedResult := image.RGBA{
		Pix: []uint8{
			0x33, 0x00, 0x32, 0xFF, 0x19, 0x19, 0x00, 0xFF,
			0x19, 0x00, 0x4C, 0xCC, 0x0C, 0x26, 0x4C, 0xCC,
		},
		Stride: 8,
		Rect:   image.Rect(0, 0, 2, 2),
	}

	result := TintedBlendFilter(&destination, &source, layer, Properties{})

	assert.Equal(t, &expectedResult, result)
}
