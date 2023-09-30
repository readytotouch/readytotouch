package env

import (
	"fmt"
)

const (
	Production  = "production"
	Development = "development"
)

var (
	environment = assertEnvironment()
)

// Environment - const
func Environment() string {
	return environment
}

func assertEnvironment() string {
	environment := Required("ENVIRONMENT")
	switch environment {
	case Production, Development:
		return environment
	default:
		panic(fmt.Sprintf("environment variable %q has unknown value %q", "ENVIRONMENT", environment))
	}
}
