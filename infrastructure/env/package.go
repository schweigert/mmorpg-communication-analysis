package env

import (
	"os"

	"github.com/krakenlab/ternary"
)

// Defaults
const (
	DefaultServiceName = "noname"
)

// ServiceName $SERVICE_NAME
func ServiceName() string {
	def := os.Getenv("SERVICE_NAME")
	return ternary.String(def == "", DefaultServiceName, def)
}
