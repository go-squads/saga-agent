package lxdclient

import (
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

var client lxd.ContainerServer

func init() {
	var err error
	client, err = lxd.ConnectLXDUnix("", nil)
	if err != nil {
		panic(err)
	}
}

func createContainer(req api.ContainersPost) (op lxd.Operation, err error) {
	op, err = client.CreateContainer(req)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func getContainer(name string) (container *api.Container, err error) {
	container, _, err = client.GetContainer(name)
	return container, err
}

func deleteContainer(name string) (op lxd.Operation, err error) {
	op, err = client.DeleteContainer(name)
	if err != nil {
		return nil, err
	}
	return op, nil
}

func getContainers() (containers []api.Container, err error) {
	return client.GetContainers()
}
