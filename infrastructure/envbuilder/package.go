package envbuilder

import (
	"fmt"

	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/env"
)

// GinAddr format Interface and Port in %s:%s format from env
func GinAddr() string {
	return fmt.Sprintf(
		"%s:%s",
		env.Interface(),
		env.Port(),
	)
}

// PostgresAddr format host, port, user, password, sslmode and dbname from env
func PostgresAddr() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=%s dbname=%s",
		env.PostgresHost(),
		env.PostgresPort(),
		env.PostgresUser(),
		env.PostgresPass(),
		env.PostgresSsl(),
		env.PostgresDb(),
	)
}
