package api

import (
	"strings"
)

// HasPrefixSome strがprefixesいずれかから始まっていればtrueを返す
func HasPrefixSome(prefixes []string, str string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}
	return false
}
