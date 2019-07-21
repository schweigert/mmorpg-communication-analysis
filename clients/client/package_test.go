package client

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PackageSuite struct {
	gspec.Suite
}

func (suite *PackageSuite) TestRandomUsername() {
	suite.Contains(RandomUsername(), "_")

	usernames := []string{RandomUsername()}

	for count := 0; count < 1000; count++ {
		username := RandomUsername()
		suite.NotContains(usernames, username)
		usernames = append(usernames, username)
	}
}

func (suite *PackageSuite) TestRandomPassword() {
	suite.GreaterOrEqual(len(RandomPassword()), 8)
	suite.LessOrEqual(len(RandomPassword()), 16)
}

func TestPackageSuite(t *testing.T) {
	gspec.Run(t, new(PackageSuite))
}
