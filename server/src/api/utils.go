package api

import (
	"reflect"
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

// Reverse string
func Reverse(s string) string {
	// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// GetStructFieldNames with struct name
func GetStructFieldNames(in interface{}) (string, []string) {
	t := reflect.ValueOf(in).Type()
	name := t.Name()
	fields := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
	}
	return name, fields
}
