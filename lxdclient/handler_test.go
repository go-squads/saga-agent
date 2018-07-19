package lxdclient

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/lxc/lxd/shared/api"
	"github.com/stretchr/testify/suite"
)

type HandlerSuite struct {
	suite.Suite
}

var handler Handler

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (suite *HandlerSuite) SetupSuite() {
	handler = Handler{}
	name := "test-container-11"
	req := api.ContainersPost{
		Name: name,
		Source: api.ContainerSource{
			Type: "none",
		},
	}
	createContainer(req)
	req.Name = "test-container-13"
	createContainer(req)
}

func (suite *HandlerSuite) TearDownSuite() {
	deleteContainer("test-container-11")
	deleteContainer("test-container-12")
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
	suite.Equal(http.StatusOK, rr.Code, "They should be equal")
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
