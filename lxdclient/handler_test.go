package lxdclient

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_lxdclient "github.com/go-squads/saga-agent/lxdclient/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/lxc/lxd/shared/api"
	"github.com/stretchr/testify/suite"
)

type HandlerSuite struct {
	suite.Suite
}

var handler Handler
var ctrl *gomock.Controller

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (suite *HandlerSuite) SetupSuite() {
	ctrl = gomock.NewController(suite.Suite.T())

	mockClient := mock_lxdclient.NewMockClient(ctrl)
	mockClient.EXPECT().DeleteContainer("test-container-11").Return(nil, nil)
	mockClient.EXPECT().DeleteContainer("test-container-12").Return(nil, nil)
	mockClient.EXPECT().GetContainers().Return(nil, nil)
	mockClient.EXPECT().GetContainer("test-container-11").Return(nil, nil)
	mockClient.EXPECT().CreateContainer(api.ContainersPost{
		Name: "test-container-12",
		Source: api.ContainerSource{
			Type: "none",
		},
	}).Return(nil, errors.New("err bro"))
	mockClient.EXPECT().DeleteContainer("test-container-13").Return(nil, nil)
	mockClient.EXPECT().UpdateContainerState("test-container-13", api.ContainerStatePut{
		Action:  "start",
		Timeout: 60,
	}).Return(nil, nil)

	handler = Handler{}
	handler.HandlerClient = mockClient
}

func (suite *HandlerSuite) TearDownSuite() {

}

func (suite *HandlerSuite) TestGetContainersHandler() {
	req, err := http.NewRequest("GET", "/api/v1/containers", nil)
	if err != nil {
		suite.Fail(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.GetContainersHandler)
	handler.ServeHTTP(rr, req)
	suite.Equal(http.StatusOK, rr.Code, "They should be equal")
}

func (suite *HandlerSuite) TestGetContainerHandler() {
	req, err := http.NewRequest("GET", "/api/v1/container/test-container-11", nil)
	if err != nil {
		suite.Fail(err.Error())
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/container/{name}", handler.GetContainerHandler)
	router.ServeHTTP(rr, req)
	suite.Equal(http.StatusOK, rr.Code, "They should be equal")
}

func (suite *HandlerSuite) TestCreateContainerHandler() {
	payload := []byte(`{"name":"test-container-12","type":"none"}`)
	req, err := http.NewRequest("POST", "/api/v1/container", bytes.NewBuffer(payload))
	if err != nil {
		suite.Fail(err.Error())
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/container", handler.CreateContainerHandler)
	router.ServeHTTP(rr, req)
	suite.Equal(http.StatusBadRequest, rr.Code, "They should be equal")
	suite.Equal(string(`{"error":"err bro"}`), fmt.Sprint(rr.Body), "They should be equal")
}

func (suite *HandlerSuite) TestDeleteContainerHandler() {
	payload := []byte(`{"name":"test-container-13"}`)
	req, err := http.NewRequest("DELETE", "/api/v1/container", bytes.NewBuffer(payload))
	if err != nil {
		suite.Fail(err.Error())
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/container", handler.DeleteContainerHandler)
	router.ServeHTTP(rr, req)
	suite.Equal(http.StatusOK, rr.Code, "They should be equal")
}

func (suite *HandlerSuite) TestUpdateContainerState() {
	payload := []byte(`{"name":"test-container-13", "state":{"action":"start", "timeout":60}}`)
	req, err := http.NewRequest("POST", "/api/v1/container/updatestate", bytes.NewBuffer(payload))
	if err != nil {
		suite.Fail(err.Error())
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/container/updatestate", handler.UpdateStateContainerHandler)
	router.ServeHTTP(rr, req)
	suite.Equal(http.StatusOK, rr.Code, "They should be equal")
}
