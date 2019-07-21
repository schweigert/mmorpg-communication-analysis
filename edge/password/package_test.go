package password

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PackageSuite struct {
	gspec.Suite
}

func (suite *PackageSuite) TestPasswordSuite() {
	pass := "securePass"
	suite.Equal("2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", Sum(pass))
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
