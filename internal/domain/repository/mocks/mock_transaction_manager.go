// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/repository/transaction_manager.go

// Package mocksrepository is a generated GoMock package.
package mocksrepository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTransactionManager is a mock of TransactionManager interface.
type MockTransactionManager struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionManagerMockRecorder
}

// MockTransactionManagerMockRecorder is the mock recorder for MockTransactionManager.
type MockTransactionManagerMockRecorder struct {
	mock *MockTransactionManager
}

// NewMockTransactionManager creates a new mock instance.
func NewMockTransactionManager(ctrl *gomock.Controller) *MockTransactionManager {
	mock := &MockTransactionManager{ctrl: ctrl}
	mock.recorder = &MockTransactionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionManager) EXPECT() *MockTransactionManagerMockRecorder {
	return m.recorder
}

// RunInTransaction mocks base method.
func (m *MockTransactionManager) RunInTransaction(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInTransaction", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunInTransaction indicates an expected call of RunInTransaction.
func (mr *MockTransactionManagerMockRecorder) RunInTransaction(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInTransaction", reflect.TypeOf((*MockTransactionManager)(nil).RunInTransaction), ctx, fn)
}
