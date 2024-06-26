// Code generated by mockery v2.13.1. DO NOT EDIT.

package utils

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	fetcher "github.com/TheArcadiaGroup/rosetta-sdk-go/fetcher"
	types "github.com/coinbase/rosetta-sdk-go/types"
)

// FetcherHelper is an autogenerated mock type for the FetcherHelper type
type FetcherHelper struct {
	mock.Mock
}

// AccountBalanceRetry provides a mock function with given fields: ctx, network, account, block, currencies
func (_m *FetcherHelper) AccountBalanceRetry(ctx context.Context, network *types.NetworkIdentifier, account *types.AccountIdentifier, block *types.PartialBlockIdentifier, currencies []*types.Currency) (*types.BlockIdentifier, []*types.Amount, map[string]interface{}, *fetcher.Error) {
	ret := _m.Called(ctx, network, account, block, currencies)

	var r0 *types.BlockIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, *types.PartialBlockIdentifier, []*types.Currency) *types.BlockIdentifier); ok {
		r0 = rf(ctx, network, account, block, currencies)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockIdentifier)
		}
	}

	var r1 []*types.Amount
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, *types.PartialBlockIdentifier, []*types.Currency) []*types.Amount); ok {
		r1 = rf(ctx, network, account, block, currencies)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*types.Amount)
		}
	}

	var r2 map[string]interface{}
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, *types.PartialBlockIdentifier, []*types.Currency) map[string]interface{}); ok {
		r2 = rf(ctx, network, account, block, currencies)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(map[string]interface{})
		}
	}

	var r3 *fetcher.Error
	if rf, ok := ret.Get(3).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, *types.PartialBlockIdentifier, []*types.Currency) *fetcher.Error); ok {
		r3 = rf(ctx, network, account, block, currencies)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(*fetcher.Error)
		}
	}

	return r0, r1, r2, r3
}

// AccountCoinsRetry provides a mock function with given fields: ctx, network, acct, includeMempool, currencies
func (_m *FetcherHelper) AccountCoinsRetry(ctx context.Context, network *types.NetworkIdentifier, acct *types.AccountIdentifier, includeMempool bool, currencies []*types.Currency) (*types.BlockIdentifier, []*types.Coin, map[string]interface{}, *fetcher.Error) {
	ret := _m.Called(ctx, network, acct, includeMempool, currencies)

	var r0 *types.BlockIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, bool, []*types.Currency) *types.BlockIdentifier); ok {
		r0 = rf(ctx, network, acct, includeMempool, currencies)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockIdentifier)
		}
	}

	var r1 []*types.Coin
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, bool, []*types.Currency) []*types.Coin); ok {
		r1 = rf(ctx, network, acct, includeMempool, currencies)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*types.Coin)
		}
	}

	var r2 map[string]interface{}
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, bool, []*types.Currency) map[string]interface{}); ok {
		r2 = rf(ctx, network, acct, includeMempool, currencies)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(map[string]interface{})
		}
	}

	var r3 *fetcher.Error
	if rf, ok := ret.Get(3).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, bool, []*types.Currency) *fetcher.Error); ok {
		r3 = rf(ctx, network, acct, includeMempool, currencies)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(*fetcher.Error)
		}
	}

	return r0, r1, r2, r3
}

// NetworkList provides a mock function with given fields: ctx, metadata
func (_m *FetcherHelper) NetworkList(ctx context.Context, metadata map[string]interface{}) (*types.NetworkListResponse, *fetcher.Error) {
	ret := _m.Called(ctx, metadata)

	var r0 *types.NetworkListResponse
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) *types.NetworkListResponse); ok {
		r0 = rf(ctx, metadata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NetworkListResponse)
		}
	}

	var r1 *fetcher.Error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) *fetcher.Error); ok {
		r1 = rf(ctx, metadata)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*fetcher.Error)
		}
	}

	return r0, r1
}

// NetworkStatusRetry provides a mock function with given fields: ctx, network, metadata
func (_m *FetcherHelper) NetworkStatusRetry(ctx context.Context, network *types.NetworkIdentifier, metadata map[string]interface{}) (*types.NetworkStatusResponse, *fetcher.Error) {
	ret := _m.Called(ctx, network, metadata)

	var r0 *types.NetworkStatusResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, map[string]interface{}) *types.NetworkStatusResponse); ok {
		r0 = rf(ctx, network, metadata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NetworkStatusResponse)
		}
	}

	var r1 *fetcher.Error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, map[string]interface{}) *fetcher.Error); ok {
		r1 = rf(ctx, network, metadata)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*fetcher.Error)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewFetcherHelper interface {
	mock.TestingT
	Cleanup(func())
}

// NewFetcherHelper creates a new instance of FetcherHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFetcherHelper(t mockConstructorTestingTNewFetcherHelper) *FetcherHelper {
	mock := &FetcherHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
