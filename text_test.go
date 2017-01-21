package main

import (
	"strings"
	"testing"
)

func TestGetWords(t *testing.T) {
	s := "Hello world"
	words := getWords(s)
	if len(words) != 2 || words[0] != "Hello" || words[1] != "world" {
		t.Errorf("Invalid result %s", words)
	}
}

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

func TestIsVowel(t *testing.T) {
	vowels := []string{"a", "o", "i", "u", "e"}
	for _, c := range vowels {
		if !isVowel(c) {
			t.Errorf("%s should return true", c)
		}
	}

	nonVowels := strings.Split("zxcvbnmqwrtypsdfghjkl", "")
	for _, c := range nonVowels {
		if isVowel(c) {
			t.Errorf("%s should return false", c)
		}
	}
}
