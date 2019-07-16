package repositories

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PackageSuite struct {
	gspec.Suite
}

func (suite *PackageSuite) TestDb() {
	suite.NotNil(db)
}

func (suite *PackageSuite) TestNewAccountRepository() {
	suite.NotNil(NewAccountRepository())
	suite.NotNil(NewAccountRepository().db)
	suite.Equal(db, NewAccountRepository().db)
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
