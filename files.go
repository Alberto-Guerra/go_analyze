package main

import (
	"fmt"
	"os"
	"strings"
)

func getTextFromFile() string {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the file path as an argument")
		os.Exit(1)
	}

	filepath := os.Args[1]
	dat, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file...")
		os.Exit(1)
	}
	text := string(dat)
	return text
}

func getStopWordsFromFile() []string {
	dat, err := os.ReadFile("stopwords.txt")
	if err != nil {
		fmt.Println("Error reading file...")
		os.Exit(1)
	}
	stopWords := string(dat)
	stopWordsList := strings.Split(stopWords, "\r\n")
	return stopWordsList

}
