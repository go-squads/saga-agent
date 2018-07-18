package lxdclient

import (
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

func createContainer(req api.ContainersPost) (op lxd.Operation, err error) {
	client, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		return nil, err
	}

	op, err = client.CreateContainer(req)
	if err != nil {
		return nil, err
	}

	err = op.Wait()
	if err != nil {
		return nil, err
	}

	return op, nil
}

func getContainer(name string) (container *api.Container, err error) {
	client, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		return nil, err
	}

	container, _, err = client.GetContainer(name)
	return container, nil
}
