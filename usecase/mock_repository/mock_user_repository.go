// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/repository/user_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	model "go_worlder_system/domain/model"
	reflect "reflect"
)

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// FindByID mocks base method
func (m *MockUserRepository) FindByID(id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockUserRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserRepository)(nil).FindByID), id)
}

// FindByEmail mocks base method
func (m *MockUserRepository) FindByEmail(email string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail
func (mr *MockUserRepositoryMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), email)
}

// FindCardByUserID mocks base method
func (m *MockUserRepository) FindCardByUserID(userID string) (*model.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCardByUserID", userID)
	ret0, _ := ret[0].(*model.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCardByUserID indicates an expected call of FindCardByUserID
func (mr *MockUserRepositoryMockRecorder) FindCardByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCardByUserID", reflect.TypeOf((*MockUserRepository)(nil).FindCardByUserID), userID)
}

// FindListByLikeBrandID mocks base method
func (m *MockUserRepository) FindListByLikeBrandID(brandID string) ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByLikeBrandID", brandID)
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByLikeBrandID indicates an expected call of FindListByLikeBrandID
func (mr *MockUserRepositoryMockRecorder) FindListByLikeBrandID(brandID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByLikeBrandID", reflect.TypeOf((*MockUserRepository)(nil).FindListByLikeBrandID), brandID)
}

// Save mocks base method
func (m *MockUserRepository) Save(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockUserRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserRepository)(nil).Save), arg0)
}

// SaveProfile mocks base method
func (m *MockUserRepository) SaveProfile(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveProfile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveProfile indicates an expected call of SaveProfile
func (mr *MockUserRepositoryMockRecorder) SaveProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveProfile", reflect.TypeOf((*MockUserRepository)(nil).SaveProfile), arg0)
}

// SaveCard mocks base method
func (m *MockUserRepository) SaveCard(arg0 *model.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCard", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCard indicates an expected call of SaveCard
func (mr *MockUserRepositoryMockRecorder) SaveCard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCard", reflect.TypeOf((*MockUserRepository)(nil).SaveCard), arg0)
}

// UpdateToActive mocks base method
func (m *MockUserRepository) UpdateToActive(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToActive", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToActive indicates an expected call of UpdateToActive
func (mr *MockUserRepositoryMockRecorder) UpdateToActive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToActive", reflect.TypeOf((*MockUserRepository)(nil).UpdateToActive), arg0)
}

// UpdatePassword mocks base method
func (m *MockUserRepository) UpdatePassword(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockUserRepositoryMockRecorder) UpdatePassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserRepository)(nil).UpdatePassword), arg0)
}

// UpdateProfile mocks base method
func (m *MockUserRepository) UpdateProfile(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile
func (mr *MockUserRepositoryMockRecorder) UpdateProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserRepository)(nil).UpdateProfile), arg0)
}
