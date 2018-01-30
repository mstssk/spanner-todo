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

func TestReverse(t *testing.T) {
	if Reverse("test") != "tset" {
		t.Errorf("ng")
	}
	if Reverse("あいうえお") != "おえういあ" {
		t.Errorf("ng")
	}
	if Reverse("𠮷野家") != "家野𠮷" {
		t.Errorf("ng")
	}
	if Reverse("🍎🍏") != "🍏🍎" {
		t.Errorf("ng")
	}
}
