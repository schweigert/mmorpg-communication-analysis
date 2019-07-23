package forms

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type JwtCredentialFormSuite struct {
	gspec.Suite

	Form     *JwtCredentialForm
	Password []byte
}

func (suite *JwtCredentialFormSuite) SetupTest() {
	defer suite.Suite.SetupTest()

	suite.Form = &JwtCredentialForm{}
	suite.Form.Username = "a_single_username"
	suite.Password = []byte{'p', 'a', 's', 's', 'w', 'o', 'r', 'd'}
}

func (suite *JwtCredentialFormSuite) TestToken() {
	suite.Contains(suite.Form.Token(suite.Password), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
}

func (suite *JwtCredentialFormSuite) TestParse() {
	token := suite.Form.Token(suite.Password)

	otherForm := &JwtCredentialForm{Username: ""}

	suite.True(otherForm.Parse(suite.Password, token))
	suite.Equal("a_single_username", otherForm.Username)
}

func TestJwtCredentialFormSuite(t *testing.T) {
	gspec.Run(t, new(JwtCredentialFormSuite))
}
