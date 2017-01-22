package main

import (
	"bufio"
	"os"
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
	var words []string
	for _, w := range strings.Split(t, " ") {
		w = strings.Trim(w, ",.;!'?:()& “”\n")
		if len(w) > 1 {
			words = append(words, w)
		}

	}
	return words
}

func countWords(t string) int {
	return len(getWords(t))
}

func calculateKincaid(t string) float32 {
	// RE = 206.835 – (1.015 x ASL) – (84.6 x ASW)
	sentenceCount := countSentences(t)
	words := getWords(t)
	wordCount := len(words)
	syllableCount := 0
	for _, w := range words {
		syllableCount += countSyllables(w)
	}

	asl := float32(wordCount) / float32(sentenceCount)
	asw := float32(syllableCount) / float32(wordCount)

	return 206.835 - (1.015 * asl) - (84.6 * asw)
}

func getTransitionWordMap() map[string]bool {
	f, _ := os.Open("data/transition-words.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	transitionWords := make(map[string]bool)
	for scanner.Scan() {
		transitionWords[scanner.Text()] = true
	}
	return transitionWords
}

func countSentencesWithTransitionWord(t string) int {
	t = strings.ToLower(t)
	s := getSentences(t)
	transitionWords := getTransitionWordMap()
	count := 0

	for _, s := range s {
		words := getWords(s)
		for _, w := range words {
			if transitionWords[w] {
				count++
				break
			}
		}
	}

	return count
}
