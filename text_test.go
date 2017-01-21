package main

import (
	"testing"
)

func TestCountWords(t *testing.T) {
	s := "Hello these are some words. Count them for me, will you?"
	count := countWords(s)
	if count != 11 {
		t.Errorf("Expected %d, got %d", 11, count)
	}
}

func TestCountSentences(t *testing.T) {
	s := "Hello these are some words. Count them for me, will you? Let's continue; with more text. And then end it!"
	count := countSentences(s)
	if count != 5 {
		t.Errorf("Expected %d, got %d", 5, count)
	}
}

func TestCountSyllables(t *testing.T) {
	s := "Ambulance"
	count := countSyllables(s)
	if count != 3 {
		t.Errorf("Expected %d, got %d", 3, count)
	}

	s = "Ape"
	count = countSyllables(s)
	if count != 1 {
		t.Errorf("Expected %d, got %d", 1, count)
	}

	s = "Booze"
	count = countSyllables(s)
	if count != 1 {
		t.Errorf("Expected %d, got %d", 1, count)
	}
}
