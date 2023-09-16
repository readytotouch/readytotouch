package env

import (
	"fmt"
	"os"
)

func Must(name string) string {
	var result = os.Getenv(name)

	if result == "" {
		panic(fmt.Sprintf("expected environment value %s", name))
	}

	return result
}
