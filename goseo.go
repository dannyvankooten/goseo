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
var results []result

type result struct {
	text string
	good bool
}

func main() {
	var doc *goquery.Document
	var err error

	if len(os.Args) < 2 {
		fmt.Println("Usage: goseo [location] [selector]")
		os.Exit(0)
	}

	loc := os.Args[1]
	if strings.HasPrefix(loc, "http") {
		doc, err = goquery.NewDocument(loc)
		if err != nil {
			fmt.Printf("error reading URL: %s\n", err)
			os.Exit(1)
		}
	} else {
		b, err := ioutil.ReadFile(loc)
		if err != nil {
			fmt.Printf("error reading file: %s\n", err)
			os.Exit(1)
		}
		doc, err = goquery.NewDocumentFromReader(bytes.NewReader(b))
		if err != nil {
			fmt.Printf("error parsing file as HTML: %s\n", err)
			os.Exit(1)
		}
	}

	// find article body
	body = doc.Find("body")
	if len(os.Args) > 2 {
		body = doc.Find(os.Args[2])
	}

	var c int
	var p float32
	sentenceCount := countSentences(body.Text())

	c = countWords(body.Text())
	addResult(fmt.Sprintf("The text contains %d words.", c), c >= 300)

	c = countHeadings(body)
	addResult(fmt.Sprintf("The text contains %d subheadings.", c), c > 0)

	c = countHeadingsWithWords(body, 300)
	addResult(fmt.Sprintf("%d of the subheadings is followed by more than 300 words.", c), c <= 1)

	c = countParagraphsWithWords(body, 150)
	addResult(fmt.Sprintf("%d of the paragraphs contains more than 150 words.", c), c < 1)

	c = countSentencesWithTransitionWord(body.Text())
	p = float32(c) / float32(sentenceCount) * 100
	addResult(fmt.Sprintf("%.1f%% of the sentences contain a transition word or phrase.", p), p >= 30)

	// long sentences
	c = countSentencesWithWords(body, 20)
	p = float32(c) / float32(sentenceCount) * 100
	addResult(fmt.Sprintf("%.1f%% of the sentences contain more than 20 words.", p), p <= 25)

	// kincaid
	kc := calculateKincaid(body.Text())
	addResult(fmt.Sprintf("The copy scores %1.f in the Flesch Reading Ease test.", kc), kc >= 60)

	fmt.Printf("Analysing \u001B[4m%s\u001B[24m\n", loc)
	for _, r := range results {
		if r.good {
			fmt.Printf("\u001B[32m+ %s\u001B[39m\n", r.text)
		} else {
			fmt.Printf("\u001B[31m- %s\u001B[39m\n", r.text)
		}

	}
}

func addResult(t string, g bool) {
	results = append(results, result{text: t, good: g})
}
