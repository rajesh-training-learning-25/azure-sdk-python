// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ../amqpLinks.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	log "github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	internal "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	exported "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/exported"
	gomock "github.com/golang/mock/gomock"
)

// MockAMQPLinks is a mock of AMQPLinks interface.
type MockAMQPLinks struct {
	ctrl     *gomock.Controller
	recorder *MockAMQPLinksMockRecorder
}

// MockAMQPLinksMockRecorder is the mock recorder for MockAMQPLinks.
type MockAMQPLinksMockRecorder struct {
	mock *MockAMQPLinks
}

// NewMockAMQPLinks creates a new mock instance.
func NewMockAMQPLinks(ctrl *gomock.Controller) *MockAMQPLinks {
	mock := &MockAMQPLinks{ctrl: ctrl}
	mock.recorder = &MockAMQPLinksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQPLinks) EXPECT() *MockAMQPLinksMockRecorder {
	return m.recorder
}

// Audience mocks base method.
func (m *MockAMQPLinks) Audience() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Audience")
	ret0, _ := ret[0].(string)
	return ret0
}

// Audience indicates an expected call of Audience.
func (mr *MockAMQPLinksMockRecorder) Audience() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Audience", reflect.TypeOf((*MockAMQPLinks)(nil).Audience))
}

// Close mocks base method.
func (m *MockAMQPLinks) Close(ctx context.Context, permanent bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx, permanent)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAMQPLinksMockRecorder) Close(ctx, permanent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAMQPLinks)(nil).Close), ctx, permanent)
}

// CloseIfNeeded mocks base method.
func (m *MockAMQPLinks) CloseIfNeeded(ctx context.Context, err error) internal.RecoveryKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseIfNeeded", ctx, err)
	ret0, _ := ret[0].(internal.RecoveryKind)
	return ret0
}

// CloseIfNeeded indicates an expected call of CloseIfNeeded.
func (mr *MockAMQPLinksMockRecorder) CloseIfNeeded(ctx, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseIfNeeded", reflect.TypeOf((*MockAMQPLinks)(nil).CloseIfNeeded), ctx, err)
}

// ClosedPermanently mocks base method.
func (m *MockAMQPLinks) ClosedPermanently() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClosedPermanently")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ClosedPermanently indicates an expected call of ClosedPermanently.
func (mr *MockAMQPLinksMockRecorder) ClosedPermanently() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClosedPermanently", reflect.TypeOf((*MockAMQPLinks)(nil).ClosedPermanently))
}

// EntityPath mocks base method.
func (m *MockAMQPLinks) EntityPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EntityPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// EntityPath indicates an expected call of EntityPath.
func (mr *MockAMQPLinksMockRecorder) EntityPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntityPath", reflect.TypeOf((*MockAMQPLinks)(nil).EntityPath))
}

// Get mocks base method.
func (m *MockAMQPLinks) Get(ctx context.Context) (*internal.LinksWithID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(*internal.LinksWithID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAMQPLinksMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAMQPLinks)(nil).Get), ctx)
}

// ManagementPath mocks base method.
func (m *MockAMQPLinks) ManagementPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManagementPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// ManagementPath indicates an expected call of ManagementPath.
func (mr *MockAMQPLinksMockRecorder) ManagementPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManagementPath", reflect.TypeOf((*MockAMQPLinks)(nil).ManagementPath))
}

// RecoverIfNeeded mocks base method.
func (m *MockAMQPLinks) RecoverIfNeeded(ctx context.Context, linkID internal.LinkID, err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecoverIfNeeded", ctx, linkID, err)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecoverIfNeeded indicates an expected call of RecoverIfNeeded.
func (mr *MockAMQPLinksMockRecorder) RecoverIfNeeded(ctx, linkID, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecoverIfNeeded", reflect.TypeOf((*MockAMQPLinks)(nil).RecoverIfNeeded), ctx, linkID, err)
}

// Retry mocks base method.
func (m *MockAMQPLinks) Retry(ctx context.Context, name log.Event, operation string, fn internal.RetryWithLinksFn, o exported.RetryOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retry", ctx, name, operation, fn, o)
	ret0, _ := ret[0].(error)
	return ret0
}

// Retry indicates an expected call of Retry.
func (mr *MockAMQPLinksMockRecorder) Retry(ctx, name, operation, fn, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retry", reflect.TypeOf((*MockAMQPLinks)(nil).Retry), ctx, name, operation, fn, o)
}
