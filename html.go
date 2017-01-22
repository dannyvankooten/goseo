package main

import (
	"github.com/PuerkitoBio/goquery"
)

// count subheadings, > 1
func countHeadings(s *goquery.Selection) int {
	return s.Find("h2, h3, h4, h5").Length()
}

func countHeadingsWithWords(s *goquery.Selection, l int) int {
	count := 0
	s.Find("h2, h3, h4, h5").Each(func(i int, s *goquery.Selection) {
		sub := s.NextUntil("h2, h3, h4, h5")
		wordCount := countWords(sub.Text())
		if wordCount > l {
			count++
		}
	})

	return count
}

// max 25% of lsentences should contain 20+ words
func countSentencesWithWords(s *goquery.Selection, l int) int {
	sentences := getSentences(s.Text())
	count := 0

	for _, s := range sentences {
		if countWords(s) > l {
			count++
		}
	}

	return count
}

// check paragraph length
func countParagraphsWithWords(s *goquery.Selection, l int) int {
	count := 0

	s.Find("p").Each(func(i int, s *goquery.Selection) {
		wordCount := countWords(s.Text())
		if wordCount > l {
			count++
		}
	})

	return count
}
