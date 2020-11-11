package io

import (
	"github.com/factorio-item-browser/export-icon-renderer/pkg/env"
	"github.com/stretchr/testify/assert"
	"image"
	"io/ioutil"
	"testing"
)

func TestCreate(t *testing.T) {
	size := 2

	expectedResult := image.RGBA{
		Pix: []uint8{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
		Stride: 8,
		Rect:   image.Rect(0, 0, 2, 2),
	}

	result := Create(size)

	assert.Equal(t, &expectedResult, result)
}

func TestLoad(t *testing.T) {
	defer func(ref1, ref2 string) {
		env.FactorioDataDirectory = ref1
		env.FactorioModsDirectory = ref2
	}(env.FactorioDataDirectory, env.FactorioModsDirectory)
	env.FactorioDataDirectory = "../../test/asset/factorio/data"
	env.FactorioModsDirectory = "../../test/asset/factorio/mods"

	tests := []struct {
		name           string
		fileName       string
		expectedResult image.Image
	}{
		{
			name:     "icon from actual mod",
			fileName: "__bar__/graphics/icon.png",
			expectedResult: &image.RGBA{
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
		},
		{
			name:     "icon from base mod",
			fileName: "__base__/graphics/icon.png",
			expectedResult: &image.RGBA{
				Pix: []uint8{
					0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF,
				},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
		},
		{
			name:     "icon from core mod",
			fileName: "__core__/graphics/icon.png",
			expectedResult: &image.RGBA{
				Pix: []uint8{
					0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF,
					0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF,
				},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Load(test.fileName)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedResult, result)
		})
	}
}

func TestEncode(t *testing.T) {
	img := image.RGBA{
		Pix: []uint8{
			0x00, 0x10, 0x20, 0x30, 0x40, 0x50, 0x60, 0x70,
			0x80, 0x90, 0xA0, 0xB0, 0xC0, 0xD0, 0xE0, 0xF0,
		},
		Stride: 8,
		Rect:   image.Rect(0, 0, 2, 2),
	}

	expectedResult, _ := ioutil.ReadFile("../../test/asset/encode.png")

	result, err := Encode(&img)

	assert.Nil(t, err)
	assert.Equal(t, string(expectedResult), result)
}
