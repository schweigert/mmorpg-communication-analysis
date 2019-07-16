package env

import (
	"os"

	"github.com/krakenlab/ternary"
)

// Defaults
const (
	DefaultServiceName  = "noname"
	DefaultInterface    = "0.0.0.0"
	DefaultPort         = "80"
	DefaultPostgresHost = "localhost"
	DefaultPostgresPort = "5432"
	DefaultPostgresUser = "postgres"
	DefaultPostgresPass = "postgres"
	DefaultPostgresSsl  = "disable"
	DefaultPostgresDb   = "test"
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

// PostgresHost $POSTGRES_HOST
func PostgresHost() string {
	def := os.Getenv("POSTGRES_HOST")
	return ternary.String(def == "", DefaultPostgresHost, def)
}

// PostgresPort $POSTGRES_PORT
func PostgresPort() string {
	def := os.Getenv("POSTGRES_PORT")
	return ternary.String(def == "", DefaultPostgresPort, def)
}

// PostgresUser $POSTGRES_USER
func PostgresUser() string {
	def := os.Getenv("POSTGRES_USER")
	return ternary.String(def == "", DefaultPostgresUser, def)
}

// PostgresPass $POSTGRES_PASS
func PostgresPass() string {
	def := os.Getenv("POSTGRES_PASS")
	return ternary.String(def == "", DefaultPostgresPass, def)
}

// PostgresSsl $POSTGRES_SSL
func PostgresSsl() string {
	def := os.Getenv("POSTGRES_SSL")
	return ternary.String(def == "", DefaultPostgresSsl, def)
}

// PostgresDb $POSTGRES_DB
func PostgresDb() string {
	def := os.Getenv("POSTGRES_DB")
	return ternary.String(def == "", DefaultPostgresDb, def)
}
