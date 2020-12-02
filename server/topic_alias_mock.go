// Code generated by MockGen. DO NOT EDIT.
// Source: server/topic_alias.go

// Package server is a generated GoMock package.
package server

import (
	packets "github.com/DrmagicE/gmqtt/pkg/packets"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTopicAliasManager is a mock of TopicAliasManager interface
type MockTopicAliasManager struct {
	ctrl     *gomock.Controller
	recorder *MockTopicAliasManagerMockRecorder
}

// MockTopicAliasManagerMockRecorder is the mock recorder for MockTopicAliasManager
type MockTopicAliasManagerMockRecorder struct {
	mock *MockTopicAliasManager
}

// NewMockTopicAliasManager creates a new mock instance
func NewMockTopicAliasManager(ctrl *gomock.Controller) *MockTopicAliasManager {
	mock := &MockTopicAliasManager{ctrl: ctrl}
	mock.recorder = &MockTopicAliasManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTopicAliasManager) EXPECT() *MockTopicAliasManagerMockRecorder {
	return m.recorder
}

// Check mocks base method
func (m *MockTopicAliasManager) Check(publish *packets.Publish) (uint16, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", publish)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Check indicates an expected call of Check
func (mr *MockTopicAliasManagerMockRecorder) Check(publish interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockTopicAliasManager)(nil).Check), publish)
}
