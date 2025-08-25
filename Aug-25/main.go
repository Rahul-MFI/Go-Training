package main

import (
	"Aug-25/counter"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string as an argument to print the count of words")
		return
	}

	sentence := os.Args[1]
	sentence = strings.ReplaceAll(sentence, `\n`, "\n")
	sentence = strings.ReplaceAll(sentence, `\t`, "\t")
	fmt.Printf("The count of words in the sentence is %d\n", counter.CountWords(sentence))
}
