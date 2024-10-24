package app

import (
	"os"
	"strings"
)

func ConfGet(key string) (string, bool) {
	v, ok := os.LookupEnv(strings.ToUpper(key))
	return v, ok
}
