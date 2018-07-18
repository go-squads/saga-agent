package lxdclient

import (
	"testing"

	"github.com/lxc/lxd/shared/api"
	"github.com/stretchr/testify/assert"
)

func TestCreateContainer(t *testing.T) {
	req := api.ContainersPost{
		Name: "test-container",
		Source: api.ContainerSource{
			Type:     "image",
			Protocol: "simplestreams",
			Server:   "https://cloud-images.ubuntu.com/daily",
			Alias:    "16.04",
		},
	}
	op, err := createContainer(req)
	if assert.NotNil(t, err) && assert.Equal(t, err.Error(), "Get http://unix.socket/1.0: dial unix /var/lib/lxd/unix.socket: connect: no such file or directory", "They should be equal") {
		return
	}

	if assert.NotNil(t, op) {
		assert.Equal(t, op.Get().Status, "Running", "They should be equal")
	}
}
