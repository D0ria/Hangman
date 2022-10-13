package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Game struct {
	Word             func     // Word composed of '_', ex: H_ll_
	ToFind           func     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	Found            bool       //If the letter was found, return true
}

func main() {
	var morigno Game
	morigno.Word = RandomWord()
	morigno.ToFind = letterInWord()
	morigno.Attempts = 3
	morigno.Found = false
	morigno.HangmanPositions = "hangman.txt"
}

func Reader() {
	f, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func RandomWord(i int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(i)
}

func HiddenWord(word string) {
	for _, letters := range word {
		if morigno.Found == false {
			letters += rune(95)
		}
	}
}

func RevealFewLetters(j string) {
	for n := 0; n < len(j)-1; n++ {
		rand.Seed(time.Now().Unix())
		lettre := j[rand.Intn(len(j))]
		fmt.Println(lettre)
	}
}

func letterInWord(find string, letters []string) bool {
	for _, letter := range letters {
		if find == letter {
			return true
		}
	}
	morigno.Attempts -= 1
	return false
}

func Step(tab) {

}

func EndOfGame() {
	if morigno.ToFind == RandomWord {
		return "Gagné"
	} if morigno.Attempts == 0 {
		return "Perdu"
	}
}
