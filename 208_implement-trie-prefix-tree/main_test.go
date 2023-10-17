package implementtrieprefixtree

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")

	if !trie.Search("apple") {
		t.Errorf("Expected %v but was %v", true, false)
	}

	if trie.Search("app") {
		t.Errorf("Expected %v but was %v", false, true)
	}

	if !trie.StartsWith("app") {
		t.Errorf("Expected %v but was %v", true, false)
	}

	trie.Insert("app")

	if !trie.Search("app") {
		t.Errorf("Expected %v but was %v", true, false)
	}
}
