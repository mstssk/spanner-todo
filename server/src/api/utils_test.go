package api

import (
	"testing"
)

func TestHasPrefixSome(t *testing.T) {
	if !HasPrefixSome([]string{"test"}, "test") {
		t.Errorf("ng")
	}
	if !HasPrefixSome([]string{"foo", "bar", "baz"}, "bar") {
		t.Errorf("ng")
	}
	if HasPrefixSome([]string{"foo", "bar", "baz"}, "hoge") {
		t.Errorf("ng")
	}
}
