package forms

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type UsernameFormSuite struct {
	gspec.Suite

	Form *UsernameForm
}

func (suite *UsernameFormSuite) SetupTest() {
	defer suite.Suite.SetupTest()

	suite.Form = &UsernameForm{}
	suite.Form.Username = "a_single_username"
}

func (suite *UsernameFormSuite) TestValid() {
	suite.True(suite.Form.Valid())
}

func TestUsernameFormSuite(t *testing.T) {
	gspec.Run(t, new(UsernameFormSuite))
}
