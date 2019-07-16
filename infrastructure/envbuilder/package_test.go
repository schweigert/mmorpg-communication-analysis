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

func (suite *PackageSuite) TestPostgresAddr() {
	suite.Equal(
		"host=localhost port=5432 user=postgres password=postgres sslmode=disable dbname=test",
		PostgresAddr(),
	)
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
