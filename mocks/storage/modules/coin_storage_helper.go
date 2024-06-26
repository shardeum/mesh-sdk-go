// Code generated by mockery v2.13.1. DO NOT EDIT.

package modules

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	database "github.com/TheArcadiaGroup/rosetta-sdk-go/storage/database"
	types "github.com/coinbase/rosetta-sdk-go/types"
)

// CoinStorageHelper is an autogenerated mock type for the CoinStorageHelper type
type CoinStorageHelper struct {
	mock.Mock
}

// CurrentBlockIdentifier provides a mock function with given fields: _a0, _a1
func (_m *CoinStorageHelper) CurrentBlockIdentifier(_a0 context.Context, _a1 database.Transaction) (*types.BlockIdentifier, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.BlockIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction) *types.BlockIdentifier); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCoinStorageHelper interface {
	mock.TestingT
	Cleanup(func())
}

// NewCoinStorageHelper creates a new instance of CoinStorageHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCoinStorageHelper(t mockConstructorTestingTNewCoinStorageHelper) *CoinStorageHelper {
	mock := &CoinStorageHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
