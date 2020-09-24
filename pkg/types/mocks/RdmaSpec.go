// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	v1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

// RdmaSpec is an autogenerated mock type for the RdmaSpec type
type RdmaSpec struct {
	mock.Mock
}

// GetRdmaDeviceSpec provides a mock function with given fields:
func (_m *RdmaSpec) GetRdmaDeviceSpec() []*v1beta1.DeviceSpec {
	ret := _m.Called()

	var r0 []*v1beta1.DeviceSpec
	if rf, ok := ret.Get(0).(func() []*v1beta1.DeviceSpec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.DeviceSpec)
		}
	}

	return r0
}

// IsRdma provides a mock function with given fields:
func (_m *RdmaSpec) IsRdma() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
