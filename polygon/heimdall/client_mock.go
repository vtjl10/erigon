// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon/polygon/heimdall (interfaces: HeimdallClient)

// Package heimdall is a generated GoMock package.
package heimdall

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHeimdallClient is a mock of HeimdallClient interface.
type MockHeimdallClient struct {
	ctrl     *gomock.Controller
	recorder *MockHeimdallClientMockRecorder
}

// MockHeimdallClientMockRecorder is the mock recorder for MockHeimdallClient.
type MockHeimdallClientMockRecorder struct {
	mock *MockHeimdallClient
}

// NewMockHeimdallClient creates a new mock instance.
func NewMockHeimdallClient(ctrl *gomock.Controller) *MockHeimdallClient {
	mock := &MockHeimdallClient{ctrl: ctrl}
	mock.recorder = &MockHeimdallClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHeimdallClient) EXPECT() *MockHeimdallClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockHeimdallClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockHeimdallClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHeimdallClient)(nil).Close))
}

// FetchCheckpoint mocks base method.
func (m *MockHeimdallClient) FetchCheckpoint(arg0 context.Context, arg1 int64) (*Checkpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCheckpoint", arg0, arg1)
	ret0, _ := ret[0].(*Checkpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCheckpoint indicates an expected call of FetchCheckpoint.
func (mr *MockHeimdallClientMockRecorder) FetchCheckpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCheckpoint", reflect.TypeOf((*MockHeimdallClient)(nil).FetchCheckpoint), arg0, arg1)
}

// FetchCheckpointCount mocks base method.
func (m *MockHeimdallClient) FetchCheckpointCount(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCheckpointCount", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCheckpointCount indicates an expected call of FetchCheckpointCount.
func (mr *MockHeimdallClientMockRecorder) FetchCheckpointCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCheckpointCount", reflect.TypeOf((*MockHeimdallClient)(nil).FetchCheckpointCount), arg0)
}

// FetchLastNoAckMilestone mocks base method.
func (m *MockHeimdallClient) FetchLastNoAckMilestone(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchLastNoAckMilestone", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchLastNoAckMilestone indicates an expected call of FetchLastNoAckMilestone.
func (mr *MockHeimdallClientMockRecorder) FetchLastNoAckMilestone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchLastNoAckMilestone", reflect.TypeOf((*MockHeimdallClient)(nil).FetchLastNoAckMilestone), arg0)
}

// FetchMilestone mocks base method.
func (m *MockHeimdallClient) FetchMilestone(arg0 context.Context, arg1 int64) (*Milestone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMilestone", arg0, arg1)
	ret0, _ := ret[0].(*Milestone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMilestone indicates an expected call of FetchMilestone.
func (mr *MockHeimdallClientMockRecorder) FetchMilestone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMilestone", reflect.TypeOf((*MockHeimdallClient)(nil).FetchMilestone), arg0, arg1)
}

// FetchMilestoneCount mocks base method.
func (m *MockHeimdallClient) FetchMilestoneCount(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMilestoneCount", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMilestoneCount indicates an expected call of FetchMilestoneCount.
func (mr *MockHeimdallClientMockRecorder) FetchMilestoneCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMilestoneCount", reflect.TypeOf((*MockHeimdallClient)(nil).FetchMilestoneCount), arg0)
}

// FetchMilestoneID mocks base method.
func (m *MockHeimdallClient) FetchMilestoneID(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMilestoneID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchMilestoneID indicates an expected call of FetchMilestoneID.
func (mr *MockHeimdallClientMockRecorder) FetchMilestoneID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMilestoneID", reflect.TypeOf((*MockHeimdallClient)(nil).FetchMilestoneID), arg0, arg1)
}

// FetchNoAckMilestone mocks base method.
func (m *MockHeimdallClient) FetchNoAckMilestone(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchNoAckMilestone", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchNoAckMilestone indicates an expected call of FetchNoAckMilestone.
func (mr *MockHeimdallClientMockRecorder) FetchNoAckMilestone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchNoAckMilestone", reflect.TypeOf((*MockHeimdallClient)(nil).FetchNoAckMilestone), arg0, arg1)
}

// Span mocks base method.
func (m *MockHeimdallClient) Span(arg0 context.Context, arg1 uint64) (*HeimdallSpan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Span", arg0, arg1)
	ret0, _ := ret[0].(*HeimdallSpan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Span indicates an expected call of Span.
func (mr *MockHeimdallClientMockRecorder) Span(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Span", reflect.TypeOf((*MockHeimdallClient)(nil).Span), arg0, arg1)
}

// StateSyncEvents mocks base method.
func (m *MockHeimdallClient) StateSyncEvents(arg0 context.Context, arg1 uint64, arg2 int64) ([]*EventRecordWithTime, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSyncEvents", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*EventRecordWithTime)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSyncEvents indicates an expected call of StateSyncEvents.
func (mr *MockHeimdallClientMockRecorder) StateSyncEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSyncEvents", reflect.TypeOf((*MockHeimdallClient)(nil).StateSyncEvents), arg0, arg1, arg2)
}