package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const WordLength = 5

func ScanWord() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter 5 letter word: ")
		word, _ := reader.ReadString('\n')
		word = strings.TrimSuffix(word, "\n")
		word = strings.ToLower(word)

		if len(word) != WordLength {
			fmt.Println("Invalid word length")
			continue
		}
		return word
	}
}
