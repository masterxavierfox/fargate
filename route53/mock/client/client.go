// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/masterxavierfox/fargate/route53 (interfaces: Client)

// Package client is a generated GoMock package.
package client

import (
	gomock "github.com/golang/mock/gomock"
	route53 "github.com/masterxavierfox/fargate/route53"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateAlias mocks base method
func (m *MockClient) CreateAlias(arg0 route53.CreateAliasInput) (string, error) {
	ret := m.ctrl.Call(m, "CreateAlias", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAlias indicates an expected call of CreateAlias
func (mr *MockClientMockRecorder) CreateAlias(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAlias", reflect.TypeOf((*MockClient)(nil).CreateAlias), arg0)
}

// CreateResourceRecord mocks base method
func (m *MockClient) CreateResourceRecord(arg0 route53.CreateResourceRecordInput) (string, error) {
	ret := m.ctrl.Call(m, "CreateResourceRecord", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateResourceRecord indicates an expected call of CreateResourceRecord
func (mr *MockClientMockRecorder) CreateResourceRecord(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResourceRecord", reflect.TypeOf((*MockClient)(nil).CreateResourceRecord), arg0)
}

// ListHostedZones mocks base method
func (m *MockClient) ListHostedZones() (route53.HostedZones, error) {
	ret := m.ctrl.Call(m, "ListHostedZones")
	ret0, _ := ret[0].(route53.HostedZones)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostedZones indicates an expected call of ListHostedZones
func (mr *MockClientMockRecorder) ListHostedZones() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostedZones", reflect.TypeOf((*MockClient)(nil).ListHostedZones))
}
