package web

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type ConstSuite struct {
	gspec.Suite
}

func (suite *ConstSuite) TestRootRoute() {
	suite.Contains(Routes, RootRoute)
}

func (suite *ConstSuite) TestRoutes() {
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

func TestConstSuite(t *testing.T) {
	gspec.Run(t, new(ConstSuite))
}
