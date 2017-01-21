package main

import (
	"strings"
)

// periods, explanation points, colons and semicolons serve as sentence delimiter
func getSentences(t string) []string {
	return strings.FieldsFunc(t, func(r rune) bool {
		switch r {
		case '.', '!', ':', ';', '?':
			return true
		}
		return false
	})
}

func countSentences(t string) int {
	return len(getSentences(t))
}

func isVowel(c string) bool {
	return c == "a" || c == "e" || c == "o" || c == "u" || c == "i"
}

// each vowel in a word is considered one syllable subject to:
// (a) -es, -ed and -e (except -le) endings are ignored;
// (b) words of three letters or shorter count as single syllables;
// (c) consecutive vowels count as one syllable.
func countSyllables(t string) int {
	if len(t) <= 3 {
		return 1
	}

	// strip ignored endings
	t = strings.ToLower(t)
	t = strings.TrimSuffix(t, "es")
	t = strings.TrimSuffix(t, "ed")
	if !strings.HasSuffix(t, "le") {
		t = strings.TrimSuffix(t, "e")
	}

	// count non-consecutive vowels
	count := 0
	chars := strings.Split(t, "")
	for i, c := range chars {
		if isVowel(c) && (i == 0 || !isVowel(chars[i-1])) {
			count++
		}
	}

	return count
}

// each group of continuous non-blank characters with beginning and ending punctuation removed counts as a word
func getWords(t string) []string {
	return strings.Split(t, " ")
}

func countWords(t string) int {
	return len(getWords(t))
}
