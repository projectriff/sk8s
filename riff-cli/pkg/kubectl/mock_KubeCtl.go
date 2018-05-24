// Code generated by mockery v1.0.0. DO NOT EDIT.
package kubectl

import mock "github.com/stretchr/testify/mock"

// MockKubeCtl is an autogenerated mock type for the KubeCtl type
type MockKubeCtl struct {
	mock.Mock
}

// Exec provides a mock function with given fields: cmdArgs
func (_m *MockKubeCtl) Exec(cmdArgs []string) (string, error) {
	ret := _m.Called(cmdArgs)

	var r0 string
	if rf, ok := ret.Get(0).(func([]string) string); ok {
		r0 = rf(cmdArgs)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(cmdArgs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecStdin provides a mock function with given fields: cmdArgs, stdin
func (_m *MockKubeCtl) ExecStdin(cmdArgs []string, stdin *[]byte) (string, error) {
	ret := _m.Called(cmdArgs, stdin)

	var r0 string
	if rf, ok := ret.Get(0).(func([]string, *[]byte) string); ok {
		r0 = rf(cmdArgs, stdin)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, *[]byte) error); ok {
		r1 = rf(cmdArgs, stdin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
