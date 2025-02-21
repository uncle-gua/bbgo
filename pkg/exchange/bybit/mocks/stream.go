// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/c9s/bbgo/pkg/exchange/bybit (interfaces: MarketInfoProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	bybitapi "github.com/c9s/bbgo/pkg/exchange/bybit/bybitapi"
	types "github.com/c9s/bbgo/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockMarketInfoProvider is a mock of MarketInfoProvider interface.
type MockMarketInfoProvider struct {
	ctrl     *gomock.Controller
	recorder *MockMarketInfoProviderMockRecorder
}

// MockMarketInfoProviderMockRecorder is the mock recorder for MockMarketInfoProvider.
type MockMarketInfoProviderMockRecorder struct {
	mock *MockMarketInfoProvider
}

// NewMockMarketInfoProvider creates a new mock instance.
func NewMockMarketInfoProvider(ctrl *gomock.Controller) *MockMarketInfoProvider {
	mock := &MockMarketInfoProvider{ctrl: ctrl}
	mock.recorder = &MockMarketInfoProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarketInfoProvider) EXPECT() *MockMarketInfoProviderMockRecorder {
	return m.recorder
}

// GetAllFeeRates mocks base method.
func (m *MockMarketInfoProvider) GetAllFeeRates(arg0 context.Context) (bybitapi.FeeRates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFeeRates", arg0)
	ret0, _ := ret[0].(bybitapi.FeeRates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFeeRates indicates an expected call of GetAllFeeRates.
func (mr *MockMarketInfoProviderMockRecorder) GetAllFeeRates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFeeRates", reflect.TypeOf((*MockMarketInfoProvider)(nil).GetAllFeeRates), arg0)
}

// QueryMarkets mocks base method.
func (m *MockMarketInfoProvider) QueryMarkets(arg0 context.Context) (types.MarketMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryMarkets", arg0)
	ret0, _ := ret[0].(types.MarketMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryMarkets indicates an expected call of QueryMarkets.
func (mr *MockMarketInfoProviderMockRecorder) QueryMarkets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryMarkets", reflect.TypeOf((*MockMarketInfoProvider)(nil).QueryMarkets), arg0)
}
