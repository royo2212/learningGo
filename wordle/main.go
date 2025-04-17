package main

import (
	"fmt"
)

const MaxTries = 6

func main() {
	ShowWelcomeMessage()
	secretWord, err := GetRandomWordFromAPI()
	if err != nil {
		fmt.Println("Error fetching word:", err)
		return
	}
	//fmt.Println("the word is: ", secretWord)
	keyboard := make(map[rune]string)

	for tries := 1; tries <= MaxTries; tries++ {
		guess := ScanWord()

		feedback := CheckGuess(secretWord, guess)

		PrintGuessWithColors(feedback)

		keyboard = UpdateKeyboardState(keyboard, feedback)

		PrintKeyboard(keyboard)

		if IsWin(feedback) {
			ShowWinMessage(secretWord, tries)
			return
		}

		fmt.Printf("Tries left: %d\n\n", MaxTries-tries)
	}

	ShowLoseMessage(secretWord)
}
