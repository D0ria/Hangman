package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Game struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
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

func Mot_random(i int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(i)
}

func RevealFewLetters(j string) {
	for n := 0; n < len(j)-1; n++ {
		rand.Seed(time.Now().Unix())
		lettre := j[rand.Intn(len(j))]
		fmt.Println(lettre)
	}
}

func HiddenWord() {

}
