// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package hw5mocks is a generated GoMock package.
package hw5mocks

import (
	context "context"
	reflect "reflect"

	hw5 "gitee.com/SeanJPK/go-basic/basic/hw/hw5"
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// FindOrCreate mocks base method.
func (m *MockUserService) FindOrCreate(ctx context.Context, phone string) (hw5.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrCreate", ctx, phone)
	ret0, _ := ret[0].(hw5.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrCreate indicates an expected call of FindOrCreate.
func (mr *MockUserServiceMockRecorder) FindOrCreate(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrCreate", reflect.TypeOf((*MockUserService)(nil).FindOrCreate), ctx, phone)
}

// FindOrCreateByWechat mocks base method.
func (m *MockUserService) FindOrCreateByWechat(ctx context.Context, wechatInfo hw5.WechatInfo) (hw5.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrCreateByWechat", ctx, wechatInfo)
	ret0, _ := ret[0].(hw5.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrCreateByWechat indicates an expected call of FindOrCreateByWechat.
func (mr *MockUserServiceMockRecorder) FindOrCreateByWechat(ctx, wechatInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrCreateByWechat", reflect.TypeOf((*MockUserService)(nil).FindOrCreateByWechat), ctx, wechatInfo)
}

// Login mocks base method.
func (m *MockUserService) Login(ctx context.Context, email, password string) (hw5.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, email, password)
	ret0, _ := ret[0].(hw5.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), ctx, email, password)
}

// Profile mocks base method.
func (m *MockUserService) Profile(ctx context.Context, id int64) (hw5.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Profile", ctx, id)
	ret0, _ := ret[0].(hw5.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Profile indicates an expected call of Profile.
func (mr *MockUserServiceMockRecorder) Profile(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Profile", reflect.TypeOf((*MockUserService)(nil).Profile), ctx, id)
}

// SignUp mocks base method.
func (m *MockUserService) SignUp(ctx context.Context, u hw5.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserServiceMockRecorder) SignUp(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUserService)(nil).SignUp), ctx, u)
}

// MockCodeService is a mock of CodeService interface.
type MockCodeService struct {
	ctrl     *gomock.Controller
	recorder *MockCodeServiceMockRecorder
}

// MockCodeServiceMockRecorder is the mock recorder for MockCodeService.
type MockCodeServiceMockRecorder struct {
	mock *MockCodeService
}

// NewMockCodeService creates a new mock instance.
func NewMockCodeService(ctrl *gomock.Controller) *MockCodeService {
	mock := &MockCodeService{ctrl: ctrl}
	mock.recorder = &MockCodeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodeService) EXPECT() *MockCodeServiceMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockCodeService) Send(ctx context.Context, biz, phone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", ctx, biz, phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockCodeServiceMockRecorder) Send(ctx, biz, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockCodeService)(nil).Send), ctx, biz, phone)
}

// Verify mocks base method.
func (m *MockCodeService) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, biz, phone, inputCode)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockCodeServiceMockRecorder) Verify(ctx, biz, phone, inputCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockCodeService)(nil).Verify), ctx, biz, phone, inputCode)
}

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// CheckSession mocks base method.
func (m *MockHandler) CheckSession(ctx *gin.Context, ssid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSession", ctx, ssid)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckSession indicates an expected call of CheckSession.
func (mr *MockHandlerMockRecorder) CheckSession(ctx, ssid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSession", reflect.TypeOf((*MockHandler)(nil).CheckSession), ctx, ssid)
}

// ClearToken mocks base method.
func (m *MockHandler) ClearToken(ctx *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearToken", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearToken indicates an expected call of ClearToken.
func (mr *MockHandlerMockRecorder) ClearToken(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearToken", reflect.TypeOf((*MockHandler)(nil).ClearToken), ctx)
}

// ExtractToken mocks base method.
func (m *MockHandler) ExtractToken(ctx *gin.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractToken", ctx)
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtractToken indicates an expected call of ExtractToken.
func (mr *MockHandlerMockRecorder) ExtractToken(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractToken", reflect.TypeOf((*MockHandler)(nil).ExtractToken), ctx)
}

// SetJWTToken mocks base method.
func (m *MockHandler) SetJWTToken(ctx *gin.Context, uid int64, ssid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetJWTToken", ctx, uid, ssid)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetJWTToken indicates an expected call of SetJWTToken.
func (mr *MockHandlerMockRecorder) SetJWTToken(ctx, uid, ssid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJWTToken", reflect.TypeOf((*MockHandler)(nil).SetJWTToken), ctx, uid, ssid)
}

// SetLoginToken mocks base method.
func (m *MockHandler) SetLoginToken(ctx *gin.Context, uid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLoginToken", ctx, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLoginToken indicates an expected call of SetLoginToken.
func (mr *MockHandlerMockRecorder) SetLoginToken(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLoginToken", reflect.TypeOf((*MockHandler)(nil).SetLoginToken), ctx, uid)
}
