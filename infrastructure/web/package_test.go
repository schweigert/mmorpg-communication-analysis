package web

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PackageSuite struct {
	gspec.Suite
}

func (suite *PackageSuite) TestNewWServer() {
	server := NewWServer()

	suite.NotNil(server)
	suite.NotNil(server.engine)
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
