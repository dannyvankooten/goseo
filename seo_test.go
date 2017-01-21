package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"testing"
)

func TestCountWords(t *testing.T) {
	r := bytes.NewReader([]byte("Hello these are some words. Count them for me, will you?"))
	doc, _ := goquery.NewDocumentFromReader(r)
	count := countWords(doc.Children().Text())
	if count != 11 {
		t.Errorf("Expected %d words, got %d", 11, count)
	}

}
