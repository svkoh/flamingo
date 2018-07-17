// Code generated by mockery v1.0.0
package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"

import web "flamingo.me/flamingo/framework/web"

// ErrorAware is an autogenerated mock type for the ErrorAware type
type ErrorAware struct {
	mock.Mock
}

// Error provides a mock function with given fields: _a0, err
func (_m *ErrorAware) Error(_a0 context.Context, err error) web.Response {
	ret := _m.Called(_a0, err)

	var r0 web.Response
	if rf, ok := ret.Get(0).(func(context.Context, error) web.Response); ok {
		r0 = rf(_a0, err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(web.Response)
		}
	}

	return r0
}

// ErrorNotFound provides a mock function with given fields: _a0, err
func (_m *ErrorAware) ErrorNotFound(_a0 context.Context, err error) web.Response {
	ret := _m.Called(_a0, err)

	var r0 web.Response
	if rf, ok := ret.Get(0).(func(context.Context, error) web.Response); ok {
		r0 = rf(_a0, err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(web.Response)
		}
	}

	return r0
}

// ErrorWithCode provides a mock function with given fields: _a0, err, httpStatus
func (_m *ErrorAware) ErrorWithCode(_a0 context.Context, err error, httpStatus int) web.Response {
	ret := _m.Called(_a0, err, httpStatus)

	var r0 web.Response
	if rf, ok := ret.Get(0).(func(context.Context, error, int) web.Response); ok {
		r0 = rf(_a0, err, httpStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(web.Response)
		}
	}

	return r0
}
