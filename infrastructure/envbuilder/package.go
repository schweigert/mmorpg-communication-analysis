package envbuilder

import (
	"fmt"

	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/env"
)

// GinAddr format Interface and Port in %s:%s format
func GinAddr() string {
	return fmt.Sprintf("%s:%s", env.Interface(), env.Port())
}
