// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Adapter is an autogenerated mock type for the Adapter type
type Adapter struct {
	mock.Mock
}

// CheckPassangerInformations provides a mock function with given fields: data
func (_m *Adapter) CheckPassangerInformations(data interface{}) {
	_m.Called(data)
}

// RequestPayment provides a mock function with given fields: data
func (_m *Adapter) RequestPayment(data interface{}) {
	_m.Called(data)
}

// SendNotification provides a mock function with given fields: data
func (_m *Adapter) SendNotification(data interface{}) {
	_m.Called(data)
}

// NewAdapter creates a new instance of Adapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdapter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Adapter {
	mock := &Adapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
