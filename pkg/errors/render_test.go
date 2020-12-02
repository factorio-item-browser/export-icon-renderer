package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderIconError(t *testing.T) {
	iconId := "abc"
	originalError := fmt.Errorf("test error")
	expectedMessage := "render icon error: test error"
	expectedLogFields := map[string]interface{}{
		"icon": "abc",
	}

	err := NewRenderIconError(iconId, originalError)

	assert.IsType(t, &RenderIconError{}, err)
	assert.Exactly(t, expectedMessage, err.Error())
	assert.Exactly(t, originalError, errors.Unwrap(err))
	assert.Exactly(t, expectedLogFields, err.(*RenderIconError).LogFields())
}
