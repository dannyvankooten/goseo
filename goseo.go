package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

	c = countHeadings(body)
	if c == 0 {
		bad = append(bad, "The text does not contain any subheadings.")
	}

	if countWords(body.Text()) <= 300 {
		bad = append(bad, "You have far too little content, please add some content to enable a good analysis.")
	} else {
		good = append(good, "The text contains more than 300 words.")
	}

	c = countParagraphsWithWords(body, 150)
	if c >= 1 {
		bad = append(bad, fmt.Sprintf("%d of the paragraphs contains more than the recommended maximum of 150 words.", c))
	} else {
		good = append(good, "None of the paragraphs contain more than 150 words.")
	}

	c = countHeadingsWithWords(body, 300)
	if c > 1 {
		bad = append(bad, fmt.Sprintf("%d of the subheadings is followed by more than 300 words.", c))
	} else {
		good = append(good, "The amount of words following each subheading doesn't exceed 300 words")
	}

	// long sentences
	sentences := countSentences(body.Text())
	longSentences := countSentencesWithWords(body, 20)
	percentage := float32(longSentences) / float32(sentences) * 100
	if percentage > 25 {
		bad = append(bad, fmt.Sprintf("%.1f%% of of the sentences contain more than 20 words.", percentage))
	} else {
		good = append(good, fmt.Sprintf("%.1f%% of of the sentences contain more than 20 words.", percentage))
	}

	// kincaid
	kc := calculateKincaid(body.Text())
	if kc < 60 {
		bad = append(bad, fmt.Sprintf("The copy scores %1.f in the Flesch Reading Ease test.", kc))
	} else {
		good = append(good, fmt.Sprintf("The copy scores %1.f in the Flesch Reading Ease test.", kc))
	}

	fmt.Printf("Analysing \u001B[4m%s\u001B[24m\n", loc)
	for _, l := range good {
		fmt.Printf("\u001B[32m+\u001B[39m %s\n", l)
	}
	for _, l := range bad {
		fmt.Printf("\u001B[31m+\u001B[39m %s\n", l)
	}

}
