package main

import (
	"fmt"
	"os"
	"sort"
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
		word := strings.Trim(tok.Text, ".,:;!?¿¡()“”")
		if(word != "" && word != " "){
			words = append(words,  word)
		}
	}
	return words
}

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

func printLongestWords(text string, n int){

	doc := getDocument(text)

	words := getWordsFromTokens(doc.Tokens())
			
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	fmt.Println("The ranking of longest words are:")
	for i, word := range words {
			if i < n {
					fmt.Printf("\t%d) %s - %d\n", i+1, word, len(word))
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

func showMenu() {
	printDivider()
	fmt.Println("Select the option you want to execute:")
	fmt.Println("1) Statistics")
	fmt.Println("2) Categorization")
	fmt.Println("3) Longest words")
	fmt.Println("4) Exit")
}

func getUserChoice() int {
    var choice int
    fmt.Scanln(&choice)
    return choice
}

func printDivider() {
	fmt.Println("----------------------------------------------------")

}


func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide the file path as an argument")
		os.Exit(1)
	}

	filepath := os.Args[1]
	dat, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file...")
		os.Exit(0)
	}
	text := string(dat)

	for  {
		showMenu()
		option := getUserChoice()
		printDivider()
		switch option {
		case 1:
			printStatistics(text)
		case 2:
			printCategorization(text)
		case 3:
			var n int
			fmt.Println("Enter the number of longest words you want to see:")
			n = getUserChoice()
			printLongestWords(text, n)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}


	

	
	


	 


}