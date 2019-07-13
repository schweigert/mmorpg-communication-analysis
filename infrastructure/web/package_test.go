package web

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PackageSuite struct {
	gspec.Suite
}

func (suite *PackageSuite) TestRootRoute() {
	suite.Contains(Routes, RootRoute)
}

func (suite *PackageSuite) TestRoutes() {
	// Routes must be uniq
	for firstIndex, firstRoute := range Routes {
		for secondIndex, secondRoute := range Routes {
			if firstIndex == secondIndex {
				continue
			}

			suite.NotEqual(firstRoute, secondRoute)
		}
	}
}

func (suite *PackageSuite) TestNewServer() {
	server := NewServer()

	suite.NotNil(server)
	suite.NotNil(server.engine)
}

func (suite *PackageSuite) TestNewWillsonServer() {
	server := NewWillsonServer()

	suite.NotNil(server)
	suite.NotNil(server.Server)
	suite.NotNil(server.engine)
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
