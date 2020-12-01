package errors

import "fmt"

type InvalidFileNameError struct {
	fileName string
}

func NewInvalidFileNameError(fileName string) error {
	return &InvalidFileNameError{
		fileName: fileName,
	}
}

func (e *InvalidFileNameError) Error() string {
	return "unable to parse file name"
}

func (e *InvalidFileNameError) LogFields() map[string]interface{} {
	return map[string]interface{}{
		"file": e.fileName,
	}
}

type ImageLoadError struct {
	fileName      string
	originalError error
}

func NewImageLoadError(fileName string, originalError error) error {
	return &ImageLoadError{
		fileName:      fileName,
		originalError: originalError,
	}
}

func (e *ImageLoadError) Error() string {
	return fmt.Sprintf("unable to load image: %s", e.originalError)
}

func (e *ImageLoadError) Unwrap() error {
	return e.originalError
}

func (e *ImageLoadError) LogFields() map[string]interface{} {
	return map[string]interface{}{
		"file": e.fileName,
	}
}

type EncodingError struct {
	originalError error
}

func NewEncodingError(originalError error) error {
	return &EncodingError{
		originalError: originalError,
	}
}

func (e *EncodingError) Error() string {
	return fmt.Sprintf("unable to encode image: %s", e.originalError)
}

func (e *EncodingError) Unwrap() error {
	return e.originalError
}
