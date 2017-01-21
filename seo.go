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

	if len(os.Args) < 2 {
		fmt.Println("Usage: seo [location] [selector]")
		os.Exit(0)
	}

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
		bad = append(bad, "The text does not contain any subheadings.")
	}

	if countWords(body.Text()) <= 300 {
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
		bad = append(bad, fmt.Sprintf("%d of the subheadings is followed by more than 300 words.", c))
	} else {
		good = append(good, "The amount of words following each subheading doesn't exceed 300 words")
	}

	fmt.Printf("Analysing \u001B[4m%s\u001B[24m\n", loc)
	for _, l := range good {
		fmt.Printf("\u001B[32m+\u001B[39m %s\n", l)
	}
	for _, l := range bad {
		fmt.Printf("\u001B[31m+\u001B[39m %s\n", l)
	}

}

func countHeadingFollowedByWords(l int) int {
	count := 0
	body.Find("h2, h3, h4, h5").Each(func(i int, s *goquery.Selection) {
		sub := s.NextUntil("h2, h3, h4, h5")
		wordCount := countWords(sub.Text())
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

// check paragraph length
func countParagraphsWithWords(l int) int {
	count := 0

	body.Find("p").Each(func(i int, s *goquery.Selection) {
		wordCount := countWords(s.Text())
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
