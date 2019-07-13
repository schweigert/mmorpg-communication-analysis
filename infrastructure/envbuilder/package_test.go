package envbuilder

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PackageSuite struct {
	gspec.Suite
}

func (suite *PackageSuite) TestGinAddr() {
	suite.Equal("0.0.0.0:80", GinAddr())
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
