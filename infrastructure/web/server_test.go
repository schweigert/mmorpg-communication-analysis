package web

import (
	"net/http"
	"testing"

	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/env"

	"github.com/krakenlab/gspec"
	"github.com/krakenlab/gspec/websuite"
)

type ServerTest struct {
	websuite.WebSuite
	*Server
}

func (suite *ServerTest) SetupTest() {
	defer suite.Suite.SetupTest()

	suite.Server = NewServer()
	suite.WebSuite.Server = suite.Server.engine
}

func (suite *ServerTest) TestGet() {
	response := suite.WebSuite.GET(RootRoute, "")
	body := response.Body.String()

	suite.Equal(http.StatusOK, response.Code)

	suite.Contains(body, "service")
	suite.Contains(body, env.DefaultServiceName)
}

func TestServerTest(t *testing.T) {
	gspec.Run(t, new(ServerTest))
}
