package models

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PackageTest struct {
	gspec.Suite
}

func (suite *PackageTest) TestNewAccount() {
	suite.NotNil(NewAccount("username", "securePass"))
}

func TestPackageTest(t *testing.T) {
	gspec.Run(t, new(PackageTest))
}
