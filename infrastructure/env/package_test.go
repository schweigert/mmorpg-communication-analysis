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

func (suite *PackageSuite) TestDefaultInterface() {
	suite.Equal("0.0.0.0", DefaultInterface)
}

func (suite *PackageSuite) TestDefaultPort() {
	suite.Equal("80", DefaultPort)
}

func (suite *PackageSuite) TestServiceName() {
	suite.Equal(DefaultServiceName, ServiceName())
}

func (suite *PackageSuite) TestInterface() {
	suite.Equal(DefaultInterface, Interface())
}

func (suite *PackageSuite) TestPort() {
	suite.Equal(DefaultPort, Port())
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
