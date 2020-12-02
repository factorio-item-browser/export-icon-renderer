package log

import (
	"errors"
)

type logFieldsProvider interface {
	LogFields() map[string]interface{}
}

func Error(err error) {
	logger.Error().Fields(extractFields(err)).Msg(err.Error())
}

func extractFields(err error) map[string]interface{} {
	fields := make(map[string]interface{})
	for err != nil {
		if err, ok := err.(logFieldsProvider); ok {
			for k, v := range err.LogFields() {
				if _, ok := fields[k]; !ok {
					fields[k] = v
				}
			}
		}

		err = errors.Unwrap(err)
	}
	return fields
}
