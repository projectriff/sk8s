/*
 * Copyright 2018 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by mockery v1.0.0. DO NOT EDIT.

package vendor_mocks

import mock "github.com/stretchr/testify/mock"
import rest "k8s.io/client-go/rest"
import v1 "k8s.io/client-go/kubernetes/typed/core/v1"

// CoreV1Interface is an autogenerated mock type for the CoreV1Interface type
type CoreV1Interface struct {
	mock.Mock
}

// ComponentStatuses provides a mock function with given fields:
func (_m *CoreV1Interface) ComponentStatuses() v1.ComponentStatusInterface {
	ret := _m.Called()

	var r0 v1.ComponentStatusInterface
	if rf, ok := ret.Get(0).(func() v1.ComponentStatusInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.ComponentStatusInterface)
		}
	}

	return r0
}

// ConfigMaps provides a mock function with given fields: namespace
func (_m *CoreV1Interface) ConfigMaps(namespace string) v1.ConfigMapInterface {
	ret := _m.Called(namespace)

	var r0 v1.ConfigMapInterface
	if rf, ok := ret.Get(0).(func(string) v1.ConfigMapInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.ConfigMapInterface)
		}
	}

	return r0
}

// Endpoints provides a mock function with given fields: namespace
func (_m *CoreV1Interface) Endpoints(namespace string) v1.EndpointsInterface {
	ret := _m.Called(namespace)

	var r0 v1.EndpointsInterface
	if rf, ok := ret.Get(0).(func(string) v1.EndpointsInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.EndpointsInterface)
		}
	}

	return r0
}

// Events provides a mock function with given fields: namespace
func (_m *CoreV1Interface) Events(namespace string) v1.EventInterface {
	ret := _m.Called(namespace)

	var r0 v1.EventInterface
	if rf, ok := ret.Get(0).(func(string) v1.EventInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.EventInterface)
		}
	}

	return r0
}

// LimitRanges provides a mock function with given fields: namespace
func (_m *CoreV1Interface) LimitRanges(namespace string) v1.LimitRangeInterface {
	ret := _m.Called(namespace)

	var r0 v1.LimitRangeInterface
	if rf, ok := ret.Get(0).(func(string) v1.LimitRangeInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.LimitRangeInterface)
		}
	}

	return r0
}

// Namespaces provides a mock function with given fields:
func (_m *CoreV1Interface) Namespaces() v1.NamespaceInterface {
	ret := _m.Called()

	var r0 v1.NamespaceInterface
	if rf, ok := ret.Get(0).(func() v1.NamespaceInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.NamespaceInterface)
		}
	}

	return r0
}

// Nodes provides a mock function with given fields:
func (_m *CoreV1Interface) Nodes() v1.NodeInterface {
	ret := _m.Called()

	var r0 v1.NodeInterface
	if rf, ok := ret.Get(0).(func() v1.NodeInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.NodeInterface)
		}
	}

	return r0
}

// PersistentVolumeClaims provides a mock function with given fields: namespace
func (_m *CoreV1Interface) PersistentVolumeClaims(namespace string) v1.PersistentVolumeClaimInterface {
	ret := _m.Called(namespace)

	var r0 v1.PersistentVolumeClaimInterface
	if rf, ok := ret.Get(0).(func(string) v1.PersistentVolumeClaimInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.PersistentVolumeClaimInterface)
		}
	}

	return r0
}

// PersistentVolumes provides a mock function with given fields:
func (_m *CoreV1Interface) PersistentVolumes() v1.PersistentVolumeInterface {
	ret := _m.Called()

	var r0 v1.PersistentVolumeInterface
	if rf, ok := ret.Get(0).(func() v1.PersistentVolumeInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.PersistentVolumeInterface)
		}
	}

	return r0
}

// PodTemplates provides a mock function with given fields: namespace
func (_m *CoreV1Interface) PodTemplates(namespace string) v1.PodTemplateInterface {
	ret := _m.Called(namespace)

	var r0 v1.PodTemplateInterface
	if rf, ok := ret.Get(0).(func(string) v1.PodTemplateInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.PodTemplateInterface)
		}
	}

	return r0
}

// Pods provides a mock function with given fields: namespace
func (_m *CoreV1Interface) Pods(namespace string) v1.PodInterface {
	ret := _m.Called(namespace)

	var r0 v1.PodInterface
	if rf, ok := ret.Get(0).(func(string) v1.PodInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.PodInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *CoreV1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// ReplicationControllers provides a mock function with given fields: namespace
func (_m *CoreV1Interface) ReplicationControllers(namespace string) v1.ReplicationControllerInterface {
	ret := _m.Called(namespace)

	var r0 v1.ReplicationControllerInterface
	if rf, ok := ret.Get(0).(func(string) v1.ReplicationControllerInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.ReplicationControllerInterface)
		}
	}

	return r0
}

// ResourceQuotas provides a mock function with given fields: namespace
func (_m *CoreV1Interface) ResourceQuotas(namespace string) v1.ResourceQuotaInterface {
	ret := _m.Called(namespace)

	var r0 v1.ResourceQuotaInterface
	if rf, ok := ret.Get(0).(func(string) v1.ResourceQuotaInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.ResourceQuotaInterface)
		}
	}

	return r0
}

// Secrets provides a mock function with given fields: namespace
func (_m *CoreV1Interface) Secrets(namespace string) v1.SecretInterface {
	ret := _m.Called(namespace)

	var r0 v1.SecretInterface
	if rf, ok := ret.Get(0).(func(string) v1.SecretInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.SecretInterface)
		}
	}

	return r0
}

// ServiceAccounts provides a mock function with given fields: namespace
func (_m *CoreV1Interface) ServiceAccounts(namespace string) v1.ServiceAccountInterface {
	ret := _m.Called(namespace)

	var r0 v1.ServiceAccountInterface
	if rf, ok := ret.Get(0).(func(string) v1.ServiceAccountInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.ServiceAccountInterface)
		}
	}

	return r0
}

// Services provides a mock function with given fields: namespace
func (_m *CoreV1Interface) Services(namespace string) v1.ServiceInterface {
	ret := _m.Called(namespace)

	var r0 v1.ServiceInterface
	if rf, ok := ret.Get(0).(func(string) v1.ServiceInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.ServiceInterface)
		}
	}

	return r0
}
