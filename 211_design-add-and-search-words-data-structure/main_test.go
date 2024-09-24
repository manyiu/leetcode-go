package designaddandsearchwordsdatastructure

import "testing"

func TestWordDictionary(t *testing.T) {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")

	if obj.Search("pad") {
		t.Fatalf("expecting false")
	}

	if !obj.Search("bad") {
		t.Fatalf("expecting true")
	}

	if !obj.Search(".ad") {
		t.Fatalf("expecting true")
	}

	if !obj.Search("b..") {
		t.Fatalf("expecting true")
	}
}
