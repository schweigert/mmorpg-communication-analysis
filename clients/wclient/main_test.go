package main

import (
	"testing"

	"github.com/schweigert/mmorpg-communication-analysis/clients/client"

	"github.com/krakenlab/gspec"
)

type ApplicationSuite struct {
	gspec.Suite
}

func (suite *ApplicationSuite) TestInit() {
	suite.NotEmpty(client.Username)
	suite.NotEmpty(client.Password)
	suite.NotEmpty(client.SecurePassword)
}

func TestApplicationSuite(t *testing.T) {
	gspec.Run(t, new(ApplicationSuite))
}
