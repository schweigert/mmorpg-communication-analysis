package forms

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type CredentialFormSuite struct {
	gspec.Suite

	Form *CredentialForm
}

func (suite *CredentialFormSuite) SetupTest() {
	defer suite.Suite.SetupTest()

	suite.Form = &CredentialForm{}
	suite.Form.Username = "a_single_username"
	suite.Form.Password = "a_single_password"
}

func (suite *CredentialFormSuite) TestValid() {
	suite.True(suite.Form.Valid())
}

func TestCredentialFormSuite(t *testing.T) {
	gspec.Run(t, new(CredentialFormSuite))
}
