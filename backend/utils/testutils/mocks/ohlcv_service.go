// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	types "github.com/citypayorg/udex/udex-backend/types"
	mock "github.com/stretchr/testify/mock"

	ws "github.com/citypayorg/udex/udex-backend/ws"
)

// OHLCVService is an autogenerated mock type for the OHLCVService type
type OHLCVService struct {
	mock.Mock
}

// GetOHLCV provides a mock function with given fields: p, duration, unit, timeInterval
func (_m *OHLCVService) GetOHLCV(p []types.PairAssets, duration int64, unit string, timeInterval ...int64) ([]*types.Tick, error) {
	_va := make([]interface{}, len(timeInterval))
	for _i := range timeInterval {
		_va[_i] = timeInterval[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, p, duration, unit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*types.Tick
	if rf, ok := ret.Get(0).(func([]types.PairAssets, int64, string, ...int64) []*types.Tick); ok {
		r0 = rf(p, duration, unit, timeInterval...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Tick)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]types.PairAssets, int64, string, ...int64) error); ok {
		r1 = rf(p, duration, unit, timeInterval...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: c, p
func (_m *OHLCVService) Subscribe(c *ws.Client, p *types.SubscriptionPayload) {
	_m.Called(c, p)
}

// Unsubscribe provides a mock function with given fields: c
func (_m *OHLCVService) Unsubscribe(c *ws.Client) {
	_m.Called(c)
}

// UnsubscribeChannel provides a mock function with given fields: c, p
func (_m *OHLCVService) UnsubscribeChannel(c *ws.Client, p *types.SubscriptionPayload) {
	_m.Called(c, p)
}