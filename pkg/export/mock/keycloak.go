// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudtrust/keycloak-bridge/pkg/export (interfaces: KeycloakClient)

// Package mock is a generated GoMock package.
package mock

import (
	keycloak_client "github.com/cloudtrust/keycloak-client"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// KeycloakClient is a mock of KeycloakClient interface
type KeycloakClient struct {
	ctrl     *gomock.Controller
	recorder *KeycloakClientMockRecorder
}

// KeycloakClientMockRecorder is the mock recorder for KeycloakClient
type KeycloakClientMockRecorder struct {
	mock *KeycloakClient
}

// NewKeycloakClient creates a new mock instance
func NewKeycloakClient(ctrl *gomock.Controller) *KeycloakClient {
	mock := &KeycloakClient{ctrl: ctrl}
	mock.recorder = &KeycloakClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *KeycloakClient) EXPECT() *KeycloakClientMockRecorder {
	return m.recorder
}

// ExportRealm mocks base method
func (m *KeycloakClient) ExportRealm(arg0 string) (keycloak_client.RealmRepresentation, error) {
	ret := m.ctrl.Call(m, "ExportRealm", arg0)
	ret0, _ := ret[0].(keycloak_client.RealmRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportRealm indicates an expected call of ExportRealm
func (mr *KeycloakClientMockRecorder) ExportRealm(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportRealm", reflect.TypeOf((*KeycloakClient)(nil).ExportRealm), arg0)
}

// GetRealms mocks base method
func (m *KeycloakClient) GetRealms() ([]keycloak_client.RealmRepresentation, error) {
	ret := m.ctrl.Call(m, "GetRealms")
	ret0, _ := ret[0].([]keycloak_client.RealmRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRealms indicates an expected call of GetRealms
func (mr *KeycloakClientMockRecorder) GetRealms() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRealms", reflect.TypeOf((*KeycloakClient)(nil).GetRealms))
}
