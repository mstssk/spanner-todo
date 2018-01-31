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

type TestStruct struct {
	Foo string
	Bar int
}

func TestGetStructFieldNames(t *testing.T) {
	name, fields := GetStructFieldNames(TestStruct{})
	if name != "TestStruct" {
		t.Errorf("ng %v", name)
	}
	if len(fields) != 2 {
		t.Errorf("ng %v", len(fields))
		for i, v := range fields {
			t.Errorf("%v %v", i, v)
		}
	} else {
		if fields[0] != "Foo" {
			t.Errorf("ng %v", fields[0])
		}
		if fields[1] != "Bar" {
			t.Errorf("ng %v", fields[1])
		}
	}
}
