// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudtrust/keycloak-bridge/pkg/health (interfaces: HealthChecker)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	json "encoding/json"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// HealthChecker is a mock of HealthChecker interface
type HealthChecker struct {
	ctrl     *gomock.Controller
	recorder *HealthCheckerMockRecorder
}

// HealthCheckerMockRecorder is the mock recorder for HealthChecker
type HealthCheckerMockRecorder struct {
	mock *HealthChecker
}

// NewHealthChecker creates a new mock instance
func NewHealthChecker(ctrl *gomock.Controller) *HealthChecker {
	mock := &HealthChecker{ctrl: ctrl}
	mock.recorder = &HealthCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *HealthChecker) EXPECT() *HealthCheckerMockRecorder {
	return m.recorder
}

// AllHealthChecks mocks base method
func (m *HealthChecker) AllHealthChecks(arg0 context.Context) json.RawMessage {
	ret := m.ctrl.Call(m, "AllHealthChecks", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// AllHealthChecks indicates an expected call of AllHealthChecks
func (mr *HealthCheckerMockRecorder) AllHealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllHealthChecks", reflect.TypeOf((*HealthChecker)(nil).AllHealthChecks), arg0)
}

// ExecInfluxHealthChecks mocks base method
func (m *HealthChecker) ExecInfluxHealthChecks(arg0 context.Context) json.RawMessage {
	ret := m.ctrl.Call(m, "ExecInfluxHealthChecks", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ExecInfluxHealthChecks indicates an expected call of ExecInfluxHealthChecks
func (mr *HealthCheckerMockRecorder) ExecInfluxHealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecInfluxHealthChecks", reflect.TypeOf((*HealthChecker)(nil).ExecInfluxHealthChecks), arg0)
}

// ExecJaegerHealthChecks mocks base method
func (m *HealthChecker) ExecJaegerHealthChecks(arg0 context.Context) json.RawMessage {
	ret := m.ctrl.Call(m, "ExecJaegerHealthChecks", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ExecJaegerHealthChecks indicates an expected call of ExecJaegerHealthChecks
func (mr *HealthCheckerMockRecorder) ExecJaegerHealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecJaegerHealthChecks", reflect.TypeOf((*HealthChecker)(nil).ExecJaegerHealthChecks), arg0)
}

// ExecKeycloakHealthChecks mocks base method
func (m *HealthChecker) ExecKeycloakHealthChecks(arg0 context.Context) json.RawMessage {
	ret := m.ctrl.Call(m, "ExecKeycloakHealthChecks", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ExecKeycloakHealthChecks indicates an expected call of ExecKeycloakHealthChecks
func (mr *HealthCheckerMockRecorder) ExecKeycloakHealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecKeycloakHealthChecks", reflect.TypeOf((*HealthChecker)(nil).ExecKeycloakHealthChecks), arg0)
}

// ExecRedisHealthChecks mocks base method
func (m *HealthChecker) ExecRedisHealthChecks(arg0 context.Context) json.RawMessage {
	ret := m.ctrl.Call(m, "ExecRedisHealthChecks", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ExecRedisHealthChecks indicates an expected call of ExecRedisHealthChecks
func (mr *HealthCheckerMockRecorder) ExecRedisHealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecRedisHealthChecks", reflect.TypeOf((*HealthChecker)(nil).ExecRedisHealthChecks), arg0)
}

// ExecSentryHealthChecks mocks base method
func (m *HealthChecker) ExecSentryHealthChecks(arg0 context.Context) json.RawMessage {
	ret := m.ctrl.Call(m, "ExecSentryHealthChecks", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ExecSentryHealthChecks indicates an expected call of ExecSentryHealthChecks
func (mr *HealthCheckerMockRecorder) ExecSentryHealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecSentryHealthChecks", reflect.TypeOf((*HealthChecker)(nil).ExecSentryHealthChecks), arg0)
}

// ReadInfluxHealthChecks mocks base method
func (m *HealthChecker) ReadInfluxHealthChecks(arg0 context.Context) json.RawMessage {
	ret := m.ctrl.Call(m, "ReadInfluxHealthChecks", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ReadInfluxHealthChecks indicates an expected call of ReadInfluxHealthChecks
func (mr *HealthCheckerMockRecorder) ReadInfluxHealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadInfluxHealthChecks", reflect.TypeOf((*HealthChecker)(nil).ReadInfluxHealthChecks), arg0)
}

// ReadJaegerHealthChecks mocks base method
func (m *HealthChecker) ReadJaegerHealthChecks(arg0 context.Context) json.RawMessage {
	ret := m.ctrl.Call(m, "ReadJaegerHealthChecks", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ReadJaegerHealthChecks indicates an expected call of ReadJaegerHealthChecks
func (mr *HealthCheckerMockRecorder) ReadJaegerHealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadJaegerHealthChecks", reflect.TypeOf((*HealthChecker)(nil).ReadJaegerHealthChecks), arg0)
}

// ReadKeycloakHealthChecks mocks base method
func (m *HealthChecker) ReadKeycloakHealthChecks(arg0 context.Context) json.RawMessage {
	ret := m.ctrl.Call(m, "ReadKeycloakHealthChecks", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ReadKeycloakHealthChecks indicates an expected call of ReadKeycloakHealthChecks
func (mr *HealthCheckerMockRecorder) ReadKeycloakHealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadKeycloakHealthChecks", reflect.TypeOf((*HealthChecker)(nil).ReadKeycloakHealthChecks), arg0)
}

// ReadRedisHealthChecks mocks base method
func (m *HealthChecker) ReadRedisHealthChecks(arg0 context.Context) json.RawMessage {
	ret := m.ctrl.Call(m, "ReadRedisHealthChecks", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ReadRedisHealthChecks indicates an expected call of ReadRedisHealthChecks
func (mr *HealthCheckerMockRecorder) ReadRedisHealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadRedisHealthChecks", reflect.TypeOf((*HealthChecker)(nil).ReadRedisHealthChecks), arg0)
}

// ReadSentryHealthChecks mocks base method
func (m *HealthChecker) ReadSentryHealthChecks(arg0 context.Context) json.RawMessage {
	ret := m.ctrl.Call(m, "ReadSentryHealthChecks", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// ReadSentryHealthChecks indicates an expected call of ReadSentryHealthChecks
func (mr *HealthCheckerMockRecorder) ReadSentryHealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSentryHealthChecks", reflect.TypeOf((*HealthChecker)(nil).ReadSentryHealthChecks), arg0)
}
