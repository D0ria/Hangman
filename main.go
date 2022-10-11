package hangman

import (
	"fmt"
)

func main() {

	var hangman Game
	hangman.Word = RandomWord
	hangman.ToFind = HiddenWord
	hangman.Attempts = 3
	hangman.positions = Step()
	hangman.Found = false

	fmt.Printf(RevealFewLetters(letter))
	fmt.Println(hangman)
}
