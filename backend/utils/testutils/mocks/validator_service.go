// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	types "github.com/citypayorg/udex/backend/types"
	mock "github.com/stretchr/testify/mock"
)

// ValidatorService is an autogenerated mock type for the ValidatorService type
type ValidatorService struct {
	mock.Mock
}

// ValidateAvailableBalance provides a mock function with given fields: o, uncommittedDeltas, balanceLockedInMemoryOrders
func (_m *ValidatorService) ValidateAvailableBalance(o *types.Order, uncommittedDeltas map[string]int64, balanceLockedInMemoryOrders int64) error {
	ret := _m.Called(o, uncommittedDeltas, balanceLockedInMemoryOrders)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Order, map[string]int64, int64) error); ok {
		r0 = rf(o, uncommittedDeltas, balanceLockedInMemoryOrders)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateBalance provides a mock function with given fields: o
func (_m *ValidatorService) ValidateBalance(o *types.Order) error {
	ret := _m.Called(o)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Order) error); ok {
		r0 = rf(o)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateOperatorAddress provides a mock function with given fields: o
func (_m *ValidatorService) ValidateOperatorAddress(o *types.Order) error {
	ret := _m.Called(o)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Order) error); ok {
		r0 = rf(o)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
