// Code generated by MockGen. DO NOT EDIT.
// Source: escort-book-delete-customers/repositories (interfaces: IBankAccountRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockIBankAccountRepository is a mock of IBankAccountRepository interface.
type MockIBankAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIBankAccountRepositoryMockRecorder
}

// MockIBankAccountRepositoryMockRecorder is the mock recorder for MockIBankAccountRepository.
type MockIBankAccountRepositoryMockRecorder struct {
	mock *MockIBankAccountRepository
}

// NewMockIBankAccountRepository creates a new mock instance.
func NewMockIBankAccountRepository(ctrl *gomock.Controller) *MockIBankAccountRepository {
	mock := &MockIBankAccountRepository{ctrl: ctrl}
	mock.recorder = &MockIBankAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBankAccountRepository) EXPECT() *MockIBankAccountRepositoryMockRecorder {
	return m.recorder
}

// DeleteCustomerBankAccounts mocks base method.
func (m *MockIBankAccountRepository) DeleteCustomerBankAccounts(arg0 context.Context, arg1 primitive.M) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCustomerBankAccounts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCustomerBankAccounts indicates an expected call of DeleteCustomerBankAccounts.
func (mr *MockIBankAccountRepositoryMockRecorder) DeleteCustomerBankAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCustomerBankAccounts", reflect.TypeOf((*MockIBankAccountRepository)(nil).DeleteCustomerBankAccounts), arg0, arg1)
}

// DeleteEscortBankAccounts mocks base method.
func (m *MockIBankAccountRepository) DeleteEscortBankAccounts(arg0 context.Context, arg1 primitive.M) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEscortBankAccounts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEscortBankAccounts indicates an expected call of DeleteEscortBankAccounts.
func (mr *MockIBankAccountRepositoryMockRecorder) DeleteEscortBankAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEscortBankAccounts", reflect.TypeOf((*MockIBankAccountRepository)(nil).DeleteEscortBankAccounts), arg0, arg1)
}
