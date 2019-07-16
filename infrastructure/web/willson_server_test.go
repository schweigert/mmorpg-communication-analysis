package web

import (
	"net/http"
	"testing"

	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/repositories"

	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/env"

	"github.com/krakenlab/gspec"
	"github.com/krakenlab/gspec/websuite"
)

type WillsonServerTest struct {
	websuite.WebSuite
	*WillsonServer
}

func (suite *WillsonServerTest) SetupTest() {
	defer suite.Suite.SetupTest()

	repositories.ClearDB()

	suite.WillsonServer = NewWillsonServer()
	suite.WebSuite.Server = suite.WillsonServer.engine
}

func (suite *WillsonServerTest) TestGet() {
	response := suite.WebSuite.GET(RootRoute, "")
	body := response.Body.String()

	suite.Equal(http.StatusOK, response.Code)

	suite.Contains(body, "service")
	suite.Contains(body, env.DefaultServiceName)
}

func (suite *WillsonServerTest) TestPutAccountHandler() {
	response := suite.WebSuite.PUT(AccountRoute, "??")
	body := response.Body.String()

	suite.Equal(http.StatusInternalServerError, response.Code)
	suite.Contains(body, "error")
	suite.Contains(body, "12341234")
}

func TestWillsonServerTest(t *testing.T) {
	gspec.Run(t, new(WillsonServerTest))
}
