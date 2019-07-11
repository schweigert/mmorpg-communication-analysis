package nodes

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PackageSuite struct {
	gspec.Suite
}

func (suite *PackageSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *PackageSuite) TestNewBaseNode() {
	suite.NotNil(NewBaseNode())
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
