package utils

import "strings"

func EndpointSyntaxCheck(endpoint string) bool {
	valid := true
	if strings.Split(endpoint, "")[0] != "/" {
		valid = false
	}
	return valid
}
