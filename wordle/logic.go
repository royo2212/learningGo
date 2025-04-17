package main

type LetterFeedback struct {
	Letter rune
	Status string
}

const (
	StatusGreen  = "green"
	StatusYellow = "yellow"
	StatusGray   = "gray"
)

func CheckGuess(secretWord, guess string) []LetterFeedback {
	feedback := make([]LetterFeedback, WordLength)
	remainingLetter := make(map[rune]int)
	for i := 0; i < WordLength; i++ {
		if guess[i] == secretWord[i] {
			feedback[i] = LetterFeedback{
				Letter: rune(guess[i]),
				Status: StatusGreen,
			}
		} else {
			remainingLetter[rune(secretWord[i])]++
		}
	}
	for i := 0; i < WordLength; i++ {
		if feedback[i].Status == StatusGreen {
			continue
		}
		letter := rune(guess[i])
		if remainingLetter[letter] > 0 {
			feedback[i] = LetterFeedback{
				Letter: letter,
				Status: StatusYellow,
			}
			remainingLetter[letter]--
		} else {
			feedback[i] = LetterFeedback{
				Letter: letter,
				Status: StatusGray,
			}
		}
	}
	return feedback
}
func IsWin(feedback []LetterFeedback) bool {
	for _, letterFeedback := range feedback {
		if letterFeedback.Status != StatusGreen {
			return false
		}
	}
	return true
}

func UpdateKeyboardState(keyboard map[rune]string, feedback []LetterFeedback) map[rune]string {
	for _, f := range feedback {
		currentStatus, exists := keyboard[f.Letter]

		if !exists || isBetterStatus(currentStatus, f.Status) {
			keyboard[f.Letter] = f.Status
		}
	}

	return keyboard
}

var statusRank = map[string]int{
	StatusGray:   1,
	StatusYellow: 2,
	StatusGreen:  3,
}

func isBetterStatus(current, new string) bool {
	return statusRank[new] > statusRank[current]
}
