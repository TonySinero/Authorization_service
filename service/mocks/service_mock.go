// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	model "stlab.itechart-group.com/go/food_delivery/authorization_service/model"
)

// MockAuthUser is a mock of AuthUser interface.
type MockAuthUser struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUserMockRecorder
}

// MockAuthUserMockRecorder is the mock recorder for MockAuthUser.
type MockAuthUserMockRecorder struct {
	mock *MockAuthUser
}

// NewMockAuthUser creates a new mock instance.
func NewMockAuthUser(ctrl *gomock.Controller) *MockAuthUser {
	mock := &MockAuthUser{ctrl: ctrl}
	mock.recorder = &MockAuthUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUser) EXPECT() *MockAuthUserMockRecorder {
	return m.recorder
}

// AddRoleToUser mocks base method.
func (m *MockAuthUser) AddRoleToUser(user *authProto.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoleToUser", user)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRoleToUser indicates an expected call of AddRoleToUser.
func (mr *MockAuthUserMockRecorder) AddRoleToUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoleToUser", reflect.TypeOf((*MockAuthUser)(nil).AddRoleToUser), user)
}

// BindRoleWithPerms mocks base method.
func (m *MockAuthUser) BindRoleWithPerms(rp *model.BindRoleWithPermission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindRoleWithPerms", rp)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindRoleWithPerms indicates an expected call of BindRoleWithPerms.
func (mr *MockAuthUserMockRecorder) BindRoleWithPerms(rp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindRoleWithPerms", reflect.TypeOf((*MockAuthUser)(nil).BindRoleWithPerms), rp)
}

// CheckRights mocks base method.
func (m *MockAuthUser) CheckRights(neededPerms []string, givenPerms string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRights", neededPerms, givenPerms)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckRights indicates an expected call of CheckRights.
func (mr *MockAuthUserMockRecorder) CheckRights(neededPerms, givenPerms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRights", reflect.TypeOf((*MockAuthUser)(nil).CheckRights), neededPerms, givenPerms)
}

// CheckRole mocks base method.
func (m *MockAuthUser) CheckRole(neededRole []string, givenRole string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRole", neededRole, givenRole)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckRole indicates an expected call of CheckRole.
func (mr *MockAuthUserMockRecorder) CheckRole(neededRole, givenRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRole", reflect.TypeOf((*MockAuthUser)(nil).CheckRole), neededRole, givenRole)
}

// CreatePermission mocks base method.
func (m *MockAuthUser) CreatePermission(permission string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePermission", permission)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePermission indicates an expected call of CreatePermission.
func (mr *MockAuthUserMockRecorder) CreatePermission(permission interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePermission", reflect.TypeOf((*MockAuthUser)(nil).CreatePermission), permission)
}

// CreateRole mocks base method.
func (m *MockAuthUser) CreateRole(role string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", role)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockAuthUserMockRecorder) CreateRole(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockAuthUser)(nil).CreateRole), role)
}

// GenerateTokensByAuthUser mocks base method.
func (m *MockAuthUser) GenerateTokensByAuthUser(user *authProto.User) (*authProto.GeneratedTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTokensByAuthUser", user)
	ret0, _ := ret[0].(*authProto.GeneratedTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTokensByAuthUser indicates an expected call of GenerateTokensByAuthUser.
func (mr *MockAuthUserMockRecorder) GenerateTokensByAuthUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTokensByAuthUser", reflect.TypeOf((*MockAuthUser)(nil).GenerateTokensByAuthUser), user)
}

// GetAllPerms mocks base method.
func (m *MockAuthUser) GetAllPerms() ([]model.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPerms")
	ret0, _ := ret[0].([]model.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPerms indicates an expected call of GetAllPerms.
func (mr *MockAuthUserMockRecorder) GetAllPerms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPerms", reflect.TypeOf((*MockAuthUser)(nil).GetAllPerms))
}

// GetAllRoles mocks base method.
func (m *MockAuthUser) GetAllRoles() ([]model.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRoles")
	ret0, _ := ret[0].([]model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRoles indicates an expected call of GetAllRoles.
func (mr *MockAuthUserMockRecorder) GetAllRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRoles", reflect.TypeOf((*MockAuthUser)(nil).GetAllRoles))
}

// GetPermsByRoleId mocks base method.
func (m *MockAuthUser) GetPermsByRoleId(id int) ([]model.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermsByRoleId", id)
	ret0, _ := ret[0].([]model.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermsByRoleId indicates an expected call of GetPermsByRoleId.
func (mr *MockAuthUserMockRecorder) GetPermsByRoleId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermsByRoleId", reflect.TypeOf((*MockAuthUser)(nil).GetPermsByRoleId), id)
}

// GetRoleById mocks base method.
func (m *MockAuthUser) GetRoleById(id int) (*model.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleById", id)
	ret0, _ := ret[0].(*model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleById indicates an expected call of GetRoleById.
func (mr *MockAuthUserMockRecorder) GetRoleById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleById", reflect.TypeOf((*MockAuthUser)(nil).GetRoleById), id)
}

// GetRoleByUserId mocks base method.
func (m *MockAuthUser) GetRoleByUserId(userId int) (*model.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleByUserId", userId)
	ret0, _ := ret[0].(*model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleByUserId indicates an expected call of GetRoleByUserId.
func (mr *MockAuthUserMockRecorder) GetRoleByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleByUserId", reflect.TypeOf((*MockAuthUser)(nil).GetRoleByUserId), userId)
}

// ParseToken mocks base method.
func (m *MockAuthUser) ParseToken(token string) (*authProto.UserRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(*authProto.UserRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthUserMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthUser)(nil).ParseToken), token)
}

// RefreshTokens mocks base method.
func (m *MockAuthUser) RefreshTokens(refreshToken string) (*authProto.GeneratedTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokens", refreshToken)
	ret0, _ := ret[0].(*authProto.GeneratedTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshTokens indicates an expected call of RefreshTokens.
func (mr *MockAuthUserMockRecorder) RefreshTokens(refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokens", reflect.TypeOf((*MockAuthUser)(nil).RefreshTokens), refreshToken)
}
