// Code generated by mockery v2.43.2. DO NOT EDIT.

package pathwrap

import mock "github.com/stretchr/testify/mock"

// MockPath is an autogenerated mock type for the Path type
type MockPath struct {
	mock.Mock
}

type MockPath_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPath) EXPECT() *MockPath_Expecter {
	return &MockPath_Expecter{mock: &_m.Mock}
}

// Abs provides a mock function with given fields: path
func (_m *MockPath) Abs(path string) (string, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Abs")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPath_Abs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Abs'
type MockPath_Abs_Call struct {
	*mock.Call
}

// Abs is a helper method to define mock.On call
//   - path string
func (_e *MockPath_Expecter) Abs(path interface{}) *MockPath_Abs_Call {
	return &MockPath_Abs_Call{Call: _e.mock.On("Abs", path)}
}

func (_c *MockPath_Abs_Call) Run(run func(path string)) *MockPath_Abs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPath_Abs_Call) Return(_a0 string, _a1 error) *MockPath_Abs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPath_Abs_Call) RunAndReturn(run func(string) (string, error)) *MockPath_Abs_Call {
	_c.Call.Return(run)
	return _c
}

// Join provides a mock function with given fields: elem
func (_m *MockPath) Join(elem ...string) string {
	_va := make([]interface{}, len(elem))
	for _i := range elem {
		_va[_i] = elem[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Join")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(...string) string); ok {
		r0 = rf(elem...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPath_Join_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Join'
type MockPath_Join_Call struct {
	*mock.Call
}

// Join is a helper method to define mock.On call
//   - elem ...string
func (_e *MockPath_Expecter) Join(elem ...interface{}) *MockPath_Join_Call {
	return &MockPath_Join_Call{Call: _e.mock.On("Join",
		append([]interface{}{}, elem...)...)}
}

func (_c *MockPath_Join_Call) Run(run func(elem ...string)) *MockPath_Join_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockPath_Join_Call) Return(_a0 string) *MockPath_Join_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPath_Join_Call) RunAndReturn(run func(...string) string) *MockPath_Join_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPath creates a new instance of MockPath. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPath(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPath {
	mock := &MockPath{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
