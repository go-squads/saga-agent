package lxdclient

import (
	"testing"

	"github.com/lxc/lxd/shared/api"
	"github.com/stretchr/testify/assert"
)

func TestCreateContainer(t *testing.T) {
	name := "test-container"
	_, err := getContainer(name)
	if err == nil {
		return
	}
	req := api.ContainersPost{
		Name: name,
		Source: api.ContainerSource{
			Type:     "image",
			Protocol: "simplestreams",
			Server:   "https://cloud-images.ubuntu.com/daily",
			Alias:    "16.04",
		},
	}
	op, err := createContainer(req)
	assert.Nil(t, err)

	if assert.Nil(t, op) {
		assert.Equal(t, op.Get().Status, "Running", "They should be equal")
	}
}
