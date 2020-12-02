package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidFileNameError(t *testing.T) {
	fileName := "abc"
	expectedMessage := "unable to parse file name"
	expectedLogFields := map[string]interface{}{
		"file": "abc",
	}

	err := NewInvalidFileNameError(fileName)

	assert.IsType(t, &InvalidFileNameError{}, err)
	assert.Exactly(t, expectedMessage, err.Error())
	assert.Exactly(t, expectedLogFields, err.(*InvalidFileNameError).LogFields())
}

func TestImageLoadError(t *testing.T) {
	fileName := "abc"
	originalError := fmt.Errorf("test error")
	expectedMessage := "unable to load image: test error"
	expectedLogFields := map[string]interface{}{
		"file": "abc",
	}

	err := NewImageLoadError(fileName, originalError)

	assert.IsType(t, &ImageLoadError{}, err)
	assert.Exactly(t, expectedMessage, err.Error())
	assert.Exactly(t, originalError, errors.Unwrap(err))
	assert.Exactly(t, expectedLogFields, err.(*ImageLoadError).LogFields())
}

func TestEncodingError(t *testing.T) {
	originalError := fmt.Errorf("test error")
	expectedMessage := "unable to encode image: test error"

	err := NewEncodingError(originalError)

	assert.IsType(t, &EncodingError{}, err)
	assert.Exactly(t, expectedMessage, err.Error())
	assert.Exactly(t, originalError, errors.Unwrap(err))
}
