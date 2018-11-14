// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import core "github.com/projectriff/riff/pkg/core"
import io "io"
import mock "github.com/stretchr/testify/mock"
import servingv1alpha1 "github.com/knative/serving/pkg/apis/serving/v1alpha1"
import v1alpha1 "github.com/knative/eventing/pkg/apis/channels/v1alpha1"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateChannel provides a mock function with given fields: options
func (_m *Client) CreateChannel(options core.CreateChannelOptions) (*v1alpha1.Channel, error) {
	ret := _m.Called(options)

	var r0 *v1alpha1.Channel
	if rf, ok := ret.Get(0).(func(core.CreateChannelOptions) *v1alpha1.Channel); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.CreateChannelOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFunction provides a mock function with given fields: options, log
func (_m *Client) CreateFunction(options core.CreateFunctionOptions, log io.Writer) (*servingv1alpha1.Service, error) {
	ret := _m.Called(options, log)

	var r0 *servingv1alpha1.Service
	if rf, ok := ret.Get(0).(func(core.CreateFunctionOptions, io.Writer) *servingv1alpha1.Service); ok {
		r0 = rf(options, log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servingv1alpha1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.CreateFunctionOptions, io.Writer) error); ok {
		r1 = rf(options, log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateService provides a mock function with given fields: options
func (_m *Client) CreateService(options core.CreateOrUpdateServiceOptions) (*servingv1alpha1.Service, error) {
	ret := _m.Called(options)

	var r0 *servingv1alpha1.Service
	if rf, ok := ret.Get(0).(func(core.CreateOrUpdateServiceOptions) *servingv1alpha1.Service); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servingv1alpha1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.CreateOrUpdateServiceOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSubscription provides a mock function with given fields: options
func (_m *Client) CreateSubscription(options core.CreateSubscriptionOptions) (*v1alpha1.Subscription, error) {
	ret := _m.Called(options)

	var r0 *v1alpha1.Subscription
	if rf, ok := ret.Get(0).(func(core.CreateSubscriptionOptions) *v1alpha1.Subscription); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.CreateSubscriptionOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteChannel provides a mock function with given fields: options
func (_m *Client) DeleteChannel(options core.DeleteChannelOptions) error {
	ret := _m.Called(options)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.DeleteChannelOptions) error); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteService provides a mock function with given fields: options
func (_m *Client) DeleteService(options core.DeleteServiceOptions) error {
	ret := _m.Called(options)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.DeleteServiceOptions) error); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSubscription provides a mock function with given fields: options
func (_m *Client) DeleteSubscription(options core.DeleteSubscriptionOptions) error {
	ret := _m.Called(options)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.DeleteSubscriptionOptions) error); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListChannels provides a mock function with given fields: options
func (_m *Client) ListChannels(options core.ListChannelOptions) (*v1alpha1.ChannelList, error) {
	ret := _m.Called(options)

	var r0 *v1alpha1.ChannelList
	if rf, ok := ret.Get(0).(func(core.ListChannelOptions) *v1alpha1.ChannelList); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ChannelList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.ListChannelOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListServices provides a mock function with given fields: options
func (_m *Client) ListServices(options core.ListServiceOptions) (*servingv1alpha1.ServiceList, error) {
	ret := _m.Called(options)

	var r0 *servingv1alpha1.ServiceList
	if rf, ok := ret.Get(0).(func(core.ListServiceOptions) *servingv1alpha1.ServiceList); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servingv1alpha1.ServiceList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.ListServiceOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSubscriptions provides a mock function with given fields: options
func (_m *Client) ListSubscriptions(options core.ListSubscriptionsOptions) (*v1alpha1.SubscriptionList, error) {
	ret := _m.Called(options)

	var r0 *v1alpha1.SubscriptionList
	if rf, ok := ret.Get(0).(func(core.ListSubscriptionsOptions) *v1alpha1.SubscriptionList); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.SubscriptionList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.ListSubscriptionsOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceCoordinates provides a mock function with given fields: options
func (_m *Client) ServiceCoordinates(options core.ServiceInvokeOptions) (string, string, error) {
	ret := _m.Called(options)

	var r0 string
	if rf, ok := ret.Get(0).(func(core.ServiceInvokeOptions) string); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(core.ServiceInvokeOptions) string); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(core.ServiceInvokeOptions) error); ok {
		r2 = rf(options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ServiceStatus provides a mock function with given fields: options
func (_m *Client) ServiceStatus(options core.ServiceStatusOptions) (*servingv1alpha1.ServiceCondition, error) {
	ret := _m.Called(options)

	var r0 *servingv1alpha1.ServiceCondition
	if rf, ok := ret.Get(0).(func(core.ServiceStatusOptions) *servingv1alpha1.ServiceCondition); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servingv1alpha1.ServiceCondition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.ServiceStatusOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFunction provides a mock function with given fields: options, log
func (_m *Client) UpdateFunction(options core.UpdateFunctionOptions, log io.Writer) error {
	ret := _m.Called(options, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.UpdateFunctionOptions, io.Writer) error); ok {
		r0 = rf(options, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateService provides a mock function with given fields: options
func (_m *Client) UpdateService(options core.CreateOrUpdateServiceOptions) (*servingv1alpha1.Service, error) {
	ret := _m.Called(options)

	var r0 *servingv1alpha1.Service
	if rf, ok := ret.Get(0).(func(core.CreateOrUpdateServiceOptions) *servingv1alpha1.Service); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servingv1alpha1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.CreateOrUpdateServiceOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
