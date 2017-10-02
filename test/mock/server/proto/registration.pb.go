// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/api/registration (interfaces: RegistrationClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	registration "github.com/spiffe/spire/proto/api/registration"
	common "github.com/spiffe/spire/proto/common"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// MockRegistrationClient is a mock of RegistrationClient interface
type MockRegistrationClient struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrationClientMockRecorder
}

// MockRegistrationClientMockRecorder is the mock recorder for MockRegistrationClient
type MockRegistrationClientMockRecorder struct {
	mock *MockRegistrationClient
}

// NewMockRegistrationClient creates a new mock instance
func NewMockRegistrationClient(ctrl *gomock.Controller) *MockRegistrationClient {
	mock := &MockRegistrationClient{ctrl: ctrl}
	mock.recorder = &MockRegistrationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistrationClient) EXPECT() *MockRegistrationClientMockRecorder {
	return m.recorder
}

// CreateEntry mocks base method
func (m *MockRegistrationClient) CreateEntry(arg0 context.Context, arg1 *common.RegistrationEntry, arg2 ...grpc.CallOption) (*registration.RegistrationEntryID, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEntry", varargs...)
	ret0, _ := ret[0].(*registration.RegistrationEntryID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry
func (mr *MockRegistrationClientMockRecorder) CreateEntry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockRegistrationClient)(nil).CreateEntry), varargs...)
}

// CreateFederatedBundle mocks base method
func (m *MockRegistrationClient) CreateFederatedBundle(arg0 context.Context, arg1 *registration.CreateFederatedBundleRequest, arg2 ...grpc.CallOption) (*common.Empty, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateFederatedBundle", varargs...)
	ret0, _ := ret[0].(*common.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFederatedBundle indicates an expected call of CreateFederatedBundle
func (mr *MockRegistrationClientMockRecorder) CreateFederatedBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFederatedBundle", reflect.TypeOf((*MockRegistrationClient)(nil).CreateFederatedBundle), varargs...)
}

// DeleteEntry mocks base method
func (m *MockRegistrationClient) DeleteEntry(arg0 context.Context, arg1 *registration.RegistrationEntryID, arg2 ...grpc.CallOption) (*common.RegistrationEntry, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteEntry", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEntry indicates an expected call of DeleteEntry
func (mr *MockRegistrationClientMockRecorder) DeleteEntry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntry", reflect.TypeOf((*MockRegistrationClient)(nil).DeleteEntry), varargs...)
}

// DeleteFederatedBundle mocks base method
func (m *MockRegistrationClient) DeleteFederatedBundle(arg0 context.Context, arg1 *registration.FederatedSpiffeID, arg2 ...grpc.CallOption) (*common.Empty, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFederatedBundle", varargs...)
	ret0, _ := ret[0].(*common.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFederatedBundle indicates an expected call of DeleteFederatedBundle
func (mr *MockRegistrationClientMockRecorder) DeleteFederatedBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFederatedBundle", reflect.TypeOf((*MockRegistrationClient)(nil).DeleteFederatedBundle), varargs...)
}

// FetchEntry mocks base method
func (m *MockRegistrationClient) FetchEntry(arg0 context.Context, arg1 *registration.RegistrationEntryID, arg2 ...grpc.CallOption) (*common.RegistrationEntry, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchEntry", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchEntry indicates an expected call of FetchEntry
func (mr *MockRegistrationClientMockRecorder) FetchEntry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchEntry", reflect.TypeOf((*MockRegistrationClient)(nil).FetchEntry), varargs...)
}

// ListByParentID mocks base method
func (m *MockRegistrationClient) ListByParentID(arg0 context.Context, arg1 *registration.ParentID, arg2 ...grpc.CallOption) (*common.RegistrationEntries, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByParentID", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByParentID indicates an expected call of ListByParentID
func (mr *MockRegistrationClientMockRecorder) ListByParentID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByParentID", reflect.TypeOf((*MockRegistrationClient)(nil).ListByParentID), varargs...)
}

// ListBySelector mocks base method
func (m *MockRegistrationClient) ListBySelector(arg0 context.Context, arg1 *common.Selector, arg2 ...grpc.CallOption) (*common.RegistrationEntries, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBySelector", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySelector indicates an expected call of ListBySelector
func (mr *MockRegistrationClientMockRecorder) ListBySelector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySelector", reflect.TypeOf((*MockRegistrationClient)(nil).ListBySelector), varargs...)
}

// ListBySpiffeID mocks base method
func (m *MockRegistrationClient) ListBySpiffeID(arg0 context.Context, arg1 *registration.SpiffeID, arg2 ...grpc.CallOption) (*common.RegistrationEntries, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBySpiffeID", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySpiffeID indicates an expected call of ListBySpiffeID
func (mr *MockRegistrationClientMockRecorder) ListBySpiffeID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySpiffeID", reflect.TypeOf((*MockRegistrationClient)(nil).ListBySpiffeID), varargs...)
}

// ListFederatedBundles mocks base method
func (m *MockRegistrationClient) ListFederatedBundles(arg0 context.Context, arg1 *common.Empty, arg2 ...grpc.CallOption) (*registration.ListFederatedBundlesReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFederatedBundles", varargs...)
	ret0, _ := ret[0].(*registration.ListFederatedBundlesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFederatedBundles indicates an expected call of ListFederatedBundles
func (mr *MockRegistrationClientMockRecorder) ListFederatedBundles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFederatedBundles", reflect.TypeOf((*MockRegistrationClient)(nil).ListFederatedBundles), varargs...)
}

// UpdateEntry mocks base method
func (m *MockRegistrationClient) UpdateEntry(arg0 context.Context, arg1 *registration.UpdateEntryRequest, arg2 ...grpc.CallOption) (*common.RegistrationEntry, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEntry", varargs...)
	ret0, _ := ret[0].(*common.RegistrationEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntry indicates an expected call of UpdateEntry
func (mr *MockRegistrationClientMockRecorder) UpdateEntry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntry", reflect.TypeOf((*MockRegistrationClient)(nil).UpdateEntry), varargs...)
}

// UpdateFederatedBundle mocks base method
func (m *MockRegistrationClient) UpdateFederatedBundle(arg0 context.Context, arg1 *registration.FederatedBundle, arg2 ...grpc.CallOption) (*common.Empty, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFederatedBundle", varargs...)
	ret0, _ := ret[0].(*common.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFederatedBundle indicates an expected call of UpdateFederatedBundle
func (mr *MockRegistrationClientMockRecorder) UpdateFederatedBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFederatedBundle", reflect.TypeOf((*MockRegistrationClient)(nil).UpdateFederatedBundle), varargs...)
}