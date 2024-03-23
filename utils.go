package main

import (
	"strings"

	"github.com/jdkato/prose/v2"
)

func getDocument(text string) *prose.Document {
	doc, err := prose.NewDocument(text, prose.WithExtraction(false))
	if err != nil {
		panic(err)
	}
	return doc
}

func getWordsFromTokens(tokens []prose.Token) []string {
	words := make([]string, 0)
	for _, tok := range tokens {
		if strings.Contains(tok.Text, "—") {
			words = append(words, strings.Split(tok.Text, "—")...)
			continue
		}
		word := strings.Trim(tok.Text, ".,:;!?¿¡()\"“”")
		if word != "" && word != " " {
			words = append(words, word)
		}
	}
	return words
}

func getStopWords() []string {
	return getStopWordsFromFile()
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
