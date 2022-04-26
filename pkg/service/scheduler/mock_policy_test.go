// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scylladb/scylla-manager/v3/pkg/service/scheduler (interfaces: Policy)

// Package scheduler is a generated GoMock package.
package scheduler

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/scylladb/scylla-manager/v3/pkg/util/uuid"
)

// mockPolicy is a mock of Policy interface.
type mockPolicy struct {
	ctrl     *gomock.Controller
	recorder *mockPolicyMockRecorder
}

// mockPolicyMockRecorder is the mock recorder for mockPolicy.
type mockPolicyMockRecorder struct {
	mock *mockPolicy
}

// NewmockPolicy creates a new mock instance.
func NewmockPolicy(ctrl *gomock.Controller) *mockPolicy {
	mock := &mockPolicy{ctrl: ctrl}
	mock.recorder = &mockPolicyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *mockPolicy) EXPECT() *mockPolicyMockRecorder {
	return m.recorder
}

// PostRun mocks base method.
func (m *mockPolicy) PostRun(arg0, arg1, arg2 uuid.UUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PostRun", arg0, arg1, arg2)
}

// PostRun indicates an expected call of PostRun.
func (mr *mockPolicyMockRecorder) PostRun(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRun", reflect.TypeOf((*mockPolicy)(nil).PostRun), arg0, arg1, arg2)
}

// PreRun mocks base method.
func (m *mockPolicy) PreRun(arg0, arg1, arg2 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreRun", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PreRun indicates an expected call of PreRun.
func (mr *mockPolicyMockRecorder) PreRun(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreRun", reflect.TypeOf((*mockPolicy)(nil).PreRun), arg0, arg1, arg2)
}
