package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"os"
	"strings"
)

var body *goquery.Selection

func main() {
	var doc *goquery.Document

	loc := os.Args[1]
	if strings.HasPrefix(loc, "http") {
		doc, _ = goquery.NewDocument(loc)
	} else {
		b, _ := ioutil.ReadFile(loc)
		doc, _ = goquery.NewDocumentFromReader(bytes.NewReader(b))
	}

	// find article body
	body = doc.Find("body")
	if len(os.Args) > 2 {
		body = doc.Find(os.Args[2])
	}

	var good []string
	var bad []string
	var c int

	c = countHeadings()
	if c == 0 {
		bad = append(bad, "The text does not contain any subheadings. Add at least one subheading.")
	} else {
		good = append(good, fmt.Sprintf("The text contains %d subheadings.", c))
	}

	if countWords(body) <= 300 {
		bad = append(bad, "You have far too little content, please add some content to enable a good analysis.")
	} else {
		good = append(good, "The text contains more than 300 words.")
	}

	c = countParagraphsWithWords(150)
	if c >= 1 {
		bad = append(bad, fmt.Sprintf("%d of the paragraphs contains more than the recommended maximum of 150 words.", c))
	} else {
		good = append(good, "None of the paragraphs contain too many words.")
	}

	c = countHeadingFollowedByWords(300)
	if c > 1 {
		bad = append(bad, fmt.Sprintf("%d of the subheadings is followed by more than the recommended maximum of 300 words.", c))
	} else {
		good = append(good, "No subheadings are followed by more than 300 words.")
	}

	fmt.Printf("Analysing %s\n", loc)
	fmt.Printf("\n# The good\n")
	for _, l := range good {
		fmt.Printf("- %s\n", l)
	}
	if len(good) == 0 {
		fmt.Print("...\n")
	}

	fmt.Printf("\n# The bad\n")
	for _, l := range bad {
		fmt.Printf("- %s\n", l)
	}
	if len(bad) == 0 {
		fmt.Print("...\n")
	}
}

func countHeadingFollowedByWords(l int) int {
	count := 0
	body.Find("h2, h3, h4, h5").Each(func(i int, s *goquery.Selection) {
		sub := s.NextUntil("h2, h3, h4, h5")
		wordCount := countWords(sub)
		if wordCount > l {
			count++
		}
	})

	return count
}

// count subheadings, > 1
func countHeadings() int {
	return body.Find("h2, h3, h4, h5").Length()
}

// count words, > 300
func countWords(s *goquery.Selection) int {
	return len(strings.Split(s.Text(), " "))
}

// check paragraph length
func countParagraphsWithWords(l int) int {
	count := 0

	body.Find("p").Each(func(i int, s *goquery.Selection) {
		wordCount := countWords(s)
		if wordCount > l {
			count++
		}
	})

	return count
}

//  maximum 25% of sentences with over 20 words

// passive voice

// transition words

// flesch reading ease test
