package utils

import (
	"regexp"
)

// ValidContainerID ensures that the given container ID has the
// container runtime prefix set
func ValidContainerID(value string) bool {
	matched, err := regexp.MatchString("^(containerd|docker)://.+$", value)
	if err != nil {
		panic(err)
	}

	return matched
}
