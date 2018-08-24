package lxdclient

import (
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

// Client ...
type Client interface {
	Init()
	CreateContainer(req api.ContainersPost) (op lxd.Operation, err error)
	GetContainer(name string) (container *api.Container, err error)
	DeleteContainer(name string) (op lxd.Operation, err error)
	GetContainers() (containers []api.Container, err error)
	GetOperationInfo(ID string) (op *api.Operation, err error)
	UpdateContainerState(name string, state api.ContainerStatePut) (op lxd.Operation, err error)
	GetContainerState(string) (*api.ContainerState, string, error)
}

// LxdClient ...
type LxdClient struct {
	ContainerServer lxd.ContainerServer
}

// Init ...
func (l *LxdClient) Init() {
	var err error
	l.ContainerServer, err = lxd.ConnectLXDUnix("", nil)
	if err != nil {
		panic(err)
	}
}

// CreateContainer ...
func (l *LxdClient) CreateContainer(req api.ContainersPost) (op lxd.Operation, err error) {
	op, err = l.ContainerServer.CreateContainer(req)
	if err != nil {
		return nil, err
	}
	return op, nil
}

// GetContainer ...
func (l *LxdClient) GetContainer(name string) (container *api.Container, err error) {
	container, _, err = l.ContainerServer.GetContainer(name)
	return container, err
}

// DeleteContainer ...
func (l *LxdClient) DeleteContainer(name string) (op lxd.Operation, err error) {
	op, err = l.ContainerServer.DeleteContainer(name)
	if err != nil {
		return nil, err
	}
	return op, nil
}

// GetContainers ...
func (l *LxdClient) GetContainers() (containers []api.Container, err error) {
	return l.ContainerServer.GetContainers()
}

// GetOperationInfo ...
func (l *LxdClient) GetOperationInfo(ID string) (op *api.Operation, err error) {
	op, _, err = l.ContainerServer.GetOperation(ID)
	if err != nil {
		return nil, err
	}
	return op, nil
}

// UpdateContainerState ...
func (l *LxdClient) UpdateContainerState(name string, state api.ContainerStatePut) (op lxd.Operation, err error) {
	return l.ContainerServer.UpdateContainerState(name, state, "")
}

func (l *LxdClient) GetContainerState(containerName string) (*api.ContainerState, string, error) {
	return l.ContainerServer.GetContainerState(containerName)
}
