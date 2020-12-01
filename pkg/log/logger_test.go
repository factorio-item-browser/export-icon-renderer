package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogger(t *testing.T) {
	result1 := Logger()
	result2 := Logger()

	assert.Exactly(t, result1, result2)
}
