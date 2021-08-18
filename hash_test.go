package main

import (
	"testing"
)

// Init
func TestHashInit(t *testing.T) {
	htable := Init()
	if htable == nil || htable.table[0] == nil {
		t.Errorf("There was a problem creating table")
	}
}

// Insert
func TestHashInsert(t *testing.T) {
	htable := Init()

	list := []string{
		"ERIC",
		"MARTIN",
		"BORJA",
		"ADRIAN",
		"DANIEL",
		"LUIS",
		"URKO",
		"CRISTIAN",
	}

	for _, s := range list {
		htable.Insert(s)
	}
}

// Search
func TestHastSearch(t *testing.T) {
	htable := Init()

	list := []string{
		"ERIC",
		"MARTIN",
		"BORJA",
		"ADRIAN",
		"DANIEL",
		"LUIS",
		"URKO",
		"CRISTIAN",
	}

	for _, s := range list {
		htable.Insert(s)
	}

	if !htable.Search("MARTIN") {
		t.Errorf("Searching not working")
	}
}

// Delete
func TestHastDelete(t *testing.T) {
	htable := Init()

	keyToDelete := "MARTIN"

	htable.Insert(keyToDelete)

	if !htable.Search(keyToDelete) {
		t.Errorf("Searching not working")
	}

	htable.Delete(keyToDelete)

	if htable.Search(keyToDelete) {
		t.Errorf("Deleting not working")
	}
}
