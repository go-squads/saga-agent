package lxdclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
}

func (suite *HandlerSuite) TearDownSuite() {

}

func (suite *HandlerSuite) TestGetContainerHandler() {
	req, err := http.NewRequest("GET", "/api/v1/containers", nil)
	if err != nil {
		suite.Fail(err.Error())
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handler.GetContainersHadler)
	handler.ServeHTTP(rr, req)

	suite.Equal(http.StatusOK, rr.Code, "They should be equal")
}
