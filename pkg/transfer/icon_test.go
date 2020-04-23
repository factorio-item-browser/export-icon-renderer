package transfer

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	serializedIcon := `{"id":"9192f864-9953-4b85-8420-000b8aa2b8c7","size":42,"layers":[{"fileName":"abc","size":21,"scale":13.37},{"fileName":"def","size":27,"scale":2.7,"offset":{"x":12,"y":34},"tint":{"red":0.12,"green":0.34,"blue":0.56,"alpha":0.78}}]}`

	expectedIcon := Icon{
		Id:   "9192f864-9953-4b85-8420-000b8aa2b8c7",
		Size: 42,
		Layers: []Layer{
			{
				FileName: "abc",
				Size:     21,
				Scale:    13.37,
			},
			{
				FileName: "def",
				Size:     27,
				Scale:    2.7,
				Offset: Offset{
					X: 12,
					Y: 34,
				},
				Tint: Color{
					Red:   0.12,
					Green: 0.34,
					Blue:  0.56,
					Alpha: 0.78,
				},
			},
		},
	}

	var icon Icon
	err := json.Unmarshal([]byte(serializedIcon), &icon)

	assert.Nil(t, err)
	assert.Equal(t, expectedIcon, icon)
}
