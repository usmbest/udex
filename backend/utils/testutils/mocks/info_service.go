// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	types "github.com/citypayorg/udex/udex-backend/types"
	mock "github.com/stretchr/testify/mock"
)

// InfoService is an autogenerated mock type for the InfoService type
type InfoService struct {
	mock.Mock
}

// GetExchangeData provides a mock function with given fields:
func (_m *InfoService) GetExchangeData() (*types.ExchangeData, error) {
	ret := _m.Called()

	var r0 *types.ExchangeData
	if rf, ok := ret.Get(0).(func() *types.ExchangeData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ExchangeData)
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

// GetExchangeStats provides a mock function with given fields:
func (_m *InfoService) GetExchangeStats() (*types.ExchangeStats, error) {
	ret := _m.Called()

	var r0 *types.ExchangeStats
	if rf, ok := ret.Get(0).(func() *types.ExchangeStats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ExchangeStats)
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

// GetPairStats provides a mock function with given fields:
func (_m *InfoService) GetPairStats() (*types.PairStats, error) {
	ret := _m.Called()

	var r0 *types.PairStats
	if rf, ok := ret.Get(0).(func() *types.PairStats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.PairStats)
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