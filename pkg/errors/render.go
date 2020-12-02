package errors

import "fmt"

type RenderIconError struct {
	iconId        string
	originalError error
}

func NewRenderIconError(iconId string, originalError error) error {
	return &RenderIconError{
		iconId:        iconId,
		originalError: originalError,
	}
}

func (e *RenderIconError) Error() string {
	return fmt.Sprintf("render icon error: %s", e.originalError)
}

func (e *RenderIconError) Unwrap() error {
	return e.originalError
}

func (e *RenderIconError) LogFields() map[string]interface{} {
	return map[string]interface{}{
		"icon": e.iconId,
	}
}
