package environment

import (
	"fmt"
	"os"
)

// Retrieves the value of the environmental variable named by the key, or
// panics if it isn't set.
//
// For a less drastic approach, use [Lookup].
func ValueFor(key string) string {
	if p, isSet := Lookup(key); !isSet {
		panic(fmt.Sprintf("%s isn't set!", key))
	} else {
		return p
	}
}

// Retrieves the value of the environmental variable named by the key and a
// boolean representing if it was set. (Set to empty also counts!)
func Lookup(key string) (string, bool) {
	return os.LookupEnv(key)
}
