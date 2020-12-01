package log

import (
	"bytes"
	"fmt"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/errors"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	defer func(ref1 zerolog.Logger) {
		logger = ref1
	}(logger)

	buf := bytes.NewBufferString("")
	logger = zlog.Output(buf)

	error1 := fmt.Errorf("initial error")
	error2 := errors.NewImageLoadError("abc", error1)
	error3 := errors.NewRenderIconError("def", error2)

	Error(error3)

	result := buf.String()

	assert.Contains(t, result, `"file":"abc"`)
	assert.Contains(t, result, `"icon":"def"`)
	assert.Contains(t, result, "initial error")
}
