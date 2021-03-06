// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_lxdclient is a generated GoMock package.
package mock_lxdclient

import (
	gomock "github.com/golang/mock/gomock"
	client "github.com/lxc/lxd/client"
	api "github.com/lxc/lxd/shared/api"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockClient) Init() {
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init
func (mr *MockClientMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockClient)(nil).Init))
}

// CreateContainer mocks base method
func (m *MockClient) CreateContainer(req api.ContainersPost) (client.Operation, error) {
	ret := m.ctrl.Call(m, "CreateContainer", req)
	ret0, _ := ret[0].(client.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer
func (mr *MockClientMockRecorder) CreateContainer(req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockClient)(nil).CreateContainer), req)
}

// GetContainer mocks base method
func (m *MockClient) GetContainer(name string) (*api.Container, error) {
	ret := m.ctrl.Call(m, "GetContainer", name)
	ret0, _ := ret[0].(*api.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainer indicates an expected call of GetContainer
func (mr *MockClientMockRecorder) GetContainer(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainer", reflect.TypeOf((*MockClient)(nil).GetContainer), name)
}

// DeleteContainer mocks base method
func (m *MockClient) DeleteContainer(name string) (client.Operation, error) {
	ret := m.ctrl.Call(m, "DeleteContainer", name)
	ret0, _ := ret[0].(client.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteContainer indicates an expected call of DeleteContainer
func (mr *MockClientMockRecorder) DeleteContainer(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainer", reflect.TypeOf((*MockClient)(nil).DeleteContainer), name)
}

// GetContainers mocks base method
func (m *MockClient) GetContainers() ([]api.Container, error) {
	ret := m.ctrl.Call(m, "GetContainers")
	ret0, _ := ret[0].([]api.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainers indicates an expected call of GetContainers
func (mr *MockClientMockRecorder) GetContainers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainers", reflect.TypeOf((*MockClient)(nil).GetContainers))
}

// GetOperationInfo mocks base method
func (m *MockClient) GetOperationInfo(ID string) (*api.Operation, error) {
	ret := m.ctrl.Call(m, "GetOperationInfo", ID)
	ret0, _ := ret[0].(*api.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperationInfo indicates an expected call of GetOperationInfo
func (mr *MockClientMockRecorder) GetOperationInfo(ID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperationInfo", reflect.TypeOf((*MockClient)(nil).GetOperationInfo), ID)
}

// UpdateContainerState mocks base method
func (m *MockClient) UpdateContainerState(name string, state api.ContainerStatePut) (client.Operation, error) {
	ret := m.ctrl.Call(m, "UpdateContainerState", name, state)
	ret0, _ := ret[0].(client.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateContainerState indicates an expected call of UpdateContainerState
func (mr *MockClientMockRecorder) UpdateContainerState(name, state interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContainerState", reflect.TypeOf((*MockClient)(nil).UpdateContainerState), name, state)
}
