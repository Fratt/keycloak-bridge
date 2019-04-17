// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudtrust/keycloak-bridge/pkg/management (interfaces: ManagementComponent)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	management "github.com/cloudtrust/keycloak-bridge/api/management"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// ManagementComponent is a mock of ManagementComponent interface
type ManagementComponent struct {
	ctrl     *gomock.Controller
	recorder *ManagementComponentMockRecorder
}

// ManagementComponentMockRecorder is the mock recorder for ManagementComponent
type ManagementComponentMockRecorder struct {
	mock *ManagementComponent
}

// NewManagementComponent creates a new mock instance
func NewManagementComponent(ctrl *gomock.Controller) *ManagementComponent {
	mock := &ManagementComponent{ctrl: ctrl}
	mock.recorder = &ManagementComponentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *ManagementComponent) EXPECT() *ManagementComponentMockRecorder {
	return m.recorder
}

// AddClientRolesToUser mocks base method
func (m *ManagementComponent) AddClientRolesToUser(arg0 context.Context, arg1, arg2, arg3 string, arg4 []management.RoleRepresentation) error {
	ret := m.ctrl.Call(m, "AddClientRolesToUser", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddClientRolesToUser indicates an expected call of AddClientRolesToUser
func (mr *ManagementComponentMockRecorder) AddClientRolesToUser(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClientRolesToUser", reflect.TypeOf((*ManagementComponent)(nil).AddClientRolesToUser), arg0, arg1, arg2, arg3, arg4)
}

// CreateClientRole mocks base method
func (m *ManagementComponent) CreateClientRole(arg0 context.Context, arg1, arg2 string, arg3 management.RoleRepresentation) (string, error) {
	ret := m.ctrl.Call(m, "CreateClientRole", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClientRole indicates an expected call of CreateClientRole
func (mr *ManagementComponentMockRecorder) CreateClientRole(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClientRole", reflect.TypeOf((*ManagementComponent)(nil).CreateClientRole), arg0, arg1, arg2, arg3)
}

// CreateUser mocks base method
func (m *ManagementComponent) CreateUser(arg0 context.Context, arg1 string, arg2 management.UserRepresentation) (string, error) {
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *ManagementComponentMockRecorder) CreateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*ManagementComponent)(nil).CreateUser), arg0, arg1, arg2)
}

// DeleteCredentialsForUser mocks base method
func (m *ManagementComponent) DeleteCredentialsForUser(arg0 context.Context, arg1, arg2, arg3 string) error {
	ret := m.ctrl.Call(m, "DeleteCredentialsForUser", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCredentialsForUser indicates an expected call of DeleteCredentialsForUser
func (mr *ManagementComponentMockRecorder) DeleteCredentialsForUser(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCredentialsForUser", reflect.TypeOf((*ManagementComponent)(nil).DeleteCredentialsForUser), arg0, arg1, arg2, arg3)
}

// DeleteUser mocks base method
func (m *ManagementComponent) DeleteUser(arg0 context.Context, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *ManagementComponentMockRecorder) DeleteUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*ManagementComponent)(nil).DeleteUser), arg0, arg1, arg2)
}

// GetClient mocks base method
func (m *ManagementComponent) GetClient(arg0 context.Context, arg1, arg2 string) (management.ClientRepresentation, error) {
	ret := m.ctrl.Call(m, "GetClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(management.ClientRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClient indicates an expected call of GetClient
func (mr *ManagementComponentMockRecorder) GetClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*ManagementComponent)(nil).GetClient), arg0, arg1, arg2)
}

