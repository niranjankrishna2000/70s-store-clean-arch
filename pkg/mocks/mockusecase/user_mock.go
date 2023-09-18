// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interface/user.go

// Package mockusecase is a generated GoMock package.
package mockusecase

import (
	domain "main/pkg/domain"
	models "main/pkg/utils/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserUseCase) AddAddress(id int, address models.AddAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", id, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserUseCaseMockRecorder) AddAddress(id, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserUseCase)(nil).AddAddress), id, address)
}

// ChangePassword mocks base method.
func (m *MockUserUseCase) ChangePassword(id int, old, password, repassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", id, old, password, repassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserUseCaseMockRecorder) ChangePassword(id, old, password, repassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserUseCase)(nil).ChangePassword), id, old, password, repassword)
}

// ClearCart mocks base method.
func (m *MockUserUseCase) ClearCart(cartID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearCart", cartID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearCart indicates an expected call of ClearCart.
func (mr *MockUserUseCaseMockRecorder) ClearCart(cartID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearCart", reflect.TypeOf((*MockUserUseCase)(nil).ClearCart), cartID)
}

// EditUser mocks base method.
func (m *MockUserUseCase) EditUser(id int, userData models.EditUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUser", id, userData)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditUser indicates an expected call of EditUser.
func (mr *MockUserUseCaseMockRecorder) EditUser(id, userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUser", reflect.TypeOf((*MockUserUseCase)(nil).EditUser), id, userData)
}

// GetAddresses mocks base method.
func (m *MockUserUseCase) GetAddresses(id int) ([]domain.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddresses", id)
	ret0, _ := ret[0].([]domain.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddresses indicates an expected call of GetAddresses.
func (mr *MockUserUseCaseMockRecorder) GetAddresses(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddresses", reflect.TypeOf((*MockUserUseCase)(nil).GetAddresses), id)
}

// GetCart mocks base method.
func (m *MockUserUseCase) GetCart(id, page, limit int) ([]models.GetCart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", id, page, limit)
	ret0, _ := ret[0].([]models.GetCart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart.
func (mr *MockUserUseCaseMockRecorder) GetCart(id, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockUserUseCase)(nil).GetCart), id, page, limit)
}

// GetCartID mocks base method.
func (m *MockUserUseCase) GetCartID(userID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartID", userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartID indicates an expected call of GetCartID.
func (mr *MockUserUseCaseMockRecorder) GetCartID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartID", reflect.TypeOf((*MockUserUseCase)(nil).GetCartID), userID)
}

// GetUserDetails mocks base method.
func (m *MockUserUseCase) GetUserDetails(id int) (models.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetails", id)
	ret0, _ := ret[0].(models.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDetails indicates an expected call of GetUserDetails.
func (mr *MockUserUseCaseMockRecorder) GetUserDetails(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetails", reflect.TypeOf((*MockUserUseCase)(nil).GetUserDetails), id)
}

// GetWallet mocks base method.
func (m *MockUserUseCase) GetWallet(id, page, limit int) (models.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWallet", id, page, limit)
	ret0, _ := ret[0].(models.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWallet indicates an expected call of GetWallet.
func (mr *MockUserUseCaseMockRecorder) GetWallet(id, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWallet", reflect.TypeOf((*MockUserUseCase)(nil).GetWallet), id, page, limit)
}

// Login mocks base method.
func (m *MockUserUseCase) Login(user models.UserLogin) (models.TokenUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", user)
	ret0, _ := ret[0].(models.TokenUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserUseCaseMockRecorder) Login(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserUseCase)(nil).Login), user)
}

// RemoveFromCart mocks base method.
func (m *MockUserUseCase) RemoveFromCart(id, inventoryID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromCart", id, inventoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromCart indicates an expected call of RemoveFromCart.
func (mr *MockUserUseCaseMockRecorder) RemoveFromCart(id, inventoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromCart", reflect.TypeOf((*MockUserUseCase)(nil).RemoveFromCart), id, inventoryID)
}

// SignUp mocks base method.
func (m *MockUserUseCase) SignUp(user models.UserDetails) (models.TokenUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", user)
	ret0, _ := ret[0].(models.TokenUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserUseCaseMockRecorder) SignUp(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUserUseCase)(nil).SignUp), user)
}

// UpdateQuantityAdd mocks base method.
func (m *MockUserUseCase) UpdateQuantityAdd(id, inv_id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuantityAdd", id, inv_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuantityAdd indicates an expected call of UpdateQuantityAdd.
func (mr *MockUserUseCaseMockRecorder) UpdateQuantityAdd(id, inv_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuantityAdd", reflect.TypeOf((*MockUserUseCase)(nil).UpdateQuantityAdd), id, inv_id)
}

// UpdateQuantityLess mocks base method.
func (m *MockUserUseCase) UpdateQuantityLess(id, inv_id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuantityLess", id, inv_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuantityLess indicates an expected call of UpdateQuantityLess.
func (mr *MockUserUseCaseMockRecorder) UpdateQuantityLess(id, inv_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuantityLess", reflect.TypeOf((*MockUserUseCase)(nil).UpdateQuantityLess), id, inv_id)
}
