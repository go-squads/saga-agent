package lxdclient

import (
	"testing"

	"github.com/lxc/lxd/shared/api"
	"github.com/stretchr/testify/suite"
)

type ContainerSuite struct {
	suite.Suite
}

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
	deleteContainer("test-container-1")
	deleteContainer("test-container-2")

	_, err := createContainer(api.ContainersPost{
		Name:   "test-container-1",
		Source: source,
	})

	if err != nil {
		panic(err)
	}

	_, err = createContainer(api.ContainersPost{
		Name:   "test-container-3",
		Source: source,
	})

	if err != nil {
		panic(err)
	}
}

func (suite *ContainerSuite) TearDownSuite() {
	deleteContainer("test-container-1")
	deleteContainer("test-container-2")
}

func (suite *ContainerSuite) TestDeleteContainerSuccessful() {
	op, err := deleteContainer("test-container-3")
	suite.Nil(err, "They should be nil")

	if suite.NotNil(op, "They should be not nil") {
		suite.Equal(api.Running, op.Get().StatusCode, "They should be equal")
	}
}

func (suite *ContainerSuite) TestCreateContainerSuccessful() {
	name := "test-container-2"
	req := api.ContainersPost{
		Name:   name,
		Source: source,
	}

	op, err := createContainer(req)
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

	op, err := createContainer(req)
	suite.Error(err, "They should be error")
	suite.Nil(op, "They should be nil")
}

func (suite *ContainerSuite) TestGetContainerSuccessful() {
	name := "test-container-1"
	container, err := getContainer(name)
	suite.NoError(err, "They should be no error")
	if suite.NotNil(container, "They should be not nil") {
		suite.Equal(name, container.Name, "They should be equal")
	}
}

func (suite *ContainerSuite) TestGetContainerFailed() {
	name := "test-container-xyz"
	container, err := getContainer(name)
	suite.Error(err, "They should be error")
	suite.Nil(container, "They should be nil")
}

func (suite *ContainerSuite) TestGetContainersSuccessful() {
	containers, err := getContainers()
	suite.NoError(err, "They should be no error")
	suite.NotEqual(0, len(containers), "They should be not equal")
}
