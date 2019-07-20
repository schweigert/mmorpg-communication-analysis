package forms

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type PasswordFormSuite struct {
	gspec.Suite

	Form *PasswordForm
}

func (suite *PasswordFormSuite) SetupTest() {
	defer suite.Suite.SetupTest()

	suite.Form = &PasswordForm{}
	suite.Form.Password = "a_single_password"
}

func (suite *PasswordFormSuite) TestValid() {
	suite.True(suite.Form.Valid())
}

func TestPasswordFormSuite(t *testing.T) {
	gspec.Run(t, new(PasswordFormSuite))
}
