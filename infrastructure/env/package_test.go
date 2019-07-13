package env

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PackageSuite struct {
	gspec.Suite
}

func (suite *PackageSuite) TestDefaultServiceName() {
	suite.Equal("noname", DefaultServiceName)
}

func (suite *PackageSuite) TestServiceName() {
	suite.Equal(DefaultServiceName, ServiceName())
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
