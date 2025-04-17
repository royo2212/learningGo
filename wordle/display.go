package main

import (
	"fmt"
)

func PrintGuessWithColors(feedback []LetterFeedback) {
	for _, letter := range feedback {
		color := getColorByStatus(letter.Status)
		fmt.Print(color, string(letter.Letter), resetColor, " ")
	}
	fmt.Println()
}

func PrintKeyboard(keyboard map[rune]string) {
	fmt.Println("\nKeyboard:")
	for ch := 'a'; ch <= 'z'; ch++ {
		status, ok := keyboard[ch]
		if ok {
			color := getColorByStatus(status)
			fmt.Print(color, string(ch), resetColor, " ")
		} else {
			fmt.Print(string(ch), " ")
		}

		if ch == 'm' || ch == 'z' { // break lines like real keyboard
			fmt.Println()
		}
	}
	fmt.Println()
}

func ShowWelcomeMessage() {
	fmt.Println("==========================")
	fmt.Println("    Welcome to Wordle!    ")
	fmt.Println("==========================")
	fmt.Println("Guess the 5-letter word.\n")
}

func ShowWinMessage(secretWord string, tries int) {
	fmt.Println("\n==========================")
	fmt.Printf("🎉 You won in %d tries!\n", tries)
	fmt.Println("The word was:", secretWord)
	fmt.Println("==========================")
}

func ShowLoseMessage(secretWord string) {
	fmt.Println("\n==========================")
	fmt.Println("💀 Game Over!")
	fmt.Println("The word was:", secretWord)
	fmt.Println("==========================")
}

const (
	greenColor  = "\033[1;42m" // Bright Green Background
	yellowColor = "\033[1;43m" // Bright Yellow Background
	grayColor   = "\033[1;47m" // White/Grey Background
	resetColor  = "\033[0m"    // Reset color
)

func getColorByStatus(status string) string {
	switch status {
	case StatusGreen:
		return greenColor
	case StatusYellow:
		return yellowColor
	case StatusGray:
		return grayColor
	default:
		return resetColor
	}
}
