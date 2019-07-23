package routes

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

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
