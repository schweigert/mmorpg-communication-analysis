package web

import (
	"net/http"
	"os"
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

func (suite *ServerTest) TestStart() {
	os.Setenv("PORT", "-1")
	defer os.Setenv("PORT", "")

	suite.Panics(func() {
		suite.Server.Start()
	})
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
