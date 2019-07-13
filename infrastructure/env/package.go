package env

import (
	"os"

	"github.com/krakenlab/ternary"
)

// Defaults
const (
	DefaultServiceName = "noname"
	DefaultInterface   = "0.0.0.0"
	DefaultPort        = "80"
)

// ServiceName $SERVICE_NAME
func ServiceName() string {
	def := os.Getenv("SERVICE_NAME")
	return ternary.String(def == "", DefaultServiceName, def)
}

// Interface $INTERFACE
func Interface() string {
	def := os.Getenv("INTERFACE")
	return ternary.String(def == "", DefaultInterface, def)
}

// Port $PORT
func Port() string {
	def := os.Getenv("PORT")
	return ternary.String(def == "", DefaultPort, def)
}
