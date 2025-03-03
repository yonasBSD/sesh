// Code generated by mockery v2.52.3. DO NOT EDIT.

package connector

import (
	model "github.com/joshmedeski/sesh/v2/model"
	mock "github.com/stretchr/testify/mock"
)

// MockConnector is an autogenerated mock type for the Connector type
type MockConnector struct {
	mock.Mock
}

type MockConnector_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConnector) EXPECT() *MockConnector_Expecter {
	return &MockConnector_Expecter{mock: &_m.Mock}
}

// Connect provides a mock function with given fields: name, opts
func (_m *MockConnector) Connect(name string, opts model.ConnectOpts) (string, error) {
	ret := _m.Called(name, opts)

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, model.ConnectOpts) (string, error)); ok {
		return rf(name, opts)
	}
	if rf, ok := ret.Get(0).(func(string, model.ConnectOpts) string); ok {
		r0 = rf(name, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, model.ConnectOpts) error); ok {
		r1 = rf(name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnector_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockConnector_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
//   - name string
//   - opts model.ConnectOpts
func (_e *MockConnector_Expecter) Connect(name interface{}, opts interface{}) *MockConnector_Connect_Call {
	return &MockConnector_Connect_Call{Call: _e.mock.On("Connect", name, opts)}
}

func (_c *MockConnector_Connect_Call) Run(run func(name string, opts model.ConnectOpts)) *MockConnector_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(model.ConnectOpts))
	})
	return _c
}

func (_c *MockConnector_Connect_Call) Return(_a0 string, _a1 error) *MockConnector_Connect_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnector_Connect_Call) RunAndReturn(run func(string, model.ConnectOpts) (string, error)) *MockConnector_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConnector creates a new instance of MockConnector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConnector(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConnector {
	mock := &MockConnector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
