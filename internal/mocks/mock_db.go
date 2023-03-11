// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/AndrewSukhobok95/yagometrics.git/internal/database (interfaces: CustomDB)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	sync "sync"
	time "time"

	datastorage "github.com/AndrewSukhobok95/yagometrics.git/internal/datastorage"
	serialization "github.com/AndrewSukhobok95/yagometrics.git/internal/serialization"
	gomock "github.com/golang/mock/gomock"
)

// MockCustomDB is a mock of CustomDB interface.
type MockCustomDB struct {
	ctrl     *gomock.Controller
	recorder *MockCustomDBMockRecorder
}

// MockCustomDBMockRecorder is the mock recorder for MockCustomDB.
type MockCustomDBMockRecorder struct {
	mock *MockCustomDB
}

// NewMockCustomDB creates a new mock instance.
func NewMockCustomDB(ctrl *gomock.Controller) *MockCustomDB {
	mock := &MockCustomDB{ctrl: ctrl}
	mock.recorder = &MockCustomDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomDB) EXPECT() *MockCustomDBMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCustomDB) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockCustomDBMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCustomDB)(nil).Close))
}

// CreateTable mocks base method.
func (m *MockCustomDB) CreateTable(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateTable", arg0)
}

// CreateTable indicates an expected call of CreateTable.
func (mr *MockCustomDBMockRecorder) CreateTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTable", reflect.TypeOf((*MockCustomDB)(nil).CreateTable), arg0)
}

// PingContext mocks base method.
func (m *MockCustomDB) PingContext(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingContext", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingContext indicates an expected call of PingContext.
func (mr *MockCustomDBMockRecorder) PingContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingContext", reflect.TypeOf((*MockCustomDB)(nil).PingContext), arg0)
}

// StartWritingToDB mocks base method.
func (m *MockCustomDB) StartWritingToDB(arg0 datastorage.Storage, arg1 time.Duration, arg2 context.Context, arg3 *sync.WaitGroup) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartWritingToDB", arg0, arg1, arg2, arg3)
}

// StartWritingToDB indicates an expected call of StartWritingToDB.
func (mr *MockCustomDBMockRecorder) StartWritingToDB(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWritingToDB", reflect.TypeOf((*MockCustomDB)(nil).StartWritingToDB), arg0, arg1, arg2, arg3)
}

// UpdateDB mocks base method.
func (m *MockCustomDB) UpdateDB(arg0 datastorage.Storage, arg1 time.Duration, arg2 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDB", arg0, arg1, arg2)
}

// UpdateDB indicates an expected call of UpdateDB.
func (mr *MockCustomDBMockRecorder) UpdateDB(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDB", reflect.TypeOf((*MockCustomDB)(nil).UpdateDB), arg0, arg1, arg2)
}

// UpdateMetricInDB mocks base method.
func (m *MockCustomDB) UpdateMetricInDB(arg0 serialization.Metrics, arg1 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMetricInDB", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetricInDB indicates an expected call of UpdateMetricInDB.
func (mr *MockCustomDBMockRecorder) UpdateMetricInDB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetricInDB", reflect.TypeOf((*MockCustomDB)(nil).UpdateMetricInDB), arg0, arg1)
}