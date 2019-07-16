package main

import (
	"os"
	"testing"

	"github.com/krakenlab/gspec"
)

type ApplicationSuite struct {
	gspec.Suite
}

func (suite *ApplicationSuite) TestMain() {
	os.Setenv("PORT", "-1")
	defer os.Setenv("PORT", "")

	suite.Panics(func() {
		main()
	})
}

func TestApplicationSuite(t *testing.T) {
	gspec.Run(t, new(ApplicationSuite))
}
