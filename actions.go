package main

import (
	"fmt"
	"sort"
	"strings"
)

func printStatistics(text string) {

	doc := getDocument(text)

	words := getWordsFromTokens(doc.Tokens())
	numberOfLetters := 0
	for _, word := range words {
		numberOfLetters += len(word)
	}

	fmt.Println("Number of phrases: ", len(doc.Sentences()))
	fmt.Println("Number of words: ", len(words))
	fmt.Println("Number of letters: ", numberOfLetters)

}

func printLongestWords(text string, n int) {

	doc := getDocument(text)

	words := getWordsFromTokens(doc.Tokens())

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	fmt.Println("The ranking of longest words are:")
	printedWords := make([]string, 0)
	for _, word := range words {
		if len(printedWords) < n {
			if !contains(printedWords, word) {
				printedWords = append(printedWords, word)
				fmt.Printf("\t%d) %s - %d\n", len(printedWords), word, len(word))

			}
		}
	}
}

func printMostFrequentWords(text string, n int) {

	stopWords := getStopWords()
	doc := getDocument(text)

	words := getWordsFromTokens(doc.Tokens())

	wordCount := make(map[string]int)

	for _, word := range words {
		word = strings.ToLower(word)
		if !contains(stopWords, word) {
			wordCount[word]++
		}

	}

	keys := make([]string, 0, len(wordCount))
	for key := range wordCount {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return wordCount[keys[i]] > wordCount[keys[j]]
	})

	fmt.Println("The ranking of most frequent words are:")
	for i, key := range keys {
		if i < n {
			fmt.Printf("\t%d) %s - %d\n", i+1, key, wordCount[key])
		}
	}

}

func printCategorization(text string) {

	doc := getDocument(text)

	var nounCount, verbCount, adjectiveCount int
	for _, tok := range doc.Tokens() {
		switch {
		case tok.Tag == "NN" || tok.Tag == "NNS" || tok.Tag == "NNP" || tok.Tag == "NNPS":
			nounCount++
		case tok.Tag == "VB" || tok.Tag == "VBD" || tok.Tag == "VBG" || tok.Tag == "VBN" || tok.Tag == "VBP" || tok.Tag == "VBZ":
			verbCount++
		case tok.Tag == "JJ" || tok.Tag == "JJR" || tok.Tag == "JJS":
			adjectiveCount++
		}
	}

	fmt.Printf("Nouns: %d, Verbs: %d, Adjectives: %d\n", nounCount, verbCount, adjectiveCount)

}