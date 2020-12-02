package filter

import (
	"github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPropertiesFromIcon(t *testing.T) {
	tests := []struct {
		name           string
		icon           transfer.Icon
		expectedResult Properties
	}{
		{
			name: "no layers",
			icon: transfer.Icon{
				Size: 64,
			},
			expectedResult: Properties{
				RenderedSize:  64,
				RenderedScale: 1,
				OutputSize:    64,
			},
		},
		{
			name: "large first layer",
			icon: transfer.Icon{
				Size: 64,
				Layers: []transfer.Layer{
					{
						Scale: 1.5,
						Size:  128,
					},
				},
			},
			expectedResult: Properties{
				RenderedSize:  192,
				RenderedScale: 1,
				OutputSize:    64,
			},
		},
		{
			name: "small first layer",
			icon: transfer.Icon{
				Size: 64,
				Layers: []transfer.Layer{
					{
						Scale: 0.5,
						Size:  32,
					},
				},
			},
			expectedResult: Properties{
				RenderedSize:  64,
				RenderedScale: 4,
				OutputSize:    64,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := NewPropertiesFromIcon(test.icon)
			assert.Equal(t, test.expectedResult, result)
		})
	}
}
