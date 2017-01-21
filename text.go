package main

import (
	"strings"
)

// instance, periods, explanation points, colons and semicolons serve as sentence delimiter
func countSentences(t string) int {
	return 0
}

// each vowel in a word is considered one syllable subject to:
// (a) -es, -ed and -e (except -le) endings are ignored;
// (b) words of three letters or shorter count as single syllables;
// (c) consecutive vowels count as one syllable.
func countSyllables(t string) int {
	return 0
}

// each group of continuous non-blank characters with beginning and ending punctuation removed counts as a word
func getWords(t string) []string {
	return strings.Split(t, " ")
}

func countWords(t string) int {
	return len(getWords(t))
}
