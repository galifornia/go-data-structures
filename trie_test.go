package main

import (
	"testing"
)

func TestInit(t *testing.T) {
	trie := InitTrie()
	if trie == nil || trie.root == nil {
		t.Errorf("Creation didn't happen")
	}

}

func TestInsert(t *testing.T) {
	trie := InitTrie()
	trie.Insert("alice")
	trie.Insert("shoe")
	trie.Insert("sun")

	i := 's' - 'a' // sun & shoe has been inserted, so 's' should not be nil
	if trie.root.children[i] == nil {
		t.Errorf("There should be a TrieNode in that position of the children array")
	}
}

func TestSearch(t *testing.T) {
	trie := InitTrie()
	trie.Insert("pinga")
	trie.Insert("zapato")
	trie.Insert("abrigo")

	if trie.Search("ping") {
		t.Errorf("Word has not been inserted")
	}

	if !trie.Search("pinga") {
		t.Errorf("Should return 'true' since word has been inserted")
	}
}
