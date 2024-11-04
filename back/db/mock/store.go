// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/presedo93/wedding/back/db/sqlc (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination db/mock/store.go github.com/presedo93/wedding/back/db/sqlc Store
//

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	db "github.com/presedo93/wedding/back/db/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateGuest mocks base method.
func (m *MockStore) CreateGuest(arg0 context.Context, arg1 db.CreateGuestParams) (db.Guest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGuest", arg0, arg1)
	ret0, _ := ret[0].(db.Guest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGuest indicates an expected call of CreateGuest.
func (mr *MockStoreMockRecorder) CreateGuest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGuest", reflect.TypeOf((*MockStore)(nil).CreateGuest), arg0, arg1)
}

// CreateProfile mocks base method.
func (m *MockStore) CreateProfile(arg0 context.Context, arg1 db.CreateProfileParams) (db.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfile", arg0, arg1)
	ret0, _ := ret[0].(db.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProfile indicates an expected call of CreateProfile.
func (mr *MockStoreMockRecorder) CreateProfile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockStore)(nil).CreateProfile), arg0, arg1)
}

// DeleteGuest mocks base method.
func (m *MockStore) DeleteGuest(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGuest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGuest indicates an expected call of DeleteGuest.
func (mr *MockStoreMockRecorder) DeleteGuest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGuest", reflect.TypeOf((*MockStore)(nil).DeleteGuest), arg0, arg1)
}

// DeleteProfile mocks base method.
func (m *MockStore) DeleteProfile(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfile indicates an expected call of DeleteProfile.
func (mr *MockStoreMockRecorder) DeleteProfile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockStore)(nil).DeleteProfile), arg0, arg1)
}

// DeleteUserGuest mocks base method.
func (m *MockStore) DeleteUserGuest(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserGuest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserGuest indicates an expected call of DeleteUserGuest.
func (mr *MockStoreMockRecorder) DeleteUserGuest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserGuest", reflect.TypeOf((*MockStore)(nil).DeleteUserGuest), arg0, arg1)
}

// GetGuest mocks base method.
func (m *MockStore) GetGuest(arg0 context.Context, arg1 int64) (db.Guest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGuest", arg0, arg1)
	ret0, _ := ret[0].(db.Guest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGuest indicates an expected call of GetGuest.
func (mr *MockStoreMockRecorder) GetGuest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGuest", reflect.TypeOf((*MockStore)(nil).GetGuest), arg0, arg1)
}

// GetGuests mocks base method.
func (m *MockStore) GetGuests(arg0 context.Context, arg1 db.GetGuestsParams) ([]db.Guest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGuests", arg0, arg1)
	ret0, _ := ret[0].([]db.Guest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGuests indicates an expected call of GetGuests.
func (mr *MockStoreMockRecorder) GetGuests(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGuests", reflect.TypeOf((*MockStore)(nil).GetGuests), arg0, arg1)
}

// GetProfile mocks base method.
func (m *MockStore) GetProfile(arg0 context.Context, arg1 uuid.UUID) (db.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", arg0, arg1)
	ret0, _ := ret[0].(db.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockStoreMockRecorder) GetProfile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockStore)(nil).GetProfile), arg0, arg1)
}

// GetProfiles mocks base method.
func (m *MockStore) GetProfiles(arg0 context.Context, arg1 db.GetProfilesParams) ([]db.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfiles", arg0, arg1)
	ret0, _ := ret[0].([]db.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfiles indicates an expected call of GetProfiles.
func (mr *MockStoreMockRecorder) GetProfiles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfiles", reflect.TypeOf((*MockStore)(nil).GetProfiles), arg0, arg1)
}

// GetUserGuests mocks base method.
func (m *MockStore) GetUserGuests(arg0 context.Context, arg1 uuid.UUID) ([]db.Guest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserGuests", arg0, arg1)
	ret0, _ := ret[0].([]db.Guest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserGuests indicates an expected call of GetUserGuests.
func (mr *MockStoreMockRecorder) GetUserGuests(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserGuests", reflect.TypeOf((*MockStore)(nil).GetUserGuests), arg0, arg1)
}

// UpdateGuest mocks base method.
func (m *MockStore) UpdateGuest(arg0 context.Context, arg1 db.UpdateGuestParams) (db.Guest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGuest", arg0, arg1)
	ret0, _ := ret[0].(db.Guest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGuest indicates an expected call of UpdateGuest.
func (mr *MockStoreMockRecorder) UpdateGuest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGuest", reflect.TypeOf((*MockStore)(nil).UpdateGuest), arg0, arg1)
}

// UpdateProfile mocks base method.
func (m *MockStore) UpdateProfile(arg0 context.Context, arg1 db.UpdateProfileParams) (db.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0, arg1)
	ret0, _ := ret[0].(db.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockStoreMockRecorder) UpdateProfile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockStore)(nil).UpdateProfile), arg0, arg1)
}
