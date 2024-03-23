package main

import (
	"fmt"
	"os"
)

func showMenu() {
	printDivider()
	fmt.Println("Select the option you want to execute:")
	fmt.Println("1) Statistics")
	fmt.Println("2) Categorization")
	fmt.Println("3) Longest words")
	fmt.Println("4) Most frequent words")
	fmt.Println("5) Exit")
}

func getUserChoice() int {
	var choice int
	fmt.Scanln(&choice)
	return choice
}

func printDivider() {
	fmt.Println("----------------------------------------------------")

}

func initMenu(text string) {
	for {
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
			var n int
			fmt.Println("Enter the number of most frequent words you want to see:")
			n = getUserChoice()
			printMostFrequentWords(text, n)
		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}