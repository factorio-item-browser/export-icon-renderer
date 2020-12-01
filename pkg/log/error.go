package log

import (
	"errors"
	"github.com/rs/zerolog"
)

type logFieldsProvider interface {
	LogFields() map[string]interface{}
}

func Error(event *zerolog.Event, err error) {
	event.Fields(extractFields(err)).Msg(err.Error())
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
