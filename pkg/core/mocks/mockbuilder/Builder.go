// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockbuilder

import core "github.com/projectriff/riff/pkg/core"
import io "io"
import mock "github.com/stretchr/testify/mock"

// Builder is an autogenerated mock type for the Builder type
type Builder struct {
	mock.Mock
}

// Build provides a mock function with given fields: repoName, options, log
func (_m *Builder) Build(repoName string, options core.BuildOptions, log io.Writer) error {
	ret := _m.Called(repoName, options, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, core.BuildOptions, io.Writer) error); ok {
		r0 = rf(repoName, options, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
