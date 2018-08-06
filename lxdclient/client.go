package lxdclient

import (
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

type Client interface {
	init()
	createContainer(req api.ContainersPost) (op lxd.Operation, err error)
	getContainer(name string) (container *api.Container, err error)
	deleteContainer(name string) (op lxd.Operation, err error)
	getContainers() (containers []api.Container, err error)
	getOperationInfo(ID string) (op *api.Operation, err error)
	updateContainerState(name string, state api.ContainerStatePut) (op lxd.Operation, err error)
}

type LxdClient struct {
	client lxd.ContainerServer
}

func (l *LxdClient) init() {
	var err error
	l.client, err = lxd.ConnectLXDUnix("", nil)
	if err != nil {
		panic(err)
	}
}

func (l *LxdClient) createContainer(req api.ContainersPost) (op lxd.Operation, err error) {
	op, err = l.client.CreateContainer(req)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (l *LxdClient) getContainer(name string) (container *api.Container, err error) {
	container, _, err = l.client.GetContainer(name)
	return container, err
}

func (l *LxdClient) deleteContainer(name string) (op lxd.Operation, err error) {
	op, err = l.client.DeleteContainer(name)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (l *LxdClient) getContainers() (containers []api.Container, err error) {
	return l.client.GetContainers()
}

func (l *LxdClient) getOperationInfo(ID string) (op *api.Operation, err error) {
	op, _, err = l.client.GetOperation(ID)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func (l *LxdClient) updateContainerState(name string, state api.ContainerStatePut) (op lxd.Operation, err error) {
	return l.client.UpdateContainerState(name, state, "")
}
