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

	return op, nil
}
