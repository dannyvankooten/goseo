package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"testing"
)

var s *goquery.Selection
var c int

func TestMain(m *testing.M) {
	b, _ := ioutil.ReadFile("data/test-article.html")
	r := bytes.NewReader(b)
	doc, _ := goquery.NewDocumentFromReader(r)
	s = doc.Find("body")
	m.Run()
}

func TestCountHeadings(t *testing.T) {
	c = countHeadings(s)
	if c != 4 {
		t.Errorf("Expected 4, got %d", c)
	}
}

func TestCountHeadingsWithWords(t *testing.T) {
	c = countHeadingsWithWords(s, 1)
	if c != 4 {
		t.Errorf("Expected 4, got %d", c)
	}

	c = countHeadingsWithWords(s, 20)
	if c != 2 {
		t.Errorf("Expected 2, got %d", c)
	}
}

func TestCountSentencesWithWords(t *testing.T) {
	c = countSentencesWithWords(s, 20)
	if c != 0 {
		t.Errorf("Expected 0, got %d", c)
	}

	c = countSentencesWithWords(s, 5)
	if c != 24 {
		t.Errorf("Expected 24, got %d", c)
	}
}

func TestCountParagraphsWithWords(t *testing.T) {
	c = countParagraphsWithWords(s, 1)
	if c != 6 {
		t.Errorf("Expected 6, got %d", c)
	}
}
