package client

import (
	"fmt"
	"strings"

	"github.com/schweigert/mmorpg-communication-analysis/edge/password"

	"github.com/dmgk/faker"
)

// Client variables
var (
	Username       string
	Password       string
	SecurePassword string
)

// FillPersonalData into variables
func FillPersonalData() {
	Username = RandomUsername()
	Password = RandomPassword()
	SecurePassword = password.Sum(Password)
}

// RandomUsername to create an account
func RandomUsername() string {
	return strings.ToLower(
		fmt.Sprintf(
			"%s_%s_%s_%d",
			faker.Name().FirstName(),
			faker.Name().LastName(),
			faker.Name().LastName(),
			faker.RandomInt(1980, 2019),
		),
	)
}

// RandomPassword using a secure password
func RandomPassword() string {
	return faker.Internet().Password(8, 16)
}
