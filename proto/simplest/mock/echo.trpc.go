// Code generated by MockGen. DO NOT EDIT.
// Source: echo.trpc.go
//
// Generated by this command:
//
//	mockgen -source=echo.trpc.go -destination=mock/echo.trpc.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	simplest "github.com/Andrew-M-C/trpc-go-demo/proto/simplest"
	gomock "go.uber.org/mock/gomock"
	client "trpc.group/trpc-go/trpc-go/client"
)

// MockHelloWorldService is a mock of HelloWorldService interface.
type MockHelloWorldService struct {
	ctrl     *gomock.Controller
	recorder *MockHelloWorldServiceMockRecorder
}

// MockHelloWorldServiceMockRecorder is the mock recorder for MockHelloWorldService.
type MockHelloWorldServiceMockRecorder struct {
	mock *MockHelloWorldService
}

// NewMockHelloWorldService creates a new mock instance.
func NewMockHelloWorldService(ctrl *gomock.Controller) *MockHelloWorldService {
	mock := &MockHelloWorldService{ctrl: ctrl}
	mock.recorder = &MockHelloWorldServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelloWorldService) EXPECT() *MockHelloWorldServiceMockRecorder {
	return m.recorder
}

// Hello mocks base method.
func (m *MockHelloWorldService) Hello(ctx context.Context, req *simplest.HelloRequest) (*simplest.HelloResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hello", ctx, req)
	ret0, _ := ret[0].(*simplest.HelloResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hello indicates an expected call of Hello.
func (mr *MockHelloWorldServiceMockRecorder) Hello(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockHelloWorldService)(nil).Hello), ctx, req)
}

// MockHelloWorldClientProxy is a mock of HelloWorldClientProxy interface.
type MockHelloWorldClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockHelloWorldClientProxyMockRecorder
}

// MockHelloWorldClientProxyMockRecorder is the mock recorder for MockHelloWorldClientProxy.
type MockHelloWorldClientProxyMockRecorder struct {
	mock *MockHelloWorldClientProxy
}

// NewMockHelloWorldClientProxy creates a new mock instance.
func NewMockHelloWorldClientProxy(ctrl *gomock.Controller) *MockHelloWorldClientProxy {
	mock := &MockHelloWorldClientProxy{ctrl: ctrl}
	mock.recorder = &MockHelloWorldClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelloWorldClientProxy) EXPECT() *MockHelloWorldClientProxyMockRecorder {
	return m.recorder
}

// Hello mocks base method.
func (m *MockHelloWorldClientProxy) Hello(ctx context.Context, req *simplest.HelloRequest, opts ...client.Option) (*simplest.HelloResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Hello", varargs...)
	ret0, _ := ret[0].(*simplest.HelloResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hello indicates an expected call of Hello.
func (mr *MockHelloWorldClientProxyMockRecorder) Hello(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockHelloWorldClientProxy)(nil).Hello), varargs...)
}