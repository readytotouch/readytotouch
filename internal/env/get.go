package env

import (
	"fmt"
	"os"
)

// Required - assert environment variable defined
func Required(name string) string {
	var result = os.Getenv(name)

	if result == "" {
		panic(fmt.Sprintf("environment variable %s required", name))
	}

	return result
}
