// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	types "github.com/citypayorg/udex/backend/types"
	mock "github.com/stretchr/testify/mock"

	ws "github.com/citypayorg/udex/backend/ws"
)

// OrderBookService is an autogenerated mock type for the OrderBookService type
type OrderBookService struct {
	mock.Mock
}

// GetOrderBook provides a mock function with given fields: bt, qt
func (_m *OrderBookService) GetOrderBook(bt string, qt string) (map[string]interface{}, error) {
	ret := _m.Called(bt, qt)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string, string) map[string]interface{}); ok {
		r0 = rf(bt, qt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(bt, qt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRawOrderBook provides a mock function with given fields: bt, qt
func (_m *OrderBookService) GetRawOrderBook(bt string, qt string) (*types.RawOrderBook, error) {
	ret := _m.Called(bt, qt)

	var r0 *types.RawOrderBook
	if rf, ok := ret.Get(0).(func(string, string) *types.RawOrderBook); ok {
		r0 = rf(bt, qt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.RawOrderBook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(bt, qt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeOrderBook provides a mock function with given fields: c, bt, qt
func (_m *OrderBookService) SubscribeOrderBook(c *ws.Client, bt string, qt string) {
	_m.Called(c, bt, qt)
}

// SubscribeRawOrderBook provides a mock function with given fields: c, bt, qt
func (_m *OrderBookService) SubscribeRawOrderBook(c *ws.Client, bt string, qt string) {
	_m.Called(c, bt, qt)
}

// UnsubscribeOrderBook provides a mock function with given fields: c
func (_m *OrderBookService) UnsubscribeOrderBook(c *ws.Client) {
	_m.Called(c)
}

// UnsubscribeOrderBookChannel provides a mock function with given fields: c, bt, qt
func (_m *OrderBookService) UnsubscribeOrderBookChannel(c *ws.Client, bt string, qt string) {
	_m.Called(c, bt, qt)
}

// UnsubscribeRawOrderBook provides a mock function with given fields: c
func (_m *OrderBookService) UnsubscribeRawOrderBook(c *ws.Client) {
	_m.Called(c)
}

// UnsubscribeRawOrderBookChannel provides a mock function with given fields: c, bt, qt
func (_m *OrderBookService) UnsubscribeRawOrderBookChannel(c *ws.Client, bt string, qt string) {
	_m.Called(c, bt, qt)
}