// GetClientRoles mocks base method
func (m *ManagementComponent) GetClientRoles(arg0 context.Context, arg1, arg2 string) ([]management.RoleRepresentation, error) {
	ret := m.ctrl.Call(m, "GetClientRoles", arg0, arg1, arg2)
	ret0, _ := ret[0].([]management.RoleRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientRoles indicates an expected call of GetClientRoles
func (mr *ManagementComponentMockRecorder) GetClientRoles(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientRoles", reflect.TypeOf((*ManagementComponent)(nil).GetClientRoles), arg0, arg1, arg2)
}

// GetClientRolesForUser mocks base method
func (m *ManagementComponent) GetClientRolesForUser(arg0 context.Context, arg1, arg2, arg3 string) ([]management.RoleRepresentation, error) {
	ret := m.ctrl.Call(m, "GetClientRolesForUser", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]management.RoleRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientRolesForUser indicates an expected call of GetClientRolesForUser
func (mr *ManagementComponentMockRecorder) GetClientRolesForUser(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientRolesForUser", reflect.TypeOf((*ManagementComponent)(nil).GetClientRolesForUser), arg0, arg1, arg2, arg3)
}

// GetClients mocks base method
func (m *ManagementComponent) GetClients(arg0 context.Context, arg1 string) ([]management.ClientRepresentation, error) {
	ret := m.ctrl.Call(m, "GetClients", arg0, arg1)
	ret0, _ := ret[0].([]management.ClientRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClients indicates an expected call of GetClients
func (mr *ManagementComponentMockRecorder) GetClients(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClients", reflect.TypeOf((*ManagementComponent)(nil).GetClients), arg0, arg1)
}

// GetCredentialsForUser mocks base method
func (m *ManagementComponent) GetCredentialsForUser(arg0 context.Context, arg1, arg2 string) ([]management.CredentialRepresentation, error) {
	ret := m.ctrl.Call(m, "GetCredentialsForUser", arg0, arg1, arg2)
	ret0, _ := ret[0].([]management.CredentialRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialsForUser indicates an expected call of GetCredentialsForUser
func (mr *ManagementComponentMockRecorder) GetCredentialsForUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsForUser", reflect.TypeOf((*ManagementComponent)(nil).GetCredentialsForUser), arg0, arg1, arg2)
}

// GetRealm mocks base method
func (m *ManagementComponent) GetRealm(arg0 context.Context, arg1 string) (management.RealmRepresentation, error) {
	ret := m.ctrl.Call(m, "GetRealm", arg0, arg1)
	ret0, _ := ret[0].(management.RealmRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRealm indicates an expected call of GetRealm
func (mr *ManagementComponentMockRecorder) GetRealm(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRealm", reflect.TypeOf((*ManagementComponent)(nil).GetRealm), arg0, arg1)
}

// GetRealmRolesForUser mocks base method
func (m *ManagementComponent) GetRealmRolesForUser(arg0 context.Context, arg1, arg2 string) ([]management.RoleRepresentation, error) {
	ret := m.ctrl.Call(m, "GetRealmRolesForUser", arg0, arg1, arg2)
	ret0, _ := ret[0].([]management.RoleRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRealmRolesForUser indicates an expected call of GetRealmRolesForUser
func (mr *ManagementComponentMockRecorder) GetRealmRolesForUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRealmRolesForUser", reflect.TypeOf((*ManagementComponent)(nil).GetRealmRolesForUser), arg0, arg1, arg2)
}

// GetRealms mocks base method
func (m *ManagementComponent) GetRealms(arg0 context.Context) ([]management.RealmRepresentation, error) {
	ret := m.ctrl.Call(m, "GetRealms", arg0)
	ret0, _ := ret[0].([]management.RealmRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRealms indicates an expected call of GetRealms
func (mr *ManagementComponentMockRecorder) GetRealms(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRealms", reflect.TypeOf((*ManagementComponent)(nil).GetRealms), arg0)
}

// GetRole mocks base method
func (m *ManagementComponent) GetRole(arg0 context.Context, arg1, arg2 string) (management.RoleRepresentation, error) {
	ret := m.ctrl.Call(m, "GetRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(management.RoleRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole
func (mr *ManagementComponentMockRecorder) GetRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*ManagementComponent)(nil).GetRole), arg0, arg1, arg2)
}

// GetRoles mocks base method
func (m *ManagementComponent) GetRoles(arg0 context.Context, arg1 string) ([]management.RoleRepresentation, error) {
	ret := m.ctrl.Call(m, "GetRoles", arg0, arg1)
	ret0, _ := ret[0].([]management.RoleRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoles indicates an expected call of GetRoles
func (mr *ManagementComponentMockRecorder) GetRoles(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoles", reflect.TypeOf((*ManagementComponent)(nil).GetRoles), arg0, arg1)
}

// GetUser mocks base method
func (m *ManagementComponent) GetUser(arg0 context.Context, arg1, arg2 string) (management.UserRepresentation, error) {
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(management.UserRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *ManagementComponentMockRecorder) GetUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*ManagementComponent)(nil).GetUser), arg0, arg1, arg2)
}

// GetUserAccountStatus mocks base method
func (m *ManagementComponent) GetUserAccountStatus(arg0 context.Context, arg1, arg2 string) (map[string]bool, error) {
	ret := m.ctrl.Call(m, "GetUserAccountStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAccountStatus indicates an expected call of GetUserAccountStatus
func (mr *ManagementComponentMockRecorder) GetUserAccountStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAccountStatus", reflect.TypeOf((*ManagementComponent)(nil).GetUserAccountStatus), arg0, arg1, arg2)
}

// GetUsers mocks base method
func (m *ManagementComponent) GetUsers(arg0 context.Context, arg1, arg2 string, arg3 ...string) ([]management.UserRepresentation, error) {
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsers", varargs...)
	ret0, _ := ret[0].([]management.UserRepresentation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers
func (mr *ManagementComponentMockRecorder) GetUsers(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*ManagementComponent)(nil).GetUsers), varargs...)
}

// ResetPassword mocks base method
func (m *ManagementComponent) ResetPassword(arg0 context.Context, arg1, arg2 string, arg3 management.PasswordRepresentation) error {
	ret := m.ctrl.Call(m, "ResetPassword", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetPassword indicates an expected call of ResetPassword
func (mr *ManagementComponentMockRecorder) ResetPassword(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*ManagementComponent)(nil).ResetPassword), arg0, arg1, arg2, arg3)
}

// SendVerifyEmail mocks base method
func (m *ManagementComponent) SendVerifyEmail(arg0 context.Context, arg1, arg2 string, arg3 ...string) error {
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendVerifyEmail", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerifyEmail indicates an expected call of SendVerifyEmail
func (mr *ManagementComponentMockRecorder) SendVerifyEmail(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerifyEmail", reflect.TypeOf((*ManagementComponent)(nil).SendVerifyEmail), varargs...)
}

// UpdateUser mocks base method
func (m *ManagementComponent) UpdateUser(arg0 context.Context, arg1, arg2 string, arg3 management.UserRepresentation) error {
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *ManagementComponentMockRecorder) UpdateUser(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*ManagementComponent)(nil).UpdateUser), arg0, arg1, arg2, arg3)
}
