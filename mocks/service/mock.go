// Code generated by MockGen. DO NOT EDIT.
// Source: service/service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	model "bootcamp/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// CreateCustomer mocks base method.
func (m *MockIService) CreateCustomer(customer *model.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomer", customer)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCustomer indicates an expected call of CreateCustomer.
func (mr *MockIServiceMockRecorder) CreateCustomer(customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*MockIService)(nil).CreateCustomer), customer)
}

// DescreaseBalance mocks base method.
func (m *MockIService) DescreaseBalance(customer *model.Customer, amount float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescreaseBalance", customer, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescreaseBalance indicates an expected call of DescreaseBalance.
func (mr *MockIServiceMockRecorder) DescreaseBalance(customer, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescreaseBalance", reflect.TypeOf((*MockIService)(nil).DescreaseBalance), customer, amount)
}

// GetAllCustomers mocks base method.
func (m *MockIService) GetAllCustomers() *[]model.Customer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCustomers")
	ret0, _ := ret[0].(*[]model.Customer)
	return ret0
}

// GetAllCustomers indicates an expected call of GetAllCustomers.
func (mr *MockIServiceMockRecorder) GetAllCustomers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCustomers", reflect.TypeOf((*MockIService)(nil).GetAllCustomers))
}

// GetCustomerByUsername mocks base method.
func (m *MockIService) GetCustomerByUsername(username string) (*model.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomerByUsername", username)
	ret0, _ := ret[0].(*model.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomerByUsername indicates an expected call of GetCustomerByUsername.
func (mr *MockIServiceMockRecorder) GetCustomerByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomerByUsername", reflect.TypeOf((*MockIService)(nil).GetCustomerByUsername), username)
}

// IncreaseBalance mocks base method.
func (m *MockIService) IncreaseBalance(customer *model.Customer, amount float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseBalance", customer, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseBalance indicates an expected call of IncreaseBalance.
func (mr *MockIServiceMockRecorder) IncreaseBalance(customer, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseBalance", reflect.TypeOf((*MockIService)(nil).IncreaseBalance), customer, amount)
}

// UpdateCustomer mocks base method.
func (m *MockIService) UpdateCustomer(customer *model.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCustomer", customer)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCustomer indicates an expected call of UpdateCustomer.
func (mr *MockIServiceMockRecorder) UpdateCustomer(customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomer", reflect.TypeOf((*MockIService)(nil).UpdateCustomer), customer)
}
