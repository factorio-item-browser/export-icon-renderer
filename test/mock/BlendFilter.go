// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	image "image"

	filter "github.com/factorio-item-browser/export-icon-renderer/pkg/filter"

	mock "github.com/stretchr/testify/mock"

	transfer "github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
)

// BlendFilter is an autogenerated mock type for the BlendFilter type
type BlendFilter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: destination, source, layer, props
func (_m *BlendFilter) Execute(destination image.Image, source image.Image, layer transfer.Layer, props filter.Properties) image.Image {
	ret := _m.Called(destination, source, layer, props)

	var r0 image.Image
	if rf, ok := ret.Get(0).(func(image.Image, image.Image, transfer.Layer, filter.Properties) image.Image); ok {
		r0 = rf(destination, source, layer, props)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(image.Image)
		}
	}

	return r0
}
