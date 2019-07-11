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

func (suite *PackageSuite) TestNewHierarchicalNode() {
	node := NewHierarchicalNode()

	suite.NotNil(node)
	suite.NotPanics(func() { node.Update(0.0) })
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
