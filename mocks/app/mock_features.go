// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/app (interfaces: MigratorFeature,SuppressUserFeature)

// Package mock_app is a generated GoMock package.
package mock_app

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	backendconfig "github.com/rudderlabs/rudder-server/config/backend-config"
	jobsdb "github.com/rudderlabs/rudder-server/jobsdb"
	types "github.com/rudderlabs/rudder-server/utils/types"
)

// MockMigratorFeature is a mock of MigratorFeature interface.
type MockMigratorFeature struct {
	ctrl     *gomock.Controller
	recorder *MockMigratorFeatureMockRecorder
}

// MockMigratorFeatureMockRecorder is the mock recorder for MockMigratorFeature.
type MockMigratorFeatureMockRecorder struct {
	mock *MockMigratorFeature
}

// NewMockMigratorFeature creates a new mock instance.
func NewMockMigratorFeature(ctrl *gomock.Controller) *MockMigratorFeature {
	mock := &MockMigratorFeature{ctrl: ctrl}
	mock.recorder = &MockMigratorFeatureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigratorFeature) EXPECT() *MockMigratorFeatureMockRecorder {
	return m.recorder
}

// PrepareJobsdbsForImport mocks base method.
func (m *MockMigratorFeature) PrepareJobsdbsForImport(arg0, arg1, arg2 *jobsdb.HandleT) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PrepareJobsdbsForImport", arg0, arg1, arg2)
}

// PrepareJobsdbsForImport indicates an expected call of PrepareJobsdbsForImport.
func (mr *MockMigratorFeatureMockRecorder) PrepareJobsdbsForImport(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareJobsdbsForImport", reflect.TypeOf((*MockMigratorFeature)(nil).PrepareJobsdbsForImport), arg0, arg1, arg2)
}

// Setup mocks base method.
func (m *MockMigratorFeature) Setup(arg0, arg1, arg2 *jobsdb.HandleT, arg3, arg4 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Setup", arg0, arg1, arg2, arg3, arg4)
}

// Setup indicates an expected call of Setup.
func (mr *MockMigratorFeatureMockRecorder) Setup(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockMigratorFeature)(nil).Setup), arg0, arg1, arg2, arg3, arg4)
}

// MockSuppressUserFeature is a mock of SuppressUserFeature interface.
type MockSuppressUserFeature struct {
	ctrl     *gomock.Controller
	recorder *MockSuppressUserFeatureMockRecorder
}

// MockSuppressUserFeatureMockRecorder is the mock recorder for MockSuppressUserFeature.
type MockSuppressUserFeatureMockRecorder struct {
	mock *MockSuppressUserFeature
}

// NewMockSuppressUserFeature creates a new mock instance.
func NewMockSuppressUserFeature(ctrl *gomock.Controller) *MockSuppressUserFeature {
	mock := &MockSuppressUserFeature{ctrl: ctrl}
	mock.recorder = &MockSuppressUserFeatureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSuppressUserFeature) EXPECT() *MockSuppressUserFeatureMockRecorder {
	return m.recorder
}

// Setup mocks base method.
func (m *MockSuppressUserFeature) Setup(arg0 backendconfig.BackendConfig) types.SuppressUserI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup", arg0)
	ret0, _ := ret[0].(types.SuppressUserI)
	return ret0
}

// Setup indicates an expected call of Setup.
func (mr *MockSuppressUserFeatureMockRecorder) Setup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockSuppressUserFeature)(nil).Setup), arg0)
}
