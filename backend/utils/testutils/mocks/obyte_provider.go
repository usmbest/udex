// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	types "github.com/citypayorg/udex/backend/types"
	mock "github.com/stretchr/testify/mock"
)

// ObyteProvider is an autogenerated mock type for the ObyteProvider type
type ObyteProvider struct {
	mock.Mock
}

// AddOrder provides a mock function with given fields: signedOrder
func (_m *ObyteProvider) AddOrder(signedOrder *interface{}) (string, error) {
	ret := _m.Called(signedOrder)

	var r0 string
	if rf, ok := ret.Get(0).(func(*interface{}) string); ok {
		r0 = rf(signedOrder)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*interface{}) error); ok {
		r1 = rf(signedOrder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Asset provides a mock function with given fields: symbol
func (_m *ObyteProvider) Asset(symbol string) (string, error) {
	ret := _m.Called(symbol)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(symbol)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(symbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BalanceOf provides a mock function with given fields: owner, token
func (_m *ObyteProvider) BalanceOf(owner string, token string) (int64, error) {
	ret := _m.Called(owner, token)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(owner, token)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelOrder provides a mock function with given fields: signedCancel
func (_m *ObyteProvider) CancelOrder(signedCancel *interface{}) error {
	ret := _m.Called(signedCancel)

	var r0 error
	if rf, ok := ret.Get(0).(func(*interface{}) error); ok {
		r0 = rf(signedCancel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Decimals provides a mock function with given fields: token
func (_m *ObyteProvider) Decimals(token string) (uint8, error) {
	ret := _m.Called(token)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(string) uint8); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteTrade provides a mock function with given fields: m
func (_m *ObyteProvider) ExecuteTrade(m *types.Matches) ([]string, error) {
	ret := _m.Called(m)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*types.Matches) []string); ok {
		r0 = rf(m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Matches) error); ok {
		r1 = rf(m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthorizedAddresses provides a mock function with given fields: address
func (_m *ObyteProvider) GetAuthorizedAddresses(address string) ([]string, error) {
	ret := _m.Called(address)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalances provides a mock function with given fields: owner
func (_m *ObyteProvider) GetBalances(owner string) map[string]int64 {
	ret := _m.Called(owner)

	var r0 map[string]int64
	if rf, ok := ret.Get(0).(func(string) map[string]int64); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]int64)
		}
	}

	return r0
}

// GetFees provides a mock function with given fields:
func (_m *ObyteProvider) GetFees() (float64, float64) {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func() float64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(float64)
	}

	return r0, r1
}

// GetOperatorAddress provides a mock function with given fields:
func (_m *ObyteProvider) GetOperatorAddress() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ListenToEvents provides a mock function with given fields:
func (_m *ObyteProvider) ListenToEvents() (chan map[string]interface{}, error) {
	ret := _m.Called()

	var r0 chan map[string]interface{}
	if rf, ok := ret.Get(0).(func() chan map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Symbol provides a mock function with given fields: token
func (_m *ObyteProvider) Symbol(token string) (string, error) {
	ret := _m.Called(token)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
