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

func (suite *PackageSuite) TestDefaultPostgresHost() {
	suite.Equal(DefaultPostgresHost, "localhost")
}

func (suite *PackageSuite) TestDefaultPostgresPort() {
	suite.Equal(DefaultPostgresPort, "5432")
}

func (suite *PackageSuite) TestDefaultPostgresUser() {
	suite.Equal(DefaultPostgresUser, "postgres")
}

func (suite *PackageSuite) TestDefaultPostgresPass() {
	suite.Equal(DefaultPostgresPass, "postgres")
}

func (suite *PackageSuite) TestDefaultPostgresSsl() {
	suite.Equal(DefaultPostgresSsl, "disable")
}

func (suite *PackageSuite) TestDefaultPostgresDb() {
	suite.Equal(DefaultPostgresDb, "test")
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

func (suite *PackageSuite) TestPostgresHost() {
	suite.Equal(DefaultPostgresHost, PostgresHost())
}

func (suite *PackageSuite) TestPostgresPort() {
	suite.Equal(DefaultPostgresPort, PostgresPort())
}

func (suite *PackageSuite) TestPostgresUser() {
	suite.Equal(DefaultPostgresUser, PostgresUser())
}

func (suite *PackageSuite) TestPostgresPass() {
	suite.Equal(DefaultPostgresPass, PostgresPass())
}

func (suite *PackageSuite) TestPostgresSsl() {
	suite.Equal(DefaultPostgresSsl, PostgresSsl())
}

func (suite *PackageSuite) TestPostgresDb() {
	suite.Equal(DefaultPostgresDb, PostgresDb())
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
