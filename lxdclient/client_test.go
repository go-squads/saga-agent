package lxdclient

import (
	"testing"

	"github.com/lxc/lxd/shared/api"
	"github.com/stretchr/testify/suite"
)

type ContainerSuite struct {
	suite.Suite
}

var lxdClient LxdClient
var source api.ContainerSource

func init() {
	source = api.ContainerSource{
		Type: "none",
	}
}

func TestContainerSuite(t *testing.T) {
	suite.Run(t, new(ContainerSuite))
}

func (suite *ContainerSuite) SetupSuite() {
	lxdClient.init()
	lxdClient.deleteContainer("test-container-1")
	lxdClient.deleteContainer("test-container-2")

	_, err := lxdClient.createContainer(api.ContainersPost{
		Name:   "test-container-1",
		Source: source,
	})

	if err != nil {
		panic(err)
	}

	_, err = lxdClient.createContainer(api.ContainersPost{
		Name:   "test-container-3",
		Source: source,
	})

	if err != nil {
		panic(err)
	}
}

func (suite *ContainerSuite) TearDownSuite() {
	lxdClient.deleteContainer("test-container-1")
	lxdClient.deleteContainer("test-container-2")
	lxdClient.deleteContainer("test-container-4")
}

func (suite *ContainerSuite) TestDeleteContainerSuccessful() {
	op, err := lxdClient.deleteContainer("test-container-3")
	suite.NoError(err, "They should be no error")

	if suite.NotNil(op, "They should be not nil") {
		suite.Equal(api.Running, op.Get().StatusCode, "They should be equal")
	}
}

func (suite *ContainerSuite) TestDeleteContainerFailed() {
	_, err := lxdClient.deleteContainer("")
	suite.Error(err, "They should be error")
}

func (suite *ContainerSuite) TestCreateContainerSuccessful() {
	name := "test-container-2"
	req := api.ContainersPost{
		Name:   name,
		Source: source,
	}

	op, err := lxdClient.createContainer(req)
	suite.Nil(err, "They should be nil")

	if suite.NotNil(op, "They should be not nil") {
		suite.Equal(api.Running, op.Get().StatusCode, "They should be equal")
	}
}

func (suite *ContainerSuite) TestCreateContainerFailed() {
	name := "test-container-1"
	req := api.ContainersPost{
		Name: name,
	}

	op, err := lxdClient.createContainer(req)
	suite.Error(err, "They should be error")
	suite.Nil(op, "They should be nil")
}

func (suite *ContainerSuite) TestGetContainerSuccessful() {
	name := "test-container-1"
	container, err := lxdClient.getContainer(name)
	suite.NoError(err, "They should be no error")
	if suite.NotNil(container, "They should be not nil") {
		suite.Equal(name, container.Name, "They should be equal")
	}
}

func (suite *ContainerSuite) TestGetContainerFailed() {
	name := "test-container-xyz"
	container, err := lxdClient.getContainer(name)
	suite.Error(err, "They should be error")
	suite.Nil(container, "They should be nil")
}

func (suite *ContainerSuite) TestGetContainersSuccessful() {
	containers, err := lxdClient.getContainers()
	suite.NoError(err, "They should be no error")
	suite.NotEqual(0, len(containers), "They should be not equal")
}

func (suite *ContainerSuite) TestGetInfoSuccessful() {
	name := "test-container-4"
	req := api.ContainersPost{
		Name:   name,
		Source: source,
	}
	result, _ := lxdClient.createContainer(req)
	op, err := lxdClient.getOperationInfo(result.Get().ID)
	suite.NoError(err, "They should be no error")
	suite.Equal(result.Get().ID, op.ID, "They should be equal")
}

func (suite *ContainerSuite) TestGetInfoFailed() {
	_, err := lxdClient.getOperationInfo("")
	suite.Error(err, "They should be error")
}
