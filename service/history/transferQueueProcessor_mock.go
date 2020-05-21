// The MIT License (MIT)
//
// Copyright (c) 2017-2020 Uber Technologies Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: transferQueueProcessor.go

// Package history is a generated GoMock package.
package history

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	persistence "github.com/uber/cadence/common/persistence"
)

// MocktransferQueueProcessor is a mock of transferQueueProcessor interface
type MocktransferQueueProcessor struct {
	ctrl     *gomock.Controller
	recorder *MocktransferQueueProcessorMockRecorder
}

// MocktransferQueueProcessorMockRecorder is the mock recorder for MocktransferQueueProcessor
type MocktransferQueueProcessorMockRecorder struct {
	mock *MocktransferQueueProcessor
}

// NewMocktransferQueueProcessor creates a new mock instance
func NewMocktransferQueueProcessor(ctrl *gomock.Controller) *MocktransferQueueProcessor {
	mock := &MocktransferQueueProcessor{ctrl: ctrl}
	mock.recorder = &MocktransferQueueProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocktransferQueueProcessor) EXPECT() *MocktransferQueueProcessorMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MocktransferQueueProcessor) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MocktransferQueueProcessorMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MocktransferQueueProcessor)(nil).Start))
}

// Stop mocks base method
func (m *MocktransferQueueProcessor) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MocktransferQueueProcessorMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MocktransferQueueProcessor)(nil).Stop))
}

// FailoverDomain mocks base method
func (m *MocktransferQueueProcessor) FailoverDomain(domainIDs map[string]struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FailoverDomain", domainIDs)
}

// FailoverDomain indicates an expected call of FailoverDomain
func (mr *MocktransferQueueProcessorMockRecorder) FailoverDomain(domainIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailoverDomain", reflect.TypeOf((*MocktransferQueueProcessor)(nil).FailoverDomain), domainIDs)
}

// NotifyNewTask mocks base method
func (m *MocktransferQueueProcessor) NotifyNewTask(clusterName string, transferTasks []persistence.Task) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyNewTask", clusterName, transferTasks)
}

// NotifyNewTask indicates an expected call of NotifyNewTask
func (mr *MocktransferQueueProcessorMockRecorder) NotifyNewTask(clusterName, transferTasks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewTask", reflect.TypeOf((*MocktransferQueueProcessor)(nil).NotifyNewTask), clusterName, transferTasks)
}

// LockTaskProcessing mocks base method
func (m *MocktransferQueueProcessor) LockTaskProcessing() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LockTaskProcessing")
}

// LockTaskProcessing indicates an expected call of LockTaskProcessing
func (mr *MocktransferQueueProcessorMockRecorder) LockTaskProcessing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockTaskProcessing", reflect.TypeOf((*MocktransferQueueProcessor)(nil).LockTaskProcessing))
}

// UnlockTaskProcessing mocks base method
func (m *MocktransferQueueProcessor) UnlockTaskProcessing() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnlockTaskProcessing")
}

// UnlockTaskProcessing indicates an expected call of UnlockTaskProcessing
func (mr *MocktransferQueueProcessorMockRecorder) UnlockTaskProcessing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockTaskProcessing", reflect.TypeOf((*MocktransferQueueProcessor)(nil).UnlockTaskProcessing))
}
