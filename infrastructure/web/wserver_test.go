package web

import (
	"net/http"
	"os"
	"testing"

	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/env"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/web/routes"

	"github.com/krakenlab/gspec"
	"github.com/krakenlab/gspec/websuite"
)

type WServerTest struct {
	websuite.WebSuite
	*WServer
}

func (suite *WServerTest) SetupTest() {
	defer suite.Suite.SetupTest()

	suite.WServer = NewWServer()
	suite.WebSuite.Server = suite.WServer.engine
}

func (suite *WServerTest) TestStart() {
	os.Setenv("PORT", "-1")
	defer os.Setenv("PORT", "")

	suite.Panics(func() {
		suite.WServer.Start()
	})
}

func (suite *WServerTest) TestGetRootHandler() {
	response := suite.WebSuite.GET(routes.RootRoute, "")
	body := response.Body.String()

	suite.Equal(http.StatusOK, response.Code)

	suite.Contains(body, "service")
	suite.Contains(body, env.DefaultServiceName)
}

func (suite *WServerTest) TestPutAccountHandler() {
	response := suite.WebSuite.PUT(routes.AccountRoute, `{"username": "aa", "password": "aaa"}`)
	body := response.Body.String()
	suite.Equal(http.StatusInternalServerError, response.Code)
	suite.Contains(body, "Invalid Request")

	response = suite.WebSuite.PUT(routes.AccountRoute, `{"username": "this_a_test", "password": "this_a_super_test"}`)
	body = response.Body.String()
	suite.Equal(http.StatusOK, response.Code)
	suite.Contains(body, "this_a_test")
	suite.Contains(body, "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
}

func TestWServerTest(t *testing.T) {
	gspec.Run(t, new(WServerTest))
}
